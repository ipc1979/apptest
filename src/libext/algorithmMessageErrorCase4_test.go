package libext

import (
	"testing"
)

func TestMessageErrorCase4(t *testing.T) {
	message := Message{}

	messageRecovered, err := message.Recovery([]string{"", "A", "", ""}, []string{"A", "", "", "D", ""}, []string{"", "B", "C", ""})

	if messageRecovered != "A,B,C,D" && err {
		t.Errorf("Message Error - Expected '' : Recovered '%s'.", messageRecovered)
	}

}
