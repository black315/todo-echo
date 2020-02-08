package model

import (
    "../database"
)

type Todo struct {
    ID uint `gorm:"primary_key" json:"id"`
    Name string `json:"name"`
}

type Todos struct {
    Todos []Todo `json:"todos"`
}

var todo Todo
var todos Todos

func GetTodos() (todos Todos) {
    db := database.GetDBConnection()
    db.AutoMigrate(&Todo{})

    db.Find(&todos)
    return
}
