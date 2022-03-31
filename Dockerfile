FROM ubuntu:latest

COPY goenv /usr/local/bin

RUN \
    apt-get update \
    && apt-get install -y ca-certificates
