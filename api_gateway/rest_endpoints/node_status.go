package restendpoints

import "time"

type NodeStatus struct {
	NodeID          string        `json:"nodeId"`
	Region          string        `json:"region"`
	HeartbeatLag    time.Duration `json:"heartbeatLag"`
	Healthy         bool          `json:"healthy"`
	MissionCapacity int           `json:"missionCapacity"`
}

func SummarizeNodeStatus(regions map[string]string, heartbeatLag map[string]time.Duration) []NodeStatus {
	statuses := make([]NodeStatus, 0, len(regions))
	for nodeID, region := range regions {
		lag := heartbeatLag[nodeID]
		statuses = append(statuses, NodeStatus{
			NodeID:          nodeID,
			Region:          region,
			HeartbeatLag:    lag,
			Healthy:         lag < 10*time.Second,
			MissionCapacity: max(1, 8-int(lag.Seconds()/2)),
		})
	}
	return statuses
}

func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}
