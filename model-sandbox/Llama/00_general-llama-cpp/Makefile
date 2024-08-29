.PHONY: start build run clean-f clean rmi

CONTAINER_NAME = ELYZA-japanese-Llama-2-7b-fast-instruct-q8_0
IMAGE_NAME = llm-cpp

MODEL_EXIST = $(shell ls models/ 2>/dev/null)

start:
	@make --no-print-directory model-exist
	@make --no-print-directory build
	@make --no-print-directory run

model-exist:
ifeq ($(MODEL_EXIST), )
	@echo "error: models/ is empty"
	@exit 1
endif

build:
	docker build -t $(IMAGE_NAME) .

run:
	docker run -it --name $(CONTAINER_NAME) $(IMAGE_NAME)

clean-f:
	@make --no-print-directory clean
	@make --no-print-directory rmi

clean:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)

rmi:
	docker rmi $(IMAGE_NAME)
