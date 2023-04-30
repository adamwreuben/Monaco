package main

import (
	"monaco/database/db"
	"monaco/database/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//Starting Db connections
	session := db.StartRethinkDb()

	//Start Monaco Server
	server.StartMonacoDBServer(r, session)

	r.Run() // default PORT 8080
}
