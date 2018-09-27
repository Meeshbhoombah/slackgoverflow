# Scaleable Archnet `geth` node with custom configuration

FROM ethereum/client-go:alpine

LABEL version='0.0.1'
LABEL maintainer='rohanmishra2017@gmail.com'

# Create a new user "geth_archnet" with home directory `/archnet`
RUN adduser -D -g "" getharch

COPY /config /home/getharch/config
# Give persmission to interact with directory
RUN chown -R getharch:getharch /home/getharch/config
USER getharch
WORKDIR /home/getharch

