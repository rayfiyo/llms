FROM ollama/ollama:0.3.8
COPY modelfiles/ .
COPY models/ .
CMD ["/bin/sh", "-c", "export 'OLLAMA_MAX_LOADED_MODELS=4'"]
