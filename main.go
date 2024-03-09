package main

import (
	"fmt"
	"os"

	"github.com/conneroisu/dbpu/pkg/databases"
)

func main() {
	val, err := databases.CreateDatabase("conner", os.Getenv("TURSO_TOKEN"), "newerer-db", "dbpu")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
