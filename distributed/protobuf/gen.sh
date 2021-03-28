#!/usr/bin/env bash

protoc --go_out=plugins=grpc:. broadcast_protobuf.proto
protoc --go_out=plugins=grpc:. connection_protobuf.proto
protoc --go_out=plugins=grpc:. group_protobuf.proto
protoc --go_out=plugins=grpc:. user_protobuf.proto