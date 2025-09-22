package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// 環境変数読み込み
	if err := godotenv.Load(); err != nil {
		log.Println("環境変数ファイルが見つかりませんでした")
	}

	log.Println("Migration tool - まもなく提供予定")
	// TODO: マイグレーション機能を実装
}