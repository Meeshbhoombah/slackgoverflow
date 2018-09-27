# Archnet
Command Line tool for building, deploying, and interacting with the sandboxed 
`archnet` Ethereum (private) network and connected Swarm and Whisper nodes.

## Getting Started
These instructions will get you a copy of the project up and running on your 
local machine for development and testing purposes.

### Prerequisites
Archnet takes advantage of Docker to run a sandboxed replica of the MVP/test
network on your local machine.
```
$ docker --version
Docker version 18.06.1-ce, build e68fc7a
```

If you don't have Docker up and running, you can set up and/or get oriented 
using this [guide](https://docs.docker.com/v17.09/get-started/).

Archnet requires Node.js `v10.7.0` or higher as it uses ES6 syntax and features.
```
$ node -v
v10.7.0
```

### Installing
Install `archnet` globally via `npm`.
```
$ npm i -g archnet
```

You can also install `archnet` directly from the source repo. Clone this 
repository.
```
$ git clone https://github.com/archproj/archnet
Cloning into 'archnet'...
...
```

After cloning the repository `cd` into the directory and install the package 
globally from within the directory.
```
$ cd archnet
$ npm i -g
```

