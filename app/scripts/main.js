'use strict';

var angular = require('angular'),
    MainController = require('./controllers/Main');

angular.module('thaiQuizApp', [])
.controller('MainController', MainController);
