#!/bin/bash

SOURCE="./versions/v12140/adapter.go"
PARENT_DIR="./versions/"

for dir in "$PARENT_DIR"/*/; do
    cp "$SOURCE" "$dir"
done
