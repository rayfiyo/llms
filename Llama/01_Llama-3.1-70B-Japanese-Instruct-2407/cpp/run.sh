#!/usr/bin/bash

docker build -t llm-cpp . && docker run -it --name llama-3.1-70B-Japanese-Instruct-2407-gguf-cpp llm-cpp
