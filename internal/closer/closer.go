package closer

import (
	"log"
	"os"
	"os/signal"
	"sync"
)

var globalCloser = New()

// Closer – инструмент для закрытия ресурсов
type Closer struct {
	mu    sync.Mutex
	funcs []func() error
	done  chan struct{}
	once  sync.Once
}

// Add – добавить функцию для закрытия ресурса
func Add(f ...func() error) {
	globalCloser.Add(f...)
}

// Wait – ожидание закрытия ресурсов
func Wait() {
	globalCloser.Wait()
}

// CloseAll – закрытие всех ресурсов
func CloseAll() {
	globalCloser.CloseAll()
}

// New – создаем новый экземпляр Closer
func New(sig ...os.Signal) *Closer {
	c := &Closer{done: make(chan struct{})}
	if len(sig) > 0 {
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, sig...)
			<-ch
			signal.Stop(ch)
			c.CloseAll()
		}()
	}
	return c
}

// Add – добавить функцию для закрытия ресурса
func (c *Closer) Add(f ...func() error) {
	c.mu.Lock()
	c.funcs = append(c.funcs, f...)
	c.mu.Unlock()
}

// Wait – ожидание закрытия ресурса
func (c *Closer) Wait() {
	<-c.done
}

// CloseAll – закрытие всех ресурсов
func (c *Closer) CloseAll() {
	c.once.Do(func() {
		defer close(c.done)

		c.mu.Lock()
		funcs := c.funcs
		c.mu.Unlock()

		errs := make(chan error, len(funcs))
		for _, f := range funcs {
			go func(f func() error) {
				errs <- f()
			}(f)
		}

		for i := 0; i < cap(errs); i++ {
			if err := <-errs; err != nil {
				log.Println("error returned from Closer")
			}
		}
	})
}
