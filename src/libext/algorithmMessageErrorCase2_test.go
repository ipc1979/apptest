package libext

import (
	"testing"
)

func TestMessageErrorCase2(t *testing.T) {
	message := Message{}

	messageRecovered, err := message.Recovery([]string{"", "A", "", "C", ""}, []string{"", "B", "", "D"}, []string{"", "C", "", ""})

	if messageRecovered != "" && !err {
		t.Errorf("Message Error - Expected '' : Recovered '%s'.", messageRecovered)
	}

}
