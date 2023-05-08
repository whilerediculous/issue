import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)


// when run online, the address of managerImpl.hitcount is changed while call atomic add every time, and load it in other routine value is zero
var managerOnce *sync.Once = new(sync.Once)
var managerInstance manager

type manager interface {
	Get() int64
	GetInfo() int64
}

func NewManager() manager {

	managerOnce.Do(func() {
		managerInstance = &managerImpl{}
	})
	return managerInstance
}

type managerImpl struct {
	hitCount int64
}

func (m *managerImpl) Get() int64 {
	req := atomic.AddInt64(&m.hitCount, 1)
	fmt.Println(&m.hitCount)
	return req
}

func (m *managerImpl) GetInfo() int64 {
	return atomic.LoadInt64(&m.hitCount)
}

func main() {

	for i := 0; i < 1000000; i++ {

		go func() {
			a := NewManager().Get()
			fmt.Println("first ", a)
		}()

	}

	for i := 0; i < 1000000; i++ {
		info := NewManager().GetInfo()
		fmt.Println("info ", info)
	}
	time.Sleep(1000)
	go func() {
		info := NewManager().GetInfo()
		fmt.Println("info ", info)
	}()
	time.Sleep(300)

}
