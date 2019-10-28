package main

import (
	"github.com/tonyalaribe/adapters/sql"
)

func main() {
	db := sql.New()
	defer db.Cleanup()

}
