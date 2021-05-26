package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	urlExample := "postgres://drakorod:s3cr3tH4cks@localhost:5432/postgres"
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, urlExample)

	if err != nil {
		fmt.Println("Unable to connect to database")
		os.Exit(1)
	}

	defer conn.Close(ctx)

	sql := "SELECT datname, u.usename AS datowner FROM pg_database db INNER JOIN pg_user u ON u.usesysid = db.datdba"
	jsonSQL := fmt.Sprintf("SELECT array_to_json(array_agg(row_to_json(t))) FROM (%s) t", sql)

	rows, err := conn.Query(ctx, jsonSQL)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var result interface{}
	for rows.Next() {
		values, _ := rows.Values()
		for _, v := range values {
			result = v
		}
	}

	json, err := json.Marshal(result)

	fmt.Println("JSON Result::> ", string(json))

}
