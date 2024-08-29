.PHONY: start build run clean-f clean rmi

CONTAINER_NAME = llama-3.1-70B-Japanese-Instruct-2407-gguf-cpp
IMAGE_NAME = llm-cpp

MODEL_A_NAME = model/Llama-3.1-70B-Japanese-Instruct-2407-Q8_0.gguf.a
MODEL_A_EXIST = $(shell ls $(MODEL_A_NAME) 2>/dev/null)

MODEL_B_NAME = model/Llama-3.1-70B-Japanese-Instruct-2407-Q8_0.gguf.b
MODEL_B_EXIST = $(shell ls $(MODEL_B_NAME) 2>/dev/null)

start:
	@make model-exist
	@make build
	@make run

model-exist:
ifeq ($(MODEL_A_EXIST), $(MODEL_A_NAME))
	echo hoge
else
	curl -L https://huggingface.co/mmnga/Llama-3.1-70B-Japanese-Instruct-2407-gguf/resolve/main/Llama-3.1-70B-Japanese-Instruct-2407-Q8_0.gguf.a?download=true -o $(MODEL_A_NAME)
endif
ifeq ($(MODEL_B_EXIST), $(MODEL_B_NAME))
else
    curl -L https://huggingface.co/mmnga/Llama-3.1-70B-Japanese-Instruct-2407-gguf/resolve/main/Llama-3.1-70B-Japanese-Instruct-2407-Q8_0.gguf.b?download=true -o $(MODEL_B_NAME)
endif

build:
	docker build -t $(IMAGE_NAME) .

run:
	docker run -it --name $(CONTAINER_NAME) $(IMAGE_NAME)

clean-f:
	@make clean
	@make rmi

clean:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)

rmi:
	docker rmi $(IMAGE_NAME)
