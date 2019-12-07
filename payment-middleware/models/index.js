'use strict';

const mongoose = require('mongoose');

const config = require('../config');

module.exports = {

    init: function () {
        const dbURL = config.getDbConnectionString();

        mongoose.connect(dbURL, { useNewUrlParser: true });

        mongoose.connection.on('connected', function () {  
            console.log("Mongoose default connection is open to ", dbURL);
         });

        mongoose.connection.on('error', function (err) {
            console.log("Mongoose default connection has occured " + err + " error");
       });
    },

    stop: function () {
        mongoose.connection.close(function() {
            console.log("Mongoose default connection is disconnected due to application termination");
            process.exit(0);
        });
    }
}