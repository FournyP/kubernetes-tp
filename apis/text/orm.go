package main

import (
	"context"
	"log"
	"os"

	"github.com/FournyP/kubernetes-tp/apis/text/ent"

	_ "github.com/go-sql-driver/mysql"
)

func InitOrm() *ent.Client {
	client, err := ent.Open("mysql", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatalf("failed opening connection to %v: %v", os.Getenv("DATABASE_TYPE"), err)
	}

	if err = client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
