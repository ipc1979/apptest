package algorithms

// LocationAlgorithm comment
type LocationAlgorithm interface {
	Calculate(x1 float32, y1 float32,
		x2 float32, y2 float32,
		x3 float32, y3 float32,
		d1 float32, d2 float32, d3 float32) (float32, float32, bool)
}

// MessageAlgorithm comment
type MessageAlgorithm interface {
	Recovery(message1 []string, message2 []string, message3 []string) (string, bool)
}
