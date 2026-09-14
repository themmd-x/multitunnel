#!/bin/bash

DIR="$1"
FILE="$2"

find "$DIR" -type f -name "$FILE" -delete
