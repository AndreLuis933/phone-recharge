package models

import "encoding/json"

type CloudEvent struct {
	ID              string          `json:"id"`              // ID único do evento
	Source          string          `json:"source"`          // De onde veio (ex: httpsms.com)
	SpecVersion     string          `json:"specversion"`     // Versão do CloudEvent (geralmente "1.0")
	Type            string          `json:"type"`            // Tipo do evento (ex: message.phone.received)
	DataContentType string          `json:"datacontenttype"` // Tipo do conteúdo (application/json)
	Time            string          `json:"time"`            // Timestamp do evento
	Data            json.RawMessage `json:"data"`            // Os dados reais do evento (varia por tipo)
}

type MessageData struct {
	ID        string `json:"id"`         // ID da mensagem
	Owner     string `json:"owner"`      // Dono do telefone
	Contact   string `json:"contact"`    // Número de quem enviou/recebeu
	Content   string `json:"content"`    // Texto da mensagem
	Status    string `json:"status"`     // Status (received, sent, delivered, etc)
	Timestamp string `json:"timestamp"`  // Quando a mensagem foi processada
	SIM       string `json:"sim"`        // Qual SIM card (se dual chip)
	Encrypted bool   `json:"encrypted"`  // Se estava encriptada
	RequestID string `json:"request_id"` // ID da requisição (se você enviou)
}

// CallData é a estrutura para eventos de chamada perdida
type CallData struct {
	ID        string `json:"id"`
	Owner     string `json:"owner"`
	Contact   string `json:"contact"`
	Timestamp string `json:"timestamp"`
	Duration  int    `json:"duration"` // Duração em segundos
}

// HeartbeatData é para eventos de online/offline do telefone
type HeartbeatData struct {
	ID        string `json:"id"`
	Owner     string `json:"owner"`
	Timestamp string `json:"timestamp"`
	Status    string `json:"status"` // online ou offline
}
