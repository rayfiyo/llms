FROM debian:12.6-slim
RUN     apt update -y &&\
        apt upgrade &&\
        apt install -y python3 python3.11-venv gcc g++ git &&\
        python3 -m venv llama3.1 &&\
        cd llama3.1 &&\
        source bin/activate &&\
        pip install torch transformers llama-cpp-python &&\
        CMAKE_ARGS="-DLLAMA_CUBLAS=on -DLLAVA_BUILD=off" pip install llama-cpp-python &&\
        CMAKE_ARGS="-DGGML_CUDA=on -DLLAVA_BUILD=off" pip install llama-cpp-python --upgrade
