#!/bin/bash

for dir in versions/*/; do
    name="${dir%/}"
    name="${name##*/}"

    if [[ "$name" =~ ^v[0-9]+$ ]]; then
        echo "Processing $name..."
        find "$dir" -type f -exec sed -i "s/v12140/$name/g" {} +
    fi
done
