.PHONY: build run exec init start	stop clean rmi clean-all

CONTAINER = ollama-container
IMAGE = ollama-image
MODEL = Llama-3-Swallow-70B-Instruct-v0.1-Q8_0
# MODEL = llama3:70b

init:
	@make --no-print-directory build
	@make --no-print-directory run

start:
	@make --no-print-directory build
	@make --no-print-directory run
	@make --no-print-directory exec

clean-all:
	@make --no-print-directory clean
	@make --no-print-directory rmi

#

build:
	docker build -t $(IMAGE) --quiet .

run:
	docker run -d -p 11434:11434 --name $(CONTAINER) $(IMAGE)

exec:
	docker exec -it $(CONTAINER) bash -c "ollama create $(MODEL) -f /modelfiles/$(MODEL)/Modelfile && ollama run $(MODEL)"
	# docker exec -it $(CONTAINER) ollama run $(MODEL)

stop:
	docker stop $(CONTAINER)

clean:
	@make --no-print-directory stop
	docker rm $(CONTAINER)

rmi:
	docker rmi $(IMAGE)
