// services/sms.go
package services

import (
	"httpsms-webhook/internal/logger"
	"httpsms-webhook/internal/models"
	"regexp"
)

var vivoCodeRegex = regexp.MustCompile(`^Recarga Digital Vivo: seu codigo e (\d{5})\. Este codigo e para uso pessoal\. Por seguranca, nao deve ser repassado a vendedores ou outras pessoas\.$`)

func extractVivoCode(message string) (string, bool) {
	matches := vivoCodeRegex.FindStringSubmatch(message)

	if len(matches) == 2 {
		return matches[1], true
	}

	return "", false
}

func ProcessVivoSMS(msgData models.MessageData) {
	code, ok := extractVivoCode(msgData.Content)
	if !ok {
		logger.Log.Warn("Mensagem não é do formato esperado",
			"contact", msgData.Contact,
			"content", msgData.Content,
		)
		return
	}

	// Salva no Redis
	if err := SaveVivoCode(code); err != nil {
		logger.Log.Error("Erro ao salvar código no Redis",
			"error", err,
			"contact", msgData.Contact,
		)
		return
	}

	logger.Log.Info("Código salvo no Redis",
		"contact", msgData.Contact,
		"codigo", code,
	)
}
