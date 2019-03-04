package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	a := App{}
	// You need to set your Username and Password here - DB name is api_task
	a.Initialize("root", "s", "api_task")

	a.Run(":8080")
}
