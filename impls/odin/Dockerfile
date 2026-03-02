FROM ubuntu:24.04

##########################################################
# General requirements for testing or common across many
# implementations
##########################################################

RUN apt-get -y update

# Required for running tests
RUN apt-get -y install make python3
RUN ln -fs /usr/bin/python3 /usr/local/bin/python

RUN mkdir -p /mal
WORKDIR /mal

##########################################################
# Specific implementation requirements
##########################################################

RUN DEBIAN_FRONTEND=noninteractive apt-get -y install \
    build-essential ca-certificates clang curl git libreadline-dev lld llvm

RUN git clone --depth=1 https://github.com/odin-lang/Odin /opt/Odin
RUN make -C /opt/Odin release-native
RUN ln -fs /opt/Odin/odin /usr/local/bin/odin

ENV HOME /mal
