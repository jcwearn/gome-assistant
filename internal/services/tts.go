package services

import (
	"context"

	"github.com/gorilla/websocket"
	ws "saml.dev/gome-assistant/internal/websocket"
)

/* Structs */

type TTS struct {
	conn *websocket.Conn
	ctx  context.Context
}

/* Public API */

// Remove all text-to-speech cache files and RAM cache.
func (tts TTS) ClearCache() {
	req := NewBaseServiceRequest("")
	req.Domain = "tts"
	req.Service = "clear_cache"

	ws.WriteMessage(req, tts.conn, tts.ctx)
}

// Say something using text-to-speech on a media player with cloud.
// Takes an entityId and an optional
// map that is translated into service_data.
func (tts TTS) CloudSay(entityId string, serviceData ...map[string]any) {
	req := NewBaseServiceRequest(entityId)
	req.Domain = "tts"
	req.Service = "cloud_say"
	if len(serviceData) != 0 {
		req.ServiceData = serviceData[0]
	}

	ws.WriteMessage(req, tts.conn, tts.ctx)
}

// Say something using text-to-speech on a media player with google_translate.
// Takes an entityId and an optional
// map that is translated into service_data.
func (tts TTS) GoogleTranslateSay(entityId string, serviceData ...map[string]any) {
	req := NewBaseServiceRequest(entityId)
	req.Domain = "tts"
	req.Service = "google_translate_say"
	if len(serviceData) != 0 {
		req.ServiceData = serviceData[0]
	}

	ws.WriteMessage(req, tts.conn, tts.ctx)
}
