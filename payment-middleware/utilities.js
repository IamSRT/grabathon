'use strict';

module.exports = {

    getConversationAndMessages: function (params) {
        return {
            'conversationId': params.conversationId,
            'title': params.title,
            'messageHistory': params.messageBlobList
        };
    }

}