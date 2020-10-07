package libext

import (
	"testing"
)

func TestMessageErrorCase3(t *testing.T) {
	message := Message{}

	messageRecovered, err := message.Recovery([]string{"A", "B", "", ""}, []string{"", "B", "", "D", ""}, []string{"", "A", "B", ""})

	if messageRecovered != "A,B,C,D" && err {
		t.Errorf("Message Error - Expected '' : Recovered '%s'.", messageRecovered)
	}

}
