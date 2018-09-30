#!/usr/bin/env node
"use strict";


const program  = require('commander')
,     Docker   = require('dockerode')
,     daemon   = new Docker()
,     chalk    = require('chalk');


const log      = console.log
,     vital    = function(str) {log(chalk.blueBright.bold(str));}
,     shout    = function(str) {log(chalk.redBright("Error: " + str));};


/* TODO: 
[x] create `archnet` network
[] build  containers based off `-g` command
[] expose ports and connect to network
[] create coinbase for mining rewards
[] start mining using `geth` command line
*/


// $ -v, --version
program
    .description('Build, deploy, and interact with a local replica of the `archnet` testnet.')
    .version('0.0.1', '-v, --version');


// $ archnet boot
program
    .command('boot')
    .description('Build and deploy `archnet` in the default `dev` config.')
    .option('-g, --geth-nodes <# of nodes>', 
        'Set the number of `geth` nodes on the network (two by default).')
    .action(function(create) {  
        vital('BOOTING ARCHNET...');
       
        daemon.createNetwork({ Name: 'archnet', CheckDuplicate: true }, function(err, msg) {
            if (err && err.statusCode == 409)  {                  
                log('ARCHNET NETWORK EXISTS, WAITING FOR NODES...');
            } else if (err && err.errno == 'ECONNREFUSED') {
                shout('ECONNREFUSED (start Docker on local machine)')
            } else if (err) {
                log(err);
                process.exit(1);
            } else {
                log('NETWORK CREATED AT: ' + msg.id);
                vital('ARCHNET NETWORK BOOTED: WAITING FOR NODES...');
            }
        });

        

    });


// $ archnet stop 
program
    .command('destroy')
    .description('Stop the `archnet` network.')
    .option('-c, --clean',
        'Destroy the network and nodes after stoping')
    .action(function(destroy) {
        
    });
        

program.parse(process.argv);


// $ archnet
if (!program.args.length) program.help();

