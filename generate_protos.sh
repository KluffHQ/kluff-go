#!/bin/bash

get_files() {
    local dir="$1"
    local files=()

    for file in "$dir"/*; do
        if [ -f "$file" ]; then
            files+=("$(basename "$file")")
        fi
    done

    # Return the array by printing its elements
    echo "${files[@]}"
}


while IFS= read -r folder; do
    # Skip empty lines or lines starting with #
    if [[ -n "$folder" && ! "$folder" =~ ^#.* ]]; then
        echo "Generating code for folder: $folder"
        
        cd "data"
        mkdir "$folder"

        cd ..

        # Navigate to the proto folder
        cd "proto/$folder"
 
        echo "Current directory: $(pwd)"

        my_files=( $(get_files ".") )
        # List files in the current directory

        for file in "${my_files[@]}"; do
            echo "$file"
            # Run the protoc command
            protoc --go_out=../../data/"$folder" --go-grpc_out=../../data/"$folder" --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $file
        done

        
        # Navigate back to the root directory
        cd ../..
        echo "Back to: $(pwd)"
    fi
done < proto_folders.txt
