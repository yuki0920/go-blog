package main

import (
	"log"
	"os"
	"yuki0920/go-blog/domain/model"
	"yuki0920/go-blog/infra"

	_ "github.com/go-sql-driver/mysql" // MySQLのドライバーを使う
)

func main() {
	db, err := infra.ConnectDB()
	if err != nil {
		log.Fatalf("db connection failed %v", err)
	} else {
		log.Println("db connection established")
	}

	infra.SetDB(db)

	user := &model.User{}

	user.Name = os.Args[1]

	password := os.Args[2]
	if err := user.SetPassword(password); err != nil {
		log.Fatalf("Set password failed %v", err)
	}

	if _, err := infra.UserCreate(user); err != nil {
		log.Fatalf("User create failed %v", err)
	}

	log.Println("user registered successfully")
}