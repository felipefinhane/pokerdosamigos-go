package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var Router *gin.Engine

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	// Crie um novo bot com o token de acesso fornecido pelo BotFather
	bot, err := tgbotapi.NewBotAPI("6284838752:AAHl1EynXMoF2cPSzkh298anShH7I6bd2JQ")
	if err != nil {
		log.Panic(err)
	}

	// Crie uma rota POST para o caminho /webhook
	r.POST("/webhook", func(c *gin.Context) {
		// Decodifique o JSON da solicitação para uma estrutura de atualização do bot do Telegram
		var update tgbotapi.Update
		if err := json.NewDecoder(c.Request.Body).Decode(&update); err != nil {
			log.Printf("Error decoding request body: %s", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Verifique se a atualização recebida é uma mensagem
		if update.Message == nil {
			c.Status(http.StatusOK)
			return
		}

		// Responda à mensagem do usuário com uma mensagem de saudação
		reply := "Olá, " + update.Message.From.UserName + "! Eu sou um bot criado em Go!"
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)

		c.Status(http.StatusOK)
	})

	r.Run()
}
