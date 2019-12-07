'use strict';

const async = require('async');
const request = require('request');

const config = require('../config');
const Conversations = require('../models/conversations');
const conversationService = require('../services/conversation');

module.exports = {

    send: function (params, callback) {
        const conversationId = params.conversationId;
        const senderId = params.senderId;
        const requester = params.requester;
        const requestees = params.requestees;

        Conversations.findOne({ conversationId: conversationId }, function (err, conversation) {
            if (err) {
                return callback(err);
            }

            if (!conversation) {
                return callback(new Error('conversation does not exist'));
            }

            // check for any selected option
            // if yes, save message for the current
            // user's conversation
            let hasSelectedOption;
            requester.options.forEach(function (option) {
                if (option.isSelected) {
                    hasSelectedOption = true;
                }
            });

            const currMessageListLength = conversation.messageBlobList.length;
            if (hasSelectedOption && conversation.messageBlobList.length > 0) {
                conversation.messageBlobList[currMessageListLength - 1].options = requester.options;
                conversation.messageBlobList[currMessageListLength - 1].messages = requester.messages;
            } else {
                conversation.messageBlobList.push({
                    options: requester.options,
                    messages: requester.messages
                });
            }

            const options = {
                url: config.getBotConnectionString(),
                body: {
                    senderId: senderId,
                    action: requester.options,
                    requestees: requestees
                }
            };
            request.post(options, function (botError, botResponse) {
                if (botError) {
                    return callback(botError);
                }

                conversation.save(function (err) {
                    if (err) {
                        return callback(err);
                    }
    
                    // broadcasting notifications to requested users
                    async.each(requestees, function (requestee, eCB) {
                        Conversations.findOne({ conversationId: requestee.id }, function (err, requesteeConversation) {
                            if (err) {
                                return eCB(new Error('Error while sending message to: ' + requestee.id));
                            }
    
                            if (requesteeConversation) {
                                requesteeConversation.messageBlobList.push({
                                    'options': requestee.options,
                                    'messages': requestee.messages,
                                    'createdAt': Date.now()
                                });
            
                                requesteeConversation.save(eCB);
                            } else {
                                const conversationCreateParams = {
                                    'conversationId': requestee.id,
                                    'title': 'New message',
                                    'messageBlob': {
                                        'options': requestee.options,
                                        'messages': requestee.messages,
                                        'createdAt': Date.now()
                                    }
                                };
                                conversationService.create(conversationCreateParams, eCB);
                            }
                        });
                    }, callback);
    
                });
            });
        });
    }

}