package main

import (
	"fmt"
	"log"
	"os"

	"app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

// DBを起動させる
func dbInit() *gorm.DB {
	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
	dsn := fmt.Sprintf(`%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=True`,
            os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))
	// DBへの接続を行う
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// エラーが発生した場合、エラー内容を表示
	if err != nil {
		log.Fatal(err)
	}
	// 接続に成功した場合、「db connected!!」と表示する
	fmt.Println("db connected!!")
	return db
}

func main() {
	// DB起動
	db := dbInit()
	// Userテーブル作成
	// db.AutoMigrate(
	// 	&User{}
	// )
	db.AutoMigrate(
		&model.User{},
		&model.Todo{},
	)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":"Hello World",
		})
	})
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}


