package services

import (
	"context"
	"fmt"

	"github.com/gorilla/websocket"
	"saml.dev/gome-assistant/internal"
)

func BuildService[
	T AlarmControlPanel |
		Cover |
		Light |
		HomeAssistant |
		Lock |
		MediaPlayer |
		Switch |
		InputBoolean |
		InputButton |
		InputDatetime |
		InputText |
		InputNumber |
		Notify |
		Number |
		Scene |
		TTS |
		Vacuum,
](conn *websocket.Conn, ctx context.Context) *T {
	return &T{conn: conn, ctx: ctx}
}

type BaseServiceRequest struct {
	Id          string         `json:"id"`
	RequestType string         `json:"type"` // hardcoded "call_service"
	Domain      string         `json:"domain"`
	Service     string         `json:"service"`
	ServiceData map[string]any `json:"service_data,omitempty"`
	Target      struct {
		EntityId string `json:"entity_id,omitempty"`
	} `json:"target,omitempty"`
}

func NewBaseServiceRequest(entityId string) BaseServiceRequest {
	id := internal.GetId()
	bsr := BaseServiceRequest{
		Id:          fmt.Sprint(id),
		RequestType: "call_service",
	}
	if entityId != "" {
		bsr.Target.EntityId = entityId
	}
	return bsr
}
