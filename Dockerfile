FROM drewgonzales360/drew:latest

ENV PATH=/usr/local/go/bin:$PATH

COPY goenv /usr/local/bin
COPY scripts/goenv-test.sh /usr/local/bin/goenv-test
