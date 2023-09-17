#!/bin/bash

while IFS= read -r folder; do
    # Skip empty lines or lines starting with #
    if [[ -n "$folder" && ! "$folder" =~ ^#.* ]]; then
        echo "Generating code for folder: $folder"

        cd "pkg/api/"
        mkdir "$folder"

        cd ../..

        # Navigate to the proto folder
        cd "api/proto/$folder"
        echo "Current directory: $(pwd)"

        # List files in the current directory
        ls

        # Run the protoc command
        protoc --go_out=../../../pkg/api/"$folder" --go-grpc_out=../../../pkg/api/"$folder" --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $folder.proto

        # Navigate back to the root directory
        cd ../../..
        echo "Back to: $(pwd)"
    fi
done < proto_folders.txt
