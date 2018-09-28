/*
 * DockerDaemon.js
 * 
 *
 */
"use strict";


const Docker   = require('dockerode')
,     daemon   = new Docker()
,     utils    = require('./utils')


// create `archnet` network 
daemon.createNetwork({ Name: 'archnet', CheckDuplicate: true }, function(err, msg) {

    // nly if message in response
    if (msg) {
        log(msg)
        vital('ARCHNET NETWORK BOOTED: WAITING FOR NODES...')
    }

    if (err && err.statusCode == 201)  {                  
        vital('ARCHNET NETWORK EXISTS...')
    } 

    if (!err) {
        shout(err.json.message);
        process.exit(1);
    }
    });


    // build containers
    // expose ports and connect to network
    // create coinbase for mining rewards
    // start mining using `geth` command line

});

