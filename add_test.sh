#!/bin/bash

for file in ./internal/models/*_test.go; do 
    # filename=$(basename "$file")
    # file_no_ext="${filename%*.go}"
    # new_file="${file_no_ext}_test.go"
    # echo "$new_file"
    # touch "./internal/models/${new_file}"
    echo "package models" >> "$file"
done

