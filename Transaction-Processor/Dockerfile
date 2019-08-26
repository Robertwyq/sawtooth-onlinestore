# Copyright 2018 Intel Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# -----------------------------------------------------------------------------

FROM ubuntu:xenial

RUN echo "deb [arch=amd64] http://repo.sawtooth.me/ubuntu/ci xenial universe" >> /etc/apt/sources.list \
 && echo "deb http://archive.ubuntu.com/ubuntu bionic-backports universe" >> /etc/apt/sources.list \
 && (apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 8AA7AF1F1091A5FD \
   || apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 8AA7AF1F1091A5FD) \
 && apt-get update \
 && apt-get install -y \
    build-essential \
    git \
    golang-1.9-go \
    libssl-dev \
    libzmq5 \
    libzmq3-dev \
    openssl \
    python3 \
    python3-grpcio-tools=1.1.3-1 \
 && apt-get clean

RUN mkdir -p /project/onlinestore/Transaction-Processor
ENV GOPATH=/go:/go/src/github.com/hyperledger/sawtooth-sdk-go:/project/onlinestore/Transaction-Processor
ENV PATH=$PATH:/project/bin:/go/bin:/usr/lib/go-1.9/bin:/project/onlinestore/Transaction-Processor

EXPOSE 4004/tcp

WORKDIR /project/onlinestore/Transaction-Processor
COPY onlinestore /project/onlinestore/Transaction-Processor
CMD bash -C "onlinestore"
