const mongoose = require('mongoose');
const Schema = mongoose.Schema;

const MessageBlobSchema = require('./messages');

const conversationsSchema = new Schema({
	conversationId: { type: String, index: true },
	title: { type: String },
	messageBlobList: [MessageBlobSchema],
	createdAt: { type: Date, default: Date.now },
	updatedAt: { type: Date, default: Date.now }
});

const Conversations = mongoose.model('Conversations', conversationsSchema);

module.exports = Conversations;
