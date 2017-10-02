package quickfix

import (
	"bytes"
	"errors"
	"io"
	"time"

	"github.com/quickfixgo/quickfix/internal"
)

const (
	defaultBufSize = 4096
)

var bufferPool internal.BufferPool

type parser struct {
	//buffer is a slice of bigBuffer
	bigBuffer, buffer []byte
	reader            io.Reader
	lastRead          time.Time
}

func newParser(reader io.Reader) *parser {
	return &parser{reader: reader}
}

func (p *parser) readMore() (int, error) {
	if len(p.buffer) == cap(p.buffer) {
		var newBuffer []byte
		switch {
		//initialize the parser
		case len(p.bigBuffer) == 0:
			p.bigBuffer = make([]byte, defaultBufSize)
			newBuffer = p.bigBuffer[0:0]

		//shift buffer back to the start of bigBuffer
		case 2*len(p.buffer) <= len(p.bigBuffer):
			newBuffer = p.bigBuffer[0:len(p.buffer)]

		//reallocate big buffer with enough space to shift buffer
		default:
			p.bigBuffer = make([]byte, 2*len(p.buffer))
			newBuffer = p.bigBuffer[0:len(p.buffer)]
		}

		copy(newBuffer, p.buffer)
		p.buffer = newBuffer
	}

	n, e := p.reader.Read(p.buffer[len(p.buffer):cap(p.buffer)])
	p.lastRead = time.Now()
	p.buffer = p.buffer[:len(p.buffer)+n]
	return n, e
}

func (p *parser) findIndex(delim []byte) (int, error) {
	return p.findIndexAfterOffset(0, delim)
}

func (p *parser) findIndexAfterOffset(offset int, delim []byte) (int, error) {
	for {
		if offset > len(p.buffer) {
			if n, err := p.readMore(); n == 0 && err != nil {
				return -1, err
			}

			continue
		}

		if index := bytes.Index(p.buffer[offset:], delim); index != -1 {
			return index + offset, nil
		}

		n, err := p.readMore()

		if n == 0 && err != nil {
			return -1, err
		}
	}
}

func (p *parser) findStart() (int, error) {
	return p.findIndex([]byte("8="))
}

func (p *parser) findEndAfterOffset(offset int) (int, error) {
	index, err := p.findIndexAfterOffset(offset, []byte("\00110="))
	if err != nil {
		return index, err
	}

	index, err = p.findIndexAfterOffset(index+1, []byte("\001"))
	if err != nil {
		return index, err
	}

	return index + 1, nil
}

func (p *parser) jumpLength() (int, error) {
	lengthIndex, err := p.findIndex([]byte("9="))
	if err != nil {
		return 0, err
	}

	lengthIndex += 3

	offset, err := p.findIndexAfterOffset(lengthIndex, []byte("\001"))
	if err != nil {
		return 0, err
	}

	if offset == lengthIndex {
		return 0, errors.New("No length given")
	}

	length, err := atoi(p.buffer[lengthIndex:offset])
	if err != nil {
		return length, err
	}

	if length <= 0 {
		return length, errors.New("Invalid length")
	}

	return offset + length, nil
}

func (p *parser) ReadMessage() (msgBytes *bytes.Buffer, err error) {
	start, err := p.findStart()
	if err != nil {
		return
	}
	p.buffer = p.buffer[start:]

	index, err := p.jumpLength()
	if err != nil {
		return
	}

	index, err = p.findEndAfterOffset(index)
	if err != nil {
		return
	}

	msgBytes = bufferPool.Get()
	msgBytes.Reset()
	msgBytes.Write(p.buffer[:index])
	p.buffer = p.buffer[index:]

	return
}
