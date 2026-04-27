package ingestiongateways

import (
	"encoding/json"
	"time"
)

type TelemetryRecord struct {
	VehicleID string         `json:"vehicleId"`
	MissionID string         `json:"missionId"`
	Payload   map[string]any `json:"payload"`
}

type KafkaEnvelope struct {
	Topic     string          `json:"topic"`
	Key       string          `json:"key"`
	Timestamp time.Time       `json:"timestamp"`
	Payload   TelemetryRecord `json:"payload"`
}

func BuildKafkaEnvelope(topic, key string, record TelemetryRecord, now time.Time) ([]byte, error) {
	envelope := KafkaEnvelope{
		Topic:     topic,
		Key:       key,
		Timestamp: now.UTC(),
		Payload:   record,
	}
	return json.Marshal(envelope)
}