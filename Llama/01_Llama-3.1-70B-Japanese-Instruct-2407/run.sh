#!/usr/bin/bash

docker build -t llm . && docker run -it --name llama-3.1-70B-Japanese-Instruct-2407-gguf llm
