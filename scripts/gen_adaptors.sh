#!/bin/bash

SOURCE="$1"
PARENT_DIR="$2"

for dir in "$PARENT_DIR"/*/; do
    cp "$SOURCE" "$dir"
done
