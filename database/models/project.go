package models

import (
	"encoding/json"
	"fmt"
)

type MonacoProject struct {
	ProjectName string  `json:"project"`
	ProjectId   string  `json:"projectId"`
	Status      bool    `json:"status"`
	Tables      []Table `json:"tables"`
}

type Table struct {
	TableName string   `json:"table"`
	Index     int      `json:"index"`
	TableId   []string `json:"tableId"`
}

func CreateProject(name string) []byte {
	jsonProject, err := json.Marshal(MonacoProject{ProjectName: name, Status: true})
	if err != nil {
		fmt.Println(err)
	}
	return jsonProject
}
