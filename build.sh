#!/bin/bash

OUT_DIR="./out/"
mkdir -p $OUT_DIR
go build -o $OUT_DIR/iwan ./src/ && echo Build completed!
