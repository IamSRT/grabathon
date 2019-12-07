var config = require('./config');

module.exports = {

	getDbConnectionString: function () {
		return 'mongodb://' + config.mongo[process.env.NODE_ENV].host + '/middleware';
	},

	getBotConnectionString: function () {
		return 'https://' + process.env.BOT_URL;
	}

}