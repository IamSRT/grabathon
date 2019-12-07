'use strict';

const bodyParser = require('body-parser');
const request = require('request');

const config = require('../config');
const conversation = require('../services/conversation');
const messsage = require('../services/message');
const utilities = require('../utilities');

module.exports = function (app) {

	app.use(bodyParser.json());
	app.use(bodyParser.urlencoded({ extended: true }));

	app.post('/api/conversations', function (req, res) {
		const params = req.body;
		const conversationId = req.headers.userid;
		const title = params.title;

		const createConversationParams = {
			'conversationId': conversationId,
			'title': title
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

			// call Bot for first set of messages to be displayed
            const options = {
                url: config.getBotConnectionString()
            };
			request.get(options, function (botError, botResponse) {
				if (botError) {
					return res.status(200).send({
						'status': 'error',
						'statusCode': 0,
						'statusMessage': 'Internal error, please try reloading',
						'data': []
					});					
				}
				return res.status(200).send({
					'status': 'ok',
					'statusCode': 1,
					'statusMessage': 'Conversation Created',
					'data': []
				});
			});
		});
	});

	app.get('/api/conversations/:conversationId', function (req, res) {
		const userId = req.headers.userid;

		const params = {
			'conversationId': req.params.conversationId
		};
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
		messsage.send(messageSendParams, function (err, result) {
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
