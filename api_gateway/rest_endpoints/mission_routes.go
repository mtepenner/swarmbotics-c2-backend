package restendpoints

import (
	"sort"
	"time"
)

type RouteLeg struct {
	VehicleID  string    `json:"vehicleId"`
	Sequence   int       `json:"sequence"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	ETASeconds int       `json:"etaSeconds"`
	WindowOpen time.Time `json:"windowOpen"`
}

type MissionRouteResponse struct {
	MissionID string     `json:"missionId"`
	Objective string     `json:"objective"`
	IssuedAt  time.Time  `json:"issuedAt"`
	RouteLegs []RouteLeg `json:"routeLegs"`
}

func BuildMissionRouteResponse(missionID, objective string, vehicleIDs []string, now time.Time) MissionRouteResponse {
	orderedVehicles := append([]string(nil), vehicleIDs...)
	sort.Strings(orderedVehicles)

	legs := make([]RouteLeg, 0, len(orderedVehicles))
	for index, vehicleID := range orderedVehicles {
		legs = append(legs, RouteLeg{
			VehicleID:  vehicleID,
			Sequence:   index + 1,
			Latitude:   33.6500 + float64(index)*0.0012,
			Longitude:  -117.8400 - float64(index)*0.0011,
			ETASeconds: 45 * (index + 1),
			WindowOpen: now.Add(time.Duration(index) * 15 * time.Second),
		})
	}

	return MissionRouteResponse{
		MissionID: missionID,
		Objective: objective,
		IssuedAt:  now.UTC(),
		RouteLegs: legs,
	}
}
