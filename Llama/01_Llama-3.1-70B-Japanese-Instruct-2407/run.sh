#!/usr/bin/bash

docker build -t myapp . && docker run -it --name mycontainer myapp
