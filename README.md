プロジェクトをクローン
```
git clone git@github.com:skawashi/todo-go.git
```

プロジェクトディレクトリへ移動
```
cd todo-go
```

`mysql`ディレクトリの直下に`.env`を以下の内容で作成
```
MYSQL_DATABASE=test_database
MYSQL_USER=test_user
MYSQL_PASSWORD=password
MYSQL_ROOT_PASSWORD=root_password
```

docker立ち上げ
```
docker-compose up -d --build
```
※`network golang_test_network declared as external, but could not be found`と出た場合は、以下も実行した後に、dockerを再度立ち上げる。
```
docker network create golang_test_network
```

goのコンテナに入る
```
docker exec -it go-todo bash
```
`main.go`を実行。
```
go run main.go
```
以下のような出力になればOK。
```
go: downloading gorm.io/driver/mysql v1.5.2
go: downloading gorm.io/gorm v1.25.5
go: downloading github.com/go-sql-driver/mysql v1.7.0
go: downloading github.com/jinzhu/now v1.1.5
go: downloading github.com/jinzhu/inflection v1.0.0
db connected!!
```

goのコンテナを抜けて、mysqlのコンテナに入る
```
docker exec -it db-todo bash
```

mysqlと接続する。(passwordはREADME通りに.envを設定していれば`root_passowrd`)
```
mysql -u root -p
```

mysqlに接続できたら
```
use test_database;
```
を実行し、次に
```
show tables;
```
を実行する。次のように出力されたら、正しくテーブルが作成されている。
```
+-------------------------+
| Tables_in_test_database |
+-------------------------+
| users                   |
| todos                   |
+-------------------------+
```
