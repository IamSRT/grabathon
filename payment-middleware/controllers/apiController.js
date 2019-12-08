'use strict';

const bodyParser = require('body-parser');
const request = require('request');

const config = require('../config');
const conversation = require('../services/conversation');
const message = require('../services/message');
const utilities = require('../utilities');

module.exports = function (app) {

	app.use(bodyParser.json());
	app.use(bodyParser.urlencoded({ extended: true }));

	app.post('/api/conversations', function (req, res) {
		const params = req.body;
		const senderId = req.headers.userid;
		const title = params.title;
		const modeId = params.modeId;
		const messages = params.messages;
		const payment = params.payment;

		const createConversationParams = {
			'conversationId': senderId,
			'title': title,
			'messageBlob': {
				'options': [],
				'messages': messages
			}
		};
		conversation.create(createConversationParams, function (err, result) {
			if (err) {
				console.log(err);
				return res.status(200).send({
					'status': 'error',
					'statusCode': 0,
					'statusMessage': 'Error creating conversation'
				});
			}
			// console.log(result);

			// call Bot for first set of messages to be displayed
            const options = {
				method: 'POST',
				uri: config.getBotConnectionString(modeId),
				headers : {
					'content-type': 'application/json'
				},
				body: {
					isRequester: true,
					UserId: senderId,
					payment: payment,
					modeId: modeId,
                    options: [],
                    messages: messages
				},
				json: true
			};
			request(options, function (err, botResponse) {
				const messageSendParams = {
					'conversationId': senderId,
					'senderId': senderId,
					'requester': botResponse.body.data.requester,
					'requestees': botResponse.body.data.requestees
				};
				message.send(messageSendParams, function () {
					return res.status(200).send({
						'status': 'ok',
						'statusCode': 1,
						'statusMessage': 'Conversation Created',
						'data': botResponse.body.data
					});
				});
			});
		});
	});

	app.get('/api/conversations/:conversationId', function (req, res) {
		const userId = req.headers.userid;

		const params = {
			'conversationId': req.params.conversationId
		};
		console.log(req.headers);
		console.log(params);
		conversation.get(params, function (err, result) {
			if (err) {
				console.log(err);
				return res.status(200).send({
					'status': 'error',
					'statusCode': 0,
					'statusMessage': 'Error getting conversation details'
				});
			}

			if (!result) {
				return res.status(200).send({
					'status': 'error',
					'statusCode': 0,
					'statusMessage': 'Conversation not found'
				});
			}

			const conversationObj = utilities.getConversationAndMessages(result);

			if (conversationObj.conversationId != userId) {
				return res.status(403).send({
					'status': 'forbidden',
					'statusCode': 0,
					'statusMessage': 'conversation does not belong to current user'
				});
			}

			return res.status(200).send({
				'status': 'ok',
				'statusCode': 1,
				'data': conversationObj
			});
		});
	});

	app.post('/api/conversations/:conversationId/messages', function (req, res) {
		const params = req.body;
		const conversationId = req.params.conversationId;
		const senderId = req.headers.userid;

		const requester = params.requester;
		const requestees = params.requestees;

		const messageSendParams = {
			'conversationId': conversationId,
			'senderId': senderId,
			'requester': requester,
			'requestees': requestees
		};
		console.log(messageSendParams);
		message.send(messageSendParams, function (err, result) {
			if (err) {
				console.log(err);
				return res.status(200).send({
					'status': 'error',
					'statusCode': 0,
					'statusMessage': 'Error while sending message'
				});
			}

			return res.status(200).send({
				'status': 'ok',
				'statusCode': 1,
				'statusMessage': 'Message sent successfully'
			});
		});
	});

};
