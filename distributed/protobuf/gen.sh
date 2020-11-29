#!/usr/bin/env bash

protoc --go_out=plugins=grpc:. gosh_protobuf.proto
