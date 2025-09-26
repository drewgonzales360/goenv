FROM ubuntu:noble

ENV PATH=/usr/local/go/bin:$PATH
COPY goenv /usr/local/bin
COPY scripts/goenv-test.bash /usr/local/bin/goenv-test

RUN \
    apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y sudo ca-certificates curl && \
    useradd -m drew && \
    echo "drew ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers

USER drew
