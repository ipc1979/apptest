package libext

import (
	"testing"
)

func TestMessage(t *testing.T) {
	message := Message{}

	messageRecovered, err := message.Recovery([]string{"", "A", "", "C", ""}, []string{"", "B", "", "D"}, []string{"A", "", "C", ""})

	if messageRecovered != "A,B,C,D" && err {
		t.Errorf("Message Error - Expected 'A,B,C,D' : Recovered '%s'.", messageRecovered)
	}

}
