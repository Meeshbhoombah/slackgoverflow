#!/usr/bin/env node

"use strict";


var program = require('commander');


program
    .version('0.0.1', '-v, --version');
    

// $ archnet create
program
    .command('create')
    .description('Build and deploy `archnet` in the default `dev` env.')
    .option('-g, --geth <int>', 'Set number of `geth` nodes for the network (two by default).')
    .option('-p, --password <password>', 'Set a custom password for use throughout network.')
    .action(function() {
        console.log('Hello World');
    });
    

program.parse(process.argv);

