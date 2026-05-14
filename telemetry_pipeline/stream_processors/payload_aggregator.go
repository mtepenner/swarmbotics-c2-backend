package streamprocessors

type PayloadSample struct {
	VehicleID   string
	Latitude    float64
	Longitude   float64
	Confidence  float64
	TargetClass string
}

type AggregatedPayloadState struct {
	TrackCount       int      `json:"trackCount"`
	DominantClass    string   `json:"dominantClass"`
	AverageLatitude  float64  `json:"averageLatitude"`
	AverageLongitude float64  `json:"averageLongitude"`
	Contributors     []string `json:"contributors"`
}

func AggregatePayloads(samples []PayloadSample) AggregatedPayloadState {
	state := AggregatedPayloadState{Contributors: make([]string, 0, len(samples))}
	if len(samples) == 0 {
		return state
	}

	classCounts := map[string]int{}
	for _, sample := range samples {
		state.TrackCount++
		state.AverageLatitude += sample.Latitude
		state.AverageLongitude += sample.Longitude
		state.Contributors = append(state.Contributors, sample.VehicleID)
		classCounts[sample.TargetClass]++
	}

	for className, count := range classCounts {
		if count > classCounts[state.DominantClass] {
			state.DominantClass = className
		}
	}

	state.AverageLatitude /= float64(len(samples))
	state.AverageLongitude /= float64(len(samples))
	return state
}
