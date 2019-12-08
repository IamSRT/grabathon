var config = require('./config');

module.exports = {

	getDbConnectionString: function () {
		return 'mongodb://' + config.mongo[process.env.NODE_ENV].host + '/middleware';
	},

	getBotConnectionString: function (modeId) {
		return 'http://' + process.env.BOT_URL + ':8080/v1/' + modeId + '/next/steps';
	}

}