.PHONY: start build run clean-all clean rmi

CONTAINER_NAME = ollama-container
IMAGE_NAME = ollama-image

start:
	@make --no-print-directory build
	@make --no-print-directory run

build:
	docker build -t $(IMAGE_NAME) .

run:
	docker run -p 11434:11434 --name $(CONTAINER_NAME) $(IMAGE_NAME)

clean-all:
	@make --no-print-directory clean
	@make --no-print-directory rmi

clean:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)

rmi:
	docker rmi $(IMAGE_NAME)
