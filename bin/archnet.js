#!/usr/bin/env node
"use strict";


const program  = require('commander')
,     :wq
,     utils    = require('../src/utils.js');


// $ -v, --version
program
    .description('Build, deploy, and interact with a local replica of the `archnet` testnet.')
    .version('0.0.1', '-v, --version');


// $ archnet up -g
program
    .command('create')
    .description('Build and deploy `archnet` in the default `dev` config.')
    .option('-g, --geth-nodes <# of nodes>', 
        'Set the number of `geth` nodes on the network (two by default).')
    .action(function(create) {

        
    
    };
        

program.parse(process.argv);


// $ archnet
if (!program.args.length) program.help();

