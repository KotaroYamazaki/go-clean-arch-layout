package main

import (
	"log"

	"github.com/KotaroYamazaki/go-clean-arch-layout/config/env"
	db "github.com/KotaroYamazaki/go-clean-arch-layout/internal/infra/db"
)

func init() {
	env.Init()

	_, err := db.NewDBCon("hoge")
	if err != nil {
		log.Fatal("failed to connect to database")
	}
}

func main() {
	println("Hello, world.")
}
