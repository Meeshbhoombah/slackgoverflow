#!/usr/bin/env node

"use strict";


const Docker   = require('dockerode')
,     daemon   = new Docker()
,     program  = require('commander')
,     chalk    = require('chalk');


const log      = console.log;
const err      = console.err;
const notif    = function(str) {log(chalk.blueBright(str));};


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
     
        // create `archnet` network 
        notif('Creating archnet...');
        docker.createNetwork({ Name: 'archnet' }) {

        }

        // build containers
        // expose ports and connect to network
        // create coinbase for mining rewards
        // start mining using `geth` command line
    
    });


program.parse(process.argv);

// $ archnet
if (!program.args.length) program.help();

