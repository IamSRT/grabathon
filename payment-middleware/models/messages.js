const mongoose = require('mongoose');
const Schema = mongoose.Schema;

const OptionsSchema = new Schema({
    id: String,
    title: String,
    description: String,
    renderType: String,
    isSelected: { type: Boolean, default: false }
});

const MessageSchema = new Schema({
    message: String,
    messageType: String
});

const MessageBlobSchema = new Schema({
    messages: [MessageSchema],
    options: [OptionsSchema],
	createdAt: Number
});

module.exports = MessageBlobSchema;
