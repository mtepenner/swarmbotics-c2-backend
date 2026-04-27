package streamprocessors

import "time"

type VehicleHealth struct {
	VehicleID       string
	BatteryPercent  float64
	TrackIntegrity  float64
	LastHeartbeatAt time.Time
}

type HealthAlert struct {
	VehicleID string   `json:"vehicleId"`
	Severity  string   `json:"severity"`
	Reasons   []string `json:"reasons"`
}

func AssessHealth(sample VehicleHealth, now time.Time) HealthAlert {
	reasons := make([]string, 0, 3)
	severity := "nominal"

	if sample.BatteryPercent < 35.0 {
		reasons = append(reasons, "battery-below-threshold")
		severity = "warning"
	}
	if sample.TrackIntegrity < 0.6 {
		reasons = append(reasons, "mobility-damage")
		severity = "critical"
	}
	if now.Sub(sample.LastHeartbeatAt) > 15*time.Second {
		reasons = append(reasons, "stale-heartbeat")
		if severity == "nominal" {
			severity = "warning"
		}
	}

	return HealthAlert{VehicleID: sample.VehicleID, Severity: severity, Reasons: reasons}
}