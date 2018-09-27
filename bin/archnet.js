#!/usr/bin/env node

"use strict";


var Docker   = require('dockerode')
,   daemon   = new Docker()
,   program  = require('commander');


// $ -v, --version
program
    .version('0.0.1', '-v, --version');
    

// $ archnet create
program
    .command('create')
    .description('Build and deploy `archnet` in the default `dev` env.')
    .option('-g, --geth <# of nodes>', 'Set INT value of `geth` nodes for the network (two by default).')
    .option('-p, --password <password>', 'Set a custom password for use throughout network.')
    .action(function(create) {
        console.log(create.geth);
    });
    

program.parse(process.argv);

