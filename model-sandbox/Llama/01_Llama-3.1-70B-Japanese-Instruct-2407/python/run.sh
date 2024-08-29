#!/usr/bin/bash

echo ". /llama3.1/bin/activate で環境の読み込み？"
docker build -t llm-python . && docker run -it --name llama-3.1-70B-Japanese-Instruct-2407-gguf-python llm-python
