package api

import (
	"encoding/json"
	"httpsms-webhook/internal/models"
	"io"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func WebhookHandler(c *gin.Context) {
	log.Println("\n" + strings.Repeat("=", 70))
	log.Println("📨 NOVO WEBHOOK RECEBIDO")
	log.Println(strings.Repeat("=", 70))

	// 1. HEADERS
	// O HTTPSMS envia alguns headers importantes
	log.Println("\n📋 HEADERS:")
	log.Println("---")

	eventType := c.GetHeader("X-Event-Type")
	contentType := c.GetHeader("Content-Type")

	log.Printf("  X-Event-Type: %s", eventType)
	log.Printf("  Content-Type: %s", contentType)
	log.Printf("  User-Agent: %s", c.GetHeader("User-Agent"))

	// 2. BODY (CloudEvent)
	// Ler o corpo da requisição
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("❌ Erro ao ler body: %v\n", err)
		c.JSON(400, gin.H{"error": "Error reading body"})
		return
	}

	// 3. LOGAR O JSON CRU
	// Vamos logar o JSON exatamente como veio
	log.Println("\n📦 PAYLOAD RAW (JSON):")
	log.Println("---")

	// Formatar o JSON para ficar mais legível
	var rawJSON map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &rawJSON); err == nil {
		prettyJSON, _ := json.MarshalIndent(rawJSON, "", "  ")
		log.Println(string(prettyJSON))
	} else {
		log.Println(string(bodyBytes))
	}

	// 4. PARSEAR O CLOUDEVENT
	var cloudEvent models.CloudEvent
	if err := json.Unmarshal(bodyBytes, &cloudEvent); err != nil {
		log.Printf("❌ Erro ao parsear CloudEvent: %v\n", err)
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	// 5. LOGAR INFORMAÇÕES DO CLOUDEVENT
	log.Println("\n☁️  CLOUDEVENT:")
	log.Println("---")
	log.Printf("  ID: %s", cloudEvent.ID)
	log.Printf("  Source: %s", cloudEvent.Source)
	log.Printf("  Type: %s", cloudEvent.Type)
	log.Printf("  Time: %s", cloudEvent.Time)
	log.Printf("  SpecVersion: %s", cloudEvent.SpecVersion)

	// 6. PROCESSAR BASEADO NO TIPO DE EVENTO
	log.Println("\n🔍 PROCESSANDO EVENTO:", eventType)
	log.Println("---")

	switch eventType {
	case "message.phone.received":
		// Mensagem recebida no Android
		var msgData models.MessageData
		if err := json.Unmarshal(cloudEvent.Data, &msgData); err != nil {
			log.Printf("❌ Erro ao parsear dados da mensagem: %v\n", err)
		} else {
			log.Println("📱 MENSAGEM RECEBIDA:")
			log.Printf("  De: %s", msgData.Contact)
			log.Printf("  Conteúdo: %s", msgData.Content)
			log.Printf("  ID: %s", msgData.ID)
			log.Printf("  Status: %s", msgData.Status)
			log.Printf("  Timestamp: %s", msgData.Timestamp)
			log.Printf("  SIM: %s", msgData.SIM)
			log.Printf("  Encriptada: %v", msgData.Encrypted)
		}

	case "message.phone.sent":
		// Mensagem enviada pelo Android
		var msgData models.MessageData
		if err := json.Unmarshal(cloudEvent.Data, &msgData); err != nil {
			log.Printf("❌ Erro ao parsear dados da mensagem: %v\n", err)
		} else {
			log.Println("📤 MENSAGEM ENVIADA:")
			log.Printf("  Para: %s", msgData.Contact)
			log.Printf("  Conteúdo: %s", msgData.Content)
			log.Printf("  ID: %s", msgData.ID)
			log.Printf("  Status: %s", msgData.Status)
		}

	case "message.phone.delivered":
		// Mensagem foi entregue ao destinatário
		var msgData models.MessageData
		if err := json.Unmarshal(cloudEvent.Data, &msgData); err != nil {
			log.Printf("❌ Erro ao parsear dados da mensagem: %v\n", err)
		} else {
			log.Println("✅ MENSAGEM ENTREGUE:")
			log.Printf("  Para: %s", msgData.Contact)
			log.Printf("  ID: %s", msgData.ID)
		}

	case "message.send.failed":
		// Falha ao enviar mensagem
		var msgData models.MessageData
		if err := json.Unmarshal(cloudEvent.Data, &msgData); err != nil {
			log.Printf("❌ Erro ao parsear dados da mensagem: %v\n", err)
		} else {
			log.Println("❌ FALHA AO ENVIAR:")
			log.Printf("  Para: %s", msgData.Contact)
			log.Printf("  Conteúdo: %s", msgData.Content)
			log.Printf("  ID: %s", msgData.ID)
		}

	case "message.send.expired":
		// Mensagem expirou antes de ser enviada
		var msgData models.MessageData
		if err := json.Unmarshal(cloudEvent.Data, &msgData); err != nil {
			log.Printf("❌ Erro ao parsear dados da mensagem: %v\n", err)
		} else {
			log.Println("⏱️  MENSAGEM EXPIRADA:")
			log.Printf("  Para: %s", msgData.Contact)
			log.Printf("  ID: %s", msgData.ID)
		}

	case "message.call.missed":
		// Chamada perdida
		var callData models.CallData
		if err := json.Unmarshal(cloudEvent.Data, &callData); err != nil {
			log.Printf("❌ Erro ao parsear dados da chamada: %v\n", err)
		} else {
			log.Println("📞 CHAMADA PERDIDA:")
			log.Printf("  De: %s", callData.Contact)
			log.Printf("  Timestamp: %s", callData.Timestamp)
			log.Printf("  Duração: %d segundos", callData.Duration)
		}

	case "phone.heartbeat.offline":
		// Telefone ficou offline (sem heartbeat por 1 hora)
		var hbData models.HeartbeatData
		if err := json.Unmarshal(cloudEvent.Data, &hbData); err != nil {
			log.Printf("❌ Erro ao parsear dados do heartbeat: %v\n", err)
		} else {
			log.Println("🔴 TELEFONE OFFLINE:")
			log.Printf("  ID: %s", hbData.ID)
			log.Printf("  Status: %s", hbData.Status)
			log.Printf("  Timestamp: %s", hbData.Timestamp)
		}

	case "phone.heartbeat.online":
		// Telefone voltou online
		var hbData models.HeartbeatData
		if err := json.Unmarshal(cloudEvent.Data, &hbData); err != nil {
			log.Printf("❌ Erro ao parsear dados do heartbeat: %v\n", err)
		} else {
			log.Println("🟢 TELEFONE ONLINE:")
			log.Printf("  ID: %s", hbData.ID)
			log.Printf("  Status: %s", hbData.Status)
			log.Printf("  Timestamp: %s", hbData.Timestamp)
		}

	default:
		log.Printf("⚠️  Tipo de evento desconhecido: %s", eventType)
	}

	log.Println(strings.Repeat("=", 70))

	// 7. RETORNAR 200 OK
	// IMPORTANTE: O HTTPSMS espera uma resposta 200 em até 5 segundos
	// Se não responder, ele vai tentar reenviar (até 4 vezes)
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
