'use strict';

const Conversations = require('../models/conversations');

module.exports = {

    create: function (params, callback) {
        const conversationId = params.conversationId;
        const title = params.title;
        const messageBlob = params.messageBlob;

        const conversation = Conversations();
        conversation.conversationId = conversationId;
        conversation.title = title;
        if (messageBlob) {
            conversation.messageBlobList.push(messageBlob);
        }
        conversation.save(callback);
    },

    get: function (params, callback) {
        Conversations.findOne({ conversationId: params.conversationId }, callback);
    }

}
