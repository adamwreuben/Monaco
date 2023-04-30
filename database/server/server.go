package server

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var RethinkDbSession *rethinkdb.Session

func StartMonacoDBServer(router *gin.Engine, session *rethinkdb.Session) {
	RethinkDbSession = session
	router.Use(cors.Default())
	//CRUD operations
	router.POST("/add", addRowData)
	router.GET("/getTables", nil)
	router.GET("/getTable/{table}", getAllTableData)
	router.GET("/getProjects", nil)
	router.GET("/getRow/{table}/{rowId}", nil)
	router.GET("/deleteRow/{table}/{rowId}", nil)
}

func addRowData(*gin.Context) {
	var data map[string]interface{}
	_, err := rethinkdb.Table(data["table"].(string)).Insert(rethinkdb.Expr(data)).RunWrite(RethinkDbSession)
	if err != nil {
		log.Fatalf("Error inserting data: %s", err)
	}
	log.Printf("Data inserted successfully")
}

func getAllTableData(*gin.Context) {
	var table string
	var data []interface{}

	cursor, err := rethinkdb.Table(table).Run(RethinkDbSession)
	if err != nil {
		log.Fatalf("Error retrieving data: %s", err)
	}

	defer cursor.Close()

	if err := cursor.All(&data); err != nil {
		log.Fatalf("Error retrieving data as array: %s", err)
	}

	// Print the data
	log.Printf("%#v", data)

}
