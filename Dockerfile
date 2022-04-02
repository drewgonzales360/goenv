FROM drewgonzales360/drew:latest

COPY goenv /usr/local/bin
COPY scripts/goenv-test.sh /usr/local/bin/goenv-test
