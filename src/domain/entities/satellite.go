package entities

// Satellite ...comment
type Satellite struct {
	Name     string
	Position Position
}

// SatelliteMeasure comment
type SatelliteMeasure struct {
	Name     string
	Distance float32
	Message  []string
}
