# docker-go

docker goテンプレート

# コンテナ作成

```
$ docker-compose build
```

# コンテナ起動

docker-compose up -d

# コンテナのシェルに入りファイルを実行

```
$ docker-compose exec app /bin/bash

$ cd src

$ go run main.go
```