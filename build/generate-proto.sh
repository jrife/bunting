#!/bin/bash
set -ex

HOME=/home/me
BUNTING_ROOT=$HOME/bunting.io/bunting
GOGOPROTO_ROOT=$HOME/protobuf
cd $BUNTING_ROOT/api/buntingkvpb
protoc -I $HOME \
    -I=$GOGOPROTO_ROOT/protobuf/ \
    -I=$GOGOPROTO_ROOT \
    -I=. \
    --gogofaster_out=paths=source_relative:. \
    buntingkv.proto
protoc -I $HOME \
    -I=$GOGOPROTO_ROOT/protobuf/ \
    -I=$GOGOPROTO_ROOT \
    -I=. \
    --gogofaster_out=plugins=grpc:. \
    rpc.proto
cd $BUNTING_ROOT/api/interceptorpb
protoc -I $HOME \
    -I=$GOGOPROTO_ROOT/protobuf/ \
    -I=$GOGOPROTO_ROOT \
    -I=. \
    --gogofaster_out=plugins=grpc:. \
    rpc.proto
cd $BUNTING_ROOT/api/observerpb
protoc -I $HOME \
    -I=$GOGOPROTO_ROOT/protobuf/ \
    -I=$GOGOPROTO_ROOT \
    -I=. \
    --gogofaster_out=plugins=grpc:. \
    rpc.proto