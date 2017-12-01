#!/usr/bin/env bash

# TODO: add images check and create image before run
docker run -d -p 8080:8080 yclite:0.1 
echo "http://localhost:8080/list/1"