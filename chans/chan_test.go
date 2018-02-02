package chans

import (
	"testing"
	"time"
)

func TestIsClosed(t *testing.T) {
	ch := make(chan c)
	t.Log(isClosed(ch)) //false
	close(ch)
	t.Log(isClosed(ch)) //true
}

// Close Channels Gracefully in Golang

// 原则：
// don't close a channel from the receiver side and
// don't close a channel if the channel has
// multiple concurrent senders.

// Solutions Which Close Channels Rudely
func TestSafeClose(t *testing.T) {
	ch := make(chan c)
	t.Log(safeClose(ch))
}

func TestSafeSend(t *testing.T) {
	ch := make(chan c)
	t.Log(safeSend(ch, 1))
}

func TestSyncOnce(t *testing.T) {
	ch := newMyChanOnce()
	go func() {
		ch.C <- 1
		t.Log(<-ch.C)
	}()
	time.Sleep(1 * time.Second)
	ch.safeClose()
	t.Log(<-ch.C)
}

func TestSyncMutex(t *testing.T) {

	mc := &myChanMutex{C: make(chan c)}
	t.Log(mc.isClosed())
	mc.safeClose()
	t.Log(mc.isClosed())
}

func TestSenderSayStop(t *testing.T) {
	senderSayStop()
}
func TestReceiverSayStop(t *testing.T) {
	receiverSayStop()
}
