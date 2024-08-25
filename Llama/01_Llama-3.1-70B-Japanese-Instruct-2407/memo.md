# 成功したコマンド

```
RUN     apt update -y &&\
        apt upgrade &&\
        apt install -y python3 python3.11-venv gcc g++ git &&\
        python3 -m venv llama3.1 &&\
        cd llama3.1 &&\
        source bin/activate &&\
        pip install torch transformers llama-cpp-python &&\
        CMAKE_ARGS="-DLLAMA_CUBLAS=on -DLLAVA_BUILD=off" pip install llama-cpp-python &&\
        CMAKE_ARGS="-DGGML_CUDA=on -DLLAVA_BUILD=off" pip install llama-cpp-python --upgrade
COPY ./query4llama-cpp.py "/llama3.1/"
```

```
        echo "import sys
import argparse
from huggingface_hub import hf_hub_download
from llama_cpp import Llama, llama_chat_format
import time

def parse_arguments():
    parser = argparse.ArgumentParser()
    parser.add_argument(\"--model-path\", required=True)
    parser.add_argument(\"--ggml-model-path\", required=True)
    parser.add_argument(\"--ggml-model-file\", required=True)
    parser.add_argument(\"--max-tokens\", type=int, default=256)
    parser.add_argument(\"--n-ctx\", type=int, default=2048)
    parser.add_argument(\"--n-threads\", type=int, default=1)
    parser.add_argument(\"--n-gpu-layers\", type=int, default=-1)
    return parser.parse_args()

def setup_model(args):
    ggml_model_path = hf_hub_download(args.ggml_model_path, filename=args.ggml_model_file)
    chat_formatter = llama_chat_format.hf_autotokenizer_to_chat_formatter(args.model_path)
    chat_handler = llama_chat_format.hf_autotokenizer_to_chat_completion_handler(args.model_path)
    return Llama(model_path=ggml_model_path, chat_handler=chat_handler, n_ctx=args.n_ctx,
                 n_threads=args.n_threads, n_gpu_layers=args.n_gpu_layers), chat_formatter

def generate_response(model, chat_formatter, prompt, history=None):
    start = time.process_time()
    messages = [{\"role\": \"system\", \"content\": \"あなたは誠実で優秀な日本人のアシスタントです。\"}]
    if history:
        messages.extend(history)
    messages.append({\"role\": \"user\", \"content\": prompt})
    formatted_prompt = chat_formatter(messages=messages)
    outputs = model.create_chat_completion(messages=messages, temperature=0.8, top_p=0.95,
                                           top_k=40, max_tokens=args.max_tokens, repeat_penalty=1.1)
    response = outputs[\"choices\"][0][\"message\"][\"content\"]
    print(response)
    end = time.process_time()
    print(f\"prompt tokens = {outputs['usage']['prompt_tokens']}\")
    print(f\"output tokens = {outputs['usage']['completion_tokens']} ({outputs['usage']['completion_tokens'] / (end - start):.2f} [tps])\")
    print(f\"   total time = {end - start:.2f} [s]\")
    return response

if __name__ == \"__main__\":
    args = parse_arguments()
    model, chat_formatter = setup_model(args)
    history = []
    history.append({\"role\": \"assistant\", \"content\": generate_response(model, chat_formatter, \"ソードアート・オンラインとは何ですか？\")})
    history.append({\"role\": \"assistant\", \"content\": generate_response(model, chat_formatter, \"続きを教えてください\", history)})
" > query4llama-cpp.py
        CUDA_VISIBLE_DEVICES=0,1 python query4llama-cpp.py \
            --model-path cyberagent/Llama-3.1-70B-Japanese-Instruct-2407 \
            --ggml-model-path mmnga/Llama-3.1-70B-Japanese-Instruct-2407-gguf \
            --ggml-model-file Llama-3.1-70B-Japanese-Instruct-2407-IQ4_XS.gguf
```

# メモ
