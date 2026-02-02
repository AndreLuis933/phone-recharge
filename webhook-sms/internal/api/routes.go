package api

import (
	"encoding/json"
	"httpsms-webhook/internal/logger"
	"httpsms-webhook/internal/models"
	"httpsms-webhook/internal/services"

	"github.com/gin-gonic/gin"
)

func WebhookHandler(c *gin.Context) {
	eventType := c.GetHeader("X-Event-Type")

	var cloudEvent models.CloudEvent
	if err := c.ShouldBindJSON(&cloudEvent); err != nil {
		logger.Log.Error("Erro ao parsear CloudEvent", "error", err)
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	switch eventType {
	case EventMessageReceived:
		var msgData models.MessageData
		if err := json.Unmarshal(cloudEvent.Data, &msgData); err != nil {
			logger.Log.Error("Erro ao parsear dados", "error", err, "type", eventType)
			c.JSON(200, gin.H{"status": "error"})
			return
		}

		// Agora você tem acesso direto aos campos
		logger.Log.Info("Processando mensagem",
			"de", msgData.Contact,
			"conteudo", msgData.Content,
			"id", msgData.ID,
		)
		if msgData.Contact == "321"{
			services.ProcessVivoSMS(msgData)
		}

	case EventHeartbeatOffline, EventHeartbeatOnline:
		var hbData models.HeartbeatData
		if err := json.Unmarshal(cloudEvent.Data, &hbData); err != nil {
			logger.Log.Error("Erro ao parsear heartbeat", "error", err, "type", eventType)
			c.JSON(200, gin.H{"status": "error"})
			return
		}

		// Log do heartbeat
		prettyData, _ := json.MarshalIndent(hbData, "", "  ")
		logger.Log.Info("Heartbeat recebido",
			"type", eventType,
			"status", hbData.Status,
			"data", string(prettyData),
		)

	default:
		logger.Log.Warn("Tipo de evento desconhecido", "type", eventType)
		c.JSON(200, gin.H{"status": "ok"})
		return
	}

	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "Webhook processado com sucesso",
	})
}

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"service": "httpsms-webhook",
	})
}
