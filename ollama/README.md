# cmd

```bash
make start
docker exec -it ollama-container ollama run モデル

# Open WebUIのセットアップ
docker run -d -p 8080:8080 --add-host=host.docker.internal:host-gateway -v open-webui:/app/backend/data --name open-webui --restart always ghcr.io/open-webui/open-webui:main
```

| Model              | Parameters | Size  | Download            |
| ------------------ | ---------- | ----- | ------------------- |
| Llama 3            | 8B         | 4.7GB | `llama3`            |
| Llama 3            | 70B        | 40GB  | `llama3:70b`        |
| Phi-3              | 3.8B       | 2.3GB | `phi3`              |
| Mistral            | 7B         | 4.1GB | `mistral`           |
| Neural Chat        | 7B         | 4.1GB | `neural-chat`       |
| Starling           | 7B         | 4.1GB | `starling-lm`       |
| Code Llama         | 7B         | 3.8GB | `codellama`         |
| Llama 2 Uncensored | 7B         | 3.8GB | `llama2-uncensored` |
| LLaVA              | 7B         | 4.5GB | `llava`             |
| Gemma              | 2B         | 1.4GB | `gemma:2b`          |
| Gemma              | 7B         | 4.8GB | `gemma:7b`          |
| Solar              | 10.7B      | 6.1GB | `solar`             |

# Links

## Thanks

- [LLaMa-3をAPIサーバーのように使う方法 #LLM - Qiita](https://qiita.com/tasuku-revol/items/6a287fb69ce4a423dbe0)
  - モデルの表
- [「よーしパパ、Ollama で Llama-3-ELYZA-JP-8B 動かしちゃうぞー」 #LLM - Qiita](https://qiita.com/s3kzk/items/3cebb8d306fb46cabe9f)
- [Dockerを使ってOllamaとOpen WebUIでllama3を動かす](https://zenn.dev/misora/articles/1037a94c53a5f0)

## notes

- [いちばんやさしいローカル LLM｜ぬこぬこ](https://note.com/schroneko/n/n8b1a5bbc740b)
- [N番煎じですがローカルLLM(ollama)とRAG(Dify)を試してみた｜桑名市スマートシティ推進課](http://kuwana-city.note.jp/n/nd64b8dcfb830)