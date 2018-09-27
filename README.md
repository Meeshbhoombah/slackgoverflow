# Archnet
Command Line tool for building, deploying, and interacting with the sandboxed 
`archnet` Ethereum (private) network and connected Swarm and Whisper nodes.

## Getting Started
These instructions will get you a copy of the project up and running on your 
local machine for development and testing purposes. See deployment for notes 
on how to deploy the project on a live system.

### Prerequisites
Archnet takes advantage of Docker to run a sandboxed replication of the MVP
`testnet` on your local machine.
```
$ docker --version
Docker version 18.06.1-ce, build e68fc7a
```

If you don't have Docker up and running, you can set up and/or get oriented 
using the [instructions](https://docs.docker.com/v17.09/get-started/).

### Installing
Upon cloning this repo and `cd`-ing into the cloned directory, install `archnet`
globally via `npm`.
```
$ npm i -g archnet
```

This will allow you to manipulate your local copy of the `archnet` network at
any time.

