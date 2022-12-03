## 開発環境の立て方

- docker-compose upする

```sh
docker-compose build
docker-compose up -d
```

- migrationする

```bash
curl http://localhost:5018
```

- 各サービスがlocalhostのいろんなportで立ってる (`docker-compose.yml`参照)
