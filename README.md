# 画像圧縮WebAPIサンプル

負荷試験などに使用できる画像圧縮APIサーバーです。
てきとうな画像を、`/v1/compress`にPOSTすると、webpに変換・圧縮します。

## 使い方

```shell
docker run -p 1323:1323 ghcr.io/irumaru/webp-compress-api:v2025.1216.3
```

curlでとりあえずテストする
```
curl -X POST -F image=@./test/testdata/1.4MB.jpg -o test/testdata/out.webp http://127.0.0.1:1323/v1/compress
```

k6でたくさんリクエストする
```shell
k6 run test/loadtest.js
```
