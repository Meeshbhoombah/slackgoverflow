#!/usr/bin/env node
"use strict";


const program  = require('commander')
,     Docker   = require('dockerrode')
,     daemon   = new Docker()
,     chalk    = require('chalk');


const log      = console.log;
,     vital    = function(str) {log(chalk.blueBright.bold(str));};
,     shout    = function(str) {log(chalk.redBright("Error: " + str));};


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
        
        utils.vital('BOOTING ARCHNET...')
        daemon.createNetwork({ Name: 'archnet', CheckDuplicate: true }, function(err, msg) {

            // only if message in response
            if (msg) {
                utils.log(msg)
                utils.vital('ARCHNET NETWORK BOOTED: WAITING FOR NODES...')
            }

            if (err && err.statusCode == 409)  {                  
                utils.log('ARCHNET NETWORK EXISTS...');
            } else {
                shout(err.json.message);
                process.exit(1);
            }

            // build containers
            // expose ports and connect to network
            // create coinbase for mining rewards
            // start mining using `geth` command line
        });
    });


// $ archnet stop
program
    .command('destroy')
    .description('Stop the `archnet` network.')
    .option('-c, --clean',
        'Destroy and remove the network and nodes after stoping')
    .action(function(destroy) {

    });
        

program.parse(process.argv);


// $ archnet
if (!program.args.length) program.help();

