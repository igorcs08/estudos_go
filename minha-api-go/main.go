package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Tarefa struct {
	Id int `json:"id"`
	Titulo string `json:"titulo"`
	Feita bool`json:"feita"`
}

var tarefas = []Tarefa{
	{Id: 1, Titulo: "Aprender Go", Feita: false},
	{Id: 2, Titulo: "Construir uma API", Feita: false},
}

func main() {
	r := gin.Default()

	// Rota GET para listar todas as tarefas
	r.GET("/tarefas", func(c *gin.Context) {
		c.JSON(http.StatusOK, tarefas)
	})

	// Rota POST para adicionar uma nova tarefa
	r.POST("/tarefas", func(c *gin.Context) {
		var novaTarefa Tarefa

		if err := c.ShouldBindJSON(&novaTarefa); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
			return
		}

		novaTarefa.Id = len(tarefas) + 1
		tarefas = append(tarefas, novaTarefa)
		c.JSON(http.StatusCreated, novaTarefa)
	})

	r.Run()
}