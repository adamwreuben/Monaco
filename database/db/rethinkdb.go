package db

import (
	"fmt"

	"gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func StartRethinkDb() *rethinkdb.Session {
	session, err := rethinkdb.Connect(rethinkdb.ConnectOpts{
		Address:  "localhost",
		Database: "MonacoDb",
	})

	if err != nil {
		fmt.Println("Error while starting Monaco Database")
		panic(err)
	}
	fmt.Println("Monaco Database is running")

	return session
}
