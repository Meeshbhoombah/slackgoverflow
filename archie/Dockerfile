FROM ubuntu:latest
MAINTAINER Rohan Mishra @Meeshbhoombah

RUN apt-get update && \
    apt-get install -y python3-pip python3-dev git && \
    cd /usr/local/bin && \
    ln -s /usr/bin/python3 python && \
    pip3 install --upgrade pip

WORKDIR /archie

COPY requirements.txt /archie
COPY run.py /archie
COPY ./app /archie

RUN pip install -r requirements.txt
ENTRYPOINT ['python']
CMD ['run.py']
 
EXPOSE 5000

