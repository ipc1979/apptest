package entities

import (
	"appml/src/interfaces/algorithms"
)

// Alliance comment
type Alliance struct {
	Location   algorithms.LocationAlgorithm
	Message    algorithms.MessageAlgorithm
	Satellites []Satellite
}

func (alliance *Alliance) AddSatellite(satellite Satellite) {
	alliance.Satellites = append(alliance.Satellites, satellite)
}

// SetLocationAlgorithm comment
func (alliance *Alliance) SetLocationAlgorithm(location algorithms.LocationAlgorithm) {
	alliance.Location = location
}

// SetMessageAlgorithm comment
func (alliance *Alliance) SetMessageAlgorithm(message algorithms.MessageAlgorithm) {
	alliance.Message = message
}

// GetLocation comment
func (alliance *Alliance) GetLocation(distances ...float32) (Position, bool) {
	x, y, err := alliance.Location.Calculate(alliance.Satellites[0].Position.X, alliance.Satellites[0].Position.Y,
		alliance.Satellites[1].Position.X, alliance.Satellites[1].Position.Y,
		alliance.Satellites[2].Position.X, alliance.Satellites[2].Position.Y,
		distances[0], distances[1], distances[2])
	return Position{X: x, Y: y}, err
}

// GetMessage Comment
func (alliance *Alliance) GetMessage(messages ...[]string) (string, bool) {
	return alliance.Message.Recovery(messages[0], messages[1], messages[2])
}
