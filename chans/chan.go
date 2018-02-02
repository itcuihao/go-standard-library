package chans

import (
	"log"
	"math"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type c int

// 这不是检查chan是否关闭的通用方法
func isClosed(ch <-chan c) bool {
	select {
	default:
	case <-ch:
		return true
	}
	return false
}

// Solutions Which Close Channels Rudely
func safeClose(ch chan c) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()

	close(ch) // if ch not is nil to panic
	return true
}

func safeSend(ch chan c, value c) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()

	ch <- value // panic if ch is closed
	return false
}

// Solutions Which Close Channels Politely

// using sync.Once
type myChanOnce struct {
	C    chan c
	once sync.Once
}

func newMyChanOnce() *myChanOnce {
	return &myChanOnce{
		C: make(chan c),
	}
}

func (mc *myChanOnce) safeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}

// using sync.Mutex 避免多次关闭
type myChanMutex struct {
	C      chan c
	closed bool
	mutex  sync.Mutex
}

func newMyChanMutex() *myChanMutex {
	return &myChanMutex{
		C: make(chan c),
	}
}

func (mc *myChanMutex) safeClose() {
	mc.mutex.Lock()
	if !mc.closed {
		close(mc.C)
		mc.closed = true
	}
	mc.mutex.Unlock()
}

func (mc *myChanMutex) isClosed() bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	return mc.closed
}

// ------------------------------------------
// Solutions Which Close Channels Gracefully

// 1.M receivers, one sender,
// the sender says "no more sends" by closing the data channel
func senderSayStop() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const maxRandomNumber = 100
	const numReceivers = 10

	wg := sync.WaitGroup{}
	wg.Add(numReceivers)

	ch := make(chan int, 100)

	// the sender
	go func() {
		for {
			value := rand.Intn(maxRandomNumber)
			if value == 0 {
				close(ch)
				return
			}
			ch <- value
		}
	}()

	// receivers
	for i := 0; i < numReceivers; i++ {
		go func() {
			defer wg.Done()
			for value := range ch {
				log.Println(value)
			}
		}()
	}

	wg.Wait()
}

// 2.One receiver, N senders,
// the receiver says "please stop sending more"
// by closing an additional signal channel
func receiverSayStop() {
	// In this example, the channel dataCh is never closed.
	// Yes, channels don't have to be closed.
	// A channel will be eventually garbage collected if no goroutines reference it any more,
	// whether it is closed or not.

	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const maxRandomNumber = 10
	const numSenders = 10

	wg := sync.WaitGroup{}
	wg.Add(1)

	ch := make(chan int, 100)

	// stopCh is an additional signal channel.
	// Its sender is the receiver of channel dataCh.
	// Its reveivers are the senders of channel dataCh.
	stopch := make(chan struct{})

	// the senders
	for i := 0; i < math.MaxUint8; i++ {
		go func() {
			select {
			default:
			case <-stopch:
				return
			}

			select {
			case <-stopch:
				return
			case ch <- rand.Intn(maxRandomNumber):
			}
		}()
	}

	// the receiver
	go func() {
		defer wg.Done()

		for value := range ch {
			log.Println(value)

			if value == maxRandomNumber-1 {
				// The receiver of the dataCh channel is
				// also the sender of the stopCh cahnnel.
				// It is safe to close the stop channel here.
				close(stopch)
				return
			}
		}
	}()

	wg.Wait()
}

// 3. M receivers, N senders,
// random one of them says "let's end the game"
// by notifying a moderator to close an additional signal channel
func moderatorSayStop() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const maxRandomNumber = 10
	const numReceivers = 10
	const numSenders = 1000

	wg := sync.WaitGroup{}
	wg.Add(numReceivers)

	ch := make(chan int, 100)
	stopCh := make(chan struct{})
	toStop := make(chan string, 1)

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	// senders
	for i := 0; i < numSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(maxRandomNumber)
				log.Println("sender:", value)
				if value == 11 {
					select {
					default:
					case toStop <- "sender#" + id:
					}
					return
				}

				select {
				default:
				case <-stopCh:
					return
				}

				select {
				case <-stopCh:
					return
				case ch <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < numReceivers; i++ {
		go func(id string) {
			defer wg.Done()

			for {
				select {
				default:
				case <-stopCh:
					return
				}

				select {
				case <-stopCh:
					return
				case value := <-ch:
					log.Println("receiver:", value)
					if value == maxRandomNumber-1 {
						select {
						default:
						case toStop <- "receivers#" + id:
						}
						return
					}
				}
			}
		}(strconv.Itoa(i))
	}
	wg.Wait()
	log.Println("stopped by: ", stoppedBy)
}
