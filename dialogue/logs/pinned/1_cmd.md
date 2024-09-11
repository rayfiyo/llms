```
./bin/dialogue \
-model1="Llama-3-Swallow-70B-Instruct-v0.1-Q8_0" \
-model2="EvoLLM-JP-v1-10B-f16" \
-head="SYSTEM 日本語で回答してください" \
"今日は１日自由に過ごせたとすると何をしていたか空想してみてください！" ;
```

```
./bin/dialogue -model=Llama-3-Swallow-70B-Instruct-v0.1-Q8_0 \
-head="SYSTEM 日本語で回答してください" \
-head1="あなたは奇数番目のユーザーです" \
-head2="あなたは偶数番目のユーザーです" \
-tail1="あなたは奇数番目のユーザーです" \
-tail2="あなたは偶数番目のユーザーです" \
"今日は１日自由に過ごせたとすると何をしていたか空想してみてください！" ;
```
