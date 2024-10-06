FROM ollama/ollama:0.3.8
ENV OLLAMA_MAX_LOADED_MODELS 4
RUN sed -i '$aalias run="ollama run --keepalive 1h"' /root/.bashrc
COPY modelfiles .
COPY models/ .
