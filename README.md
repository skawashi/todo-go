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
goのコンテナに入る
```
docker exec -it go-todo bash
```
次のコマンドを実行して、同じような出力がされればok
```
$ go run main.py

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

次のコマンドで`mysql`と接続する。(passwordはREADME通りに.envを設定していればroot_passowrd)
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
+-------------------------+
```
