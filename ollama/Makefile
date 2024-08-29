.PHONY: build run start		clean rmi clean-all

CONTAINER_NAME = ollama-container
IMAGE_NAME = ollama-image

#

build:
	docker build -t $(IMAGE_NAME) --quiet .

run:
	docker run -d -p 11434:11434 --name $(CONTAINER_NAME) $(IMAGE_NAME)

start:
	@make --no-print-directory build
	@make --no-print-directory run

#

clean:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)

rmi:
	docker rmi $(IMAGE_NAME)

clean-all:
	@make --no-print-directory clean
	@make --no-print-directory rmi
