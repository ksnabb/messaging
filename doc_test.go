package messaging

var c chan []byte = make(chan []byte, 10)

func ExampleSubscribe() {
	Subscribe(c, "examplegroup")
}

func ExampleUnsubscribe() {
	Unsubscribe(c, "examplegroup")
}

func ExampleSend() {
	Send([]byte("hello group"), "examplegroup")
}
