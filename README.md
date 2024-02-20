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

go.mod 作成
Goのモジュール管理ファイルを作成する
`golang/app`ディレクトリに移動して以下のコマンドを実行
```
go mod init app
```

ビルドの実行
```
cd ../../ && docker-compose build
```
※`network golang_test_network declared as external, but could not be found`と出た場合は、以下も実行した後に、dockerを再度ビルドする。
```
docker network create golang_test_network
```

以下のコマンドを実行
```
docker-compose run --rm go air init
docker-compose run --rm go go mod tidy
```
※こちらはマストの作業ではないが、`golang`ディレクトリに移動し,`tree`コマンドで以下の構成になって入ればOK
```
.
├── Dockerfile
└── app
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── model
    │   ├── todo.go
    │   └── user.go
    └── tmp
        └── main
```

コンテナの立ち上げ
```
docker-compose up
```

別のターミナルを立ち上げて以下のコマンドを実行してメッセージが帰ってくればok
```
curl -X GET "http://0.0.0.0:8080/"
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
