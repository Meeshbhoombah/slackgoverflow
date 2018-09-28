/*
 * utils.js
 * 
 *
 */
"use strict";




const chalk    = require('chalk');
,     log      = console.log
,     vital    = function(str) {log(chalk.blueBright.bold(str));}
,     shout    = function(str) {log(chalk.redBright("Error: " + str));}



