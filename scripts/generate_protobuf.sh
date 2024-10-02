#!/bin/bash

protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     api/proto/task.proto
rm api/gen/*.pb.go
mv api/proto/*.pb.go api/gen/