package messaging

import (
	"math/rand"
	"testing"
)

func TestMessaging(t *testing.T) {

	// create three channels
	ch1 := make(chan []byte, 2)
	ch2 := make(chan []byte, 2)
	ch3 := make(chan []byte, 2)

	// subscribe all to one and a pair separatelly
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
	ch3 <- []byte("no 12")

	// channel one, two should receive this message
	m, ok = <-ch1
	if !ok || string(m) != "hello 12" {
		t.Fail()
	}
	m, ok = <-ch2
	if !ok || string(m) != "hello 12" {
		t.Fail()
	}
	m, ok = <-ch3
	if !ok || string(m) != "no 12" {
		t.Fail()
	}
}

// one group multiple clients
func benchmarkClientsInGroups(clients int, groups int, b *testing.B) {

	// create group names
	groupNames := make([]string, groups)
	for i := 0; i < groups; i++ {
		groupNames[i] = string(i)
	}

	for i := 0; i < clients; i++ {
		// create client channel
		cc := make(chan []byte)

		// drain the channel
		go func() {
			for {
				<-cc
			}
		}()

		// subscribe to a random channel
		Subscribe(cc, groupNames[rand.Intn(groups)])
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// send to random group
		Send([]byte("hello group"), groupNames[rand.Intn(groups)])
	}
}

// one group many clients
func BenchmarkMessaging1(b *testing.B) {
	benchmarkClientsInGroups(1, 2000, b)
}
func BenchmarkMessaging10(b *testing.B) {
	benchmarkClientsInGroups(10, 2000, b)
}
func BenchmarkMessaging100(b *testing.B) {
	benchmarkClientsInGroups(100, 2000, b)
}
func BenchmarkMessaging1000(b *testing.B) {
	benchmarkClientsInGroups(1000, 2000, b)
}
func BenchmarkMessaging10000(b *testing.B) {
	benchmarkClientsInGroups(10000, 2000, b)
}
func BenchmarkMessaging100000(b *testing.B) {
	benchmarkClientsInGroups(100000, 2000, b)
}
