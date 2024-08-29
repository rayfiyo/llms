FROM debian:12.6-slim
COPY models/ /models/
RUN apt update -y &&\
    apt upgrade &&\
    apt install -y gcc g++ git make curl &&\
    git clone https://github.com/ggerganov/llama.cpp.git &&\
    cd llama.cpp &&\
    make -j &&\
    mv /models/* /llama.cpp
    # ./llama-cli -m 'Llama-3.1-70B-Japanese-Instruct-2407-Q8_0.gguf' -p 'あなたは日本語を話すアシスタントです' -n 1024 -c 1024 -cnv
