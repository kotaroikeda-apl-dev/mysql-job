## **概要**

## **実行方法**

```sh
docker compose up -d # データベース起動
go run cmd/regist/main.go # ジョブログ登録
docker compose down # データベース停止
```

## **学習ポイント**

1. **`sql.Open()`** でデータベース接続を作成し、**`defer db.Close()`** で接続を適切にクローズすることで、リソースリークを防ぐ。
2. **`db.Exec(query, ?, ?)`** を使うことで、プレースホルダ (**`?`**) を利用し、SQL インジェクションを防ぎながらデータを安全に挿入できる。

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
