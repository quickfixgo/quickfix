package internal

import (
	"bytes"
	"sync"
)

//BufferPool is an concurrently safe pool for byte buffers.  Used to constructing inbound messages and writing outbound messages
type BufferPool struct {
	b []*bytes.Buffer
	sync.Mutex
}

//Get returns a buffer from the pool, or creates a new buffer if the pool is empty
func (p *BufferPool) Get() (buf *bytes.Buffer) {
	p.Lock()
	if len(p.b) > 0 {
		buf, p.b = p.b[len(p.b)-1], p.b[:len(p.b)-1]
	} else {
		buf = new(bytes.Buffer)
	}

	p.Unlock()

	return
}

//Put returns adds a buffer to the pool
func (p *BufferPool) Put(buf *bytes.Buffer) {
	if buf == nil {
		panic("Nil Buffer inserted into pool")
	}

	p.Lock()
	p.b = append(p.b, buf)
	p.Unlock()
}
