package entities

// MeasureRequest comment
type MeasureRequest struct {
	Distance float32  `json:"distance" binding:"required"`
	Message  []string `json:"message" binding:"required"`
}

// SatelliteRequest comment
type SatelliteRequest struct {
	Name     string   `json:"name" binding:"required"`
	Distance float32  `json:"distance" binding:"required"`
	Message  []string `json:"message" binding:"required"`
}

// SatellitesRequest comment
type SatellitesRequest struct {
	Satellites []SatelliteRequest `json:"satellites" binding:"required"`
}
