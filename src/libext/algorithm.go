package libext

import (
	"math"
	"math/cmplx"
	"strings"
)

// Location comment
type Location struct {
}

// Calculate comment
func (location Location) Calculate(x1 float32, y1 float32,
	x2 float32, y2 float32,
	x3 float32, y3 float32,
	d1 float32, d2 float32, d3 float32) (float32, float32, bool) {

	pos1 := complex(x1, y1)
	pos2 := complex(x2, y2)
	pos3 := complex(x3, y3)

	locationRecovery := true
	point := complex(0.0, 0.0)
	for theta := 0.0; theta < 2*math.Pi; theta += 0.000001 {
		referencePoint := complex128(pos1) + cmplx.Rect(float64(d1), theta)
		if math.Abs(float64(cmplx.Abs(referencePoint-complex128(pos2)))-float64(d2)) < 0.1 &&
			math.Abs(float64(cmplx.Abs(referencePoint-complex128(pos3)))-float64(d3)) < 0.1 {
			point = referencePoint
			locationRecovery = false
			break
		}
	}

	return float32(real(point)), float32(imag(point)), locationRecovery
}

// Message comment
type Message struct {
}

// Recovery comment
func (message Message) Recovery(message1 []string, message2 []string, message3 []string) (string, bool) {
	return recoverMessage(message1, message2, message3)
}

func recoverMessage(messages ...[]string) (string, bool) {

	max := max([]int{len(messages[0]), len(messages[1]), len(messages[2])})

	addEmptyMessages(&messages[0], max-len(messages[0]))
	addEmptyMessages(&messages[1], max-len(messages[1]))
	addEmptyMessages(&messages[2], max-len(messages[2]))

	messageRecovered := []string{}
	lastWord := ""
	for col := 0; col < max; col++ {
		words := getWords(messages[0][col], messages[1][col], messages[2][col])
		switch len(words) {
		case 0:
		case 1:
			if lastWord != words[0] {
				lastWord = words[0]
				messageRecovered = append(messageRecovered, lastWord)
			}
		case 2:
			if lastWord != words[0] && lastWord != words[1] {
				return "", true
			}
			if lastWord != words[0] {
				lastWord = words[0]
				messageRecovered = append(messageRecovered, lastWord)
			} else if lastWord != words[1] {
				lastWord = words[1]
				messageRecovered = append(messageRecovered, lastWord)
			}
		default:
			return "", true
		}
	}

	return strings.Join(messageRecovered, " "), false
}

func getWords(words ...string) []string {
	addwords := []string{}
	for n := 0; n < len(words); n++ {
		if words[n] != "" && !contains(addwords, words[n]) {
			addwords = append(addwords, words[n])
		}
	}
	return addwords
}

func contains(words []string, elem string) bool {
	for _, word := range words {
		if elem == word {
			return true
		}
	}
	return false
}

func max(messageSizes []int) int {
	max := messageSizes[0]
	for _, value := range messageSizes {
		if value > max {
			max = value
		}
	}
	return max
}

func addEmptyMessages(message *[]string, n int) {
	i := 0
	for i < n {
		*message = append(*message, "")
		i++
	}
}
