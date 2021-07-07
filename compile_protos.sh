
#!/bin/bash

# it maybe extend to cover multiple directories for respecting hierarchy
dir="."
proto_file="bsp_transaction.proto blockservice.proto"
protoc --proto_path="$dir" --go_out=plugins=grpc,paths=source_relative:"$dir" "$dir"/$proto_file
