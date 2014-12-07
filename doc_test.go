package messaging

func ExampleSubscribe() {
  c := make(chan []byte, 10)
  messaging.Subscribe(c, "examplegroup")
}

func ExampleUnsubscribe() {
  messaging.Unsubscribe(c, "examplegroup")
}

func ExampleSend() {
  messaging.Send([]byte("hello group"), "examplegroup")
}
