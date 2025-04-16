FROM ubuntu:latest
LABEL authors="toby"

ENTRYPOINT ["top", "-b"]