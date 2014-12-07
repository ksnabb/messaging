package messaging

import "testing"

func TestMessaging(t *testing.T) {

	// create three channels with a buffer of 2
	ch1 := make(chan []byte, 2)
	ch2 := make(chan []byte, 2)
	ch3 := make(chan []byte, 2)

	// subscribe all to one group "123" and two to "12"
	Subscribe(ch1, "123")
	Subscribe(ch2, "123")
	Subscribe(ch3, "123")
	Subscribe(ch1, "12")
	Subscribe(ch2, "12")

	Send([]byte("hello group"), "123")

	// channel one, two and three should receive this message
	m, ok := <-ch1
	if !ok || string(m) != "hello group" {
		t.Fail()
	}
	m, ok = <-ch2
	if !ok || string(m) != "hello group" {
		t.Fail()
	}
	m, ok = <-ch3
	if !ok || string(m) != "hello group" {
		t.Fail()
	}

	// send message to 12
	Send([]byte("hello 12"), "12")

	// and another message to channel 3 to not block all execution
	ch3 <- []byte("no 12")

	// channel one and two should receive "hello 12"
	m, ok = <-ch1
	if !ok || string(m) != "hello 12" {
		t.Fail()
	}
	m, ok = <-ch2
	if !ok || string(m) != "hello 12" {
		t.Fail()
	}

	// channel 3 should receive message "no 12"
	m, ok = <-ch3
	if !ok || string(m) != "no 12" {
		t.Fail()
	}
}
