package libext

import (
	"testing"
)

func TestMessageErrorCase1(t *testing.T) {
	message := Message{}

	messageRecovered, err := message.Recovery([]string{"", "A", "", "C", ""}, []string{"", "B", "", "D"}, []string{"", "B", "C", ""})

	if messageRecovered != "" && !err {
		t.Errorf("Message Error - Expected '' : Recovered '%s'.", messageRecovered)
	}

}
