# note

```bash
# make build

make start

make stop

# make clean

# Open WebUIのセットアップ
docker run -d -p 8080:8080 --add-host=host.docker.internal:host-gateway -v open-webui:/app/backend/data --name open-webui --restart always ghcr.io/open-webui/open-webui:main
```

# init

- .gguf をダウンロードする
- モデルファイルを作成する（参考: modelfiles/）

# Links

## Thanks

- [LLaMa-3 を API サーバーのように使う方法 #LLM - Qiita](https://qiita.com/tasuku-revol/items/6a287fb69ce4a423dbe0)
  - モデルの情報（表）あり
- [「よーしパパ、Ollama で Llama-3-ELYZA-JP-8B 動かしちゃうぞー」 #LLM - Qiita](https://qiita.com/s3kzk/items/3cebb8d306fb46cabe9f)
- [Docker を使って Ollama と Open WebUI で llama3 を動かす](https://zenn.dev/misora/articles/1037a94c53a5f0)
- [Llama-3-Swallow-70B を使ってみた（Mac + Ollama）](https://zenn.dev/robustonian/articles/llama3_swallow_70b#ollama-modelfile%E3%81%AE%E4%BD%9C%E6%88%90)

## notes

- [いちばんやさしいローカル LLM ｜ぬこぬこ](https://note.com/schroneko/n/n8b1a5bbc740b)
- [N 番煎じですがローカル LLM(ollama)と RAG(Dify)を試してみた｜桑名市スマートシティ推進課](http://kuwana-city.note.jp/n/nd64b8dcfb830)
