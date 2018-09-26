# Architect Ethereum Network
Configured for optimal runtime of the Architect MVP. To inspect the network use
the [Architect network explorer](#).

## Getting Started
Set up the architect network.
```
$ archnet up
```

Spawn as many nodes as necessary (recommended three max, default is two).
```
$ archnet up --geth-node 3
```

## Docker Commands (temp)
```
docker network create archnet
```

Build any number of nodes (recommended max 3).
```
docker build -t node_one .
docker build -t node_two .
```

