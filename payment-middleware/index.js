'use strict';

const express = require('express');
const app = express();

const models = require('./models');
const apiController = require('./controllers/apiController');
apiController(app);

models.init();

const port = process.env.port || 3000;

app.listen(port, function () {
    console.log('Payment Middleware is listening to port: ' + port);

    process.on('SIGINT', function () {
        models.stop();
    });
});
