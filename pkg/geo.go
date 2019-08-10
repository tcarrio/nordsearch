package nordsearch

// Location stores the latitudinal and longitudinal coordinates of a location
type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

// DistanceFrom calculates the total distance from a target location in meters
func (l Location) DistanceFrom(target Location) int64 {
	return 0
}
