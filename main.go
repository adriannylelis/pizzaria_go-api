package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas []models.Pizza

func main() {
	loadPizzas()

	router := gin.Default()
	router.GET("/pizzas", getPizzas )
	router.POST("/pizzas", postPizzas )
	router.GET("/pizzas/:id", getPizzaByID )

	nomePizzaria := "***** Pizzaria da Dri *****"
	insta, tel := "instagram.com/pizzariadadri", "9999-9999"

	fmt.Printf("%s (instagram: %s) - Telefone: %s", nomePizzaria, insta, tel)

	router.Run()
}

func getPizzas(c *gin.Context){
	
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})	
}

func postPizzas(c *gin.Context){
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)
	savePizza()
	c.JSON(201, newPizza)
}

func getPizzaByID(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	for _, p := range pizzas {
		if p.ID == id {
			c.JSON(200, p)
			return
		}
	}
	c.JSON(404, gin.H{"error": "pizza não encontrada"})
}

func loadPizzas(){
	file, err := os.Open("data/pizza.json")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo de pizzas", err)
		return
	}
	defer file.Close()

	decorder := json.NewDecoder(file)
	if err := decorder.Decode(&pizzas); err != nil {
		fmt.Println("Erro ao ler arquivo de pizzas", err)
		return
	}
}

func savePizza(){
	file, err := os.Create("data/pizza.json")
	if err != nil {
		fmt.Println("Erro ao criar arquivo de pizzas", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {
		fmt.Println("Erro ao salvar pizzas", err)
		return
	}
}