package api

const (
	EventMessageReceived  = "message.phone.received"
	EventMessageSent      = "message.phone.sent"
	EventMessageDelivered = "message.phone.delivered"
	EventMessageFailed    = "message.send.failed"
	EventMessageExpired   = "message.send.expired"
	EventCallMissed       = "message.call.missed"
	EventHeartbeatOffline = "phone.heartbeat.offline"
	EventHeartbeatOnline  = "phone.heartbeat.online"
)