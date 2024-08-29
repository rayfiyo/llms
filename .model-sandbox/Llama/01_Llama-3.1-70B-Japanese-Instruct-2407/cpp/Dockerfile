FROM debian:12.6-slim
RUN apt update -y &&\
    apt upgrade &&\
    apt install -y gcc g++ git make curl &&\
    git clone https://github.com/ggerganov/llama.cpp.git &&\
    cd llama.cpp &&\
    make -j &&\
    curl -L https://huggingface.co/mmnga/Llama-3.1-70B-Japanese-Instruct-2407-gguf/resolve/main/Llama-3.1-70B-Japanese-Instruct-2407-Q8_0.gguf.a?download=true -O "Llama-3.1-70B-Japanese-Instruct-2407-Q8_0.gguf.a" &&\
    curl -L https://huggingface.co/mmnga/Llama-3.1-70B-Japanese-Instruct-2407-gguf/resolve/main/Llama-3.1-70B-Japanese-Instruct-2407-Q8_0.gguf.b?download=true -O "Llama-3.1-70B-Japanese-Instruct-2407-Q8_0.gguf.b" &&\
    cat Llama-3.1-70B-Japanese-Instruct-2407-Q8_0.gguf.a Llama-3.1-70B-Japanese-Instruct-2407-Q8_0.gguf.b > Llama-3.1-70B-Japanese-Instruct-2407-Q8_0.gguf 
    # ./llama-cli -m 'Llama-3.1-70B-Japanese-Instruct-2407-Q8_0.gguf' -p 'あなたは日本語を話すアシスタントです' -n 1024 -c 1024 -cnv
    # huggingface-cli download mmnga/Llama-3.1-70B-Japanese-Instruct-2407-Q8_0.gguf.a mmnga/Llama-3.1-70B-Japanese-Instruct-2407-Q8_0.gguf.b --local-dir . &&\
