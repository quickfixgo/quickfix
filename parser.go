package quickfixgo

import (
	"bytes"
	"io"
	"strconv"
)

const (
	defaultBufSize = 4096
)

type parser struct {
	buffer                       []byte
	reader                       io.Reader
	bytesRead, index, startIndex int
	err                          error
}

func newParser(reader io.Reader) *parser {
	return &parser{
		buffer: make([]byte, defaultBufSize),
		reader: reader,
	}
}

func newParserSize(reader io.Reader, size int) *parser {
	return &parser{
		buffer: make([]byte, size),
		reader: reader,
	}
}

func (p *parser) readMore() {
	if p.bytesRead == cap(p.buffer) {
		newBuffer := make([]byte, defaultBufSize+cap(p.buffer)-p.bytesRead)
		p.bytesRead = copy(newBuffer, p.buffer[p.startIndex:])
		p.buffer = newBuffer
		p.index = (p.index - p.startIndex)
		p.startIndex = 0
	}

	n, e := p.reader.Read(p.buffer[p.bytesRead:])
	p.bytesRead += n
	if e != nil {
		p.err = e
	}
}

func (p *parser) findIndexOffset(delim []byte) (start int, err error) {
	for {
		if p.index > p.bytesRead {
			p.readMore()
		}

		start = bytes.Index(p.buffer[p.index:p.bytesRead], delim)

		if start != -1 {
			return
		}

		p.readMore()

		if p.err != nil {
			err = p.err
			return
		}
	}

	return
}

func (p *parser) findStart() (err error) {
	var offset int
	if offset, err = p.findIndexOffset([]byte("8=")); err != nil {
		return
	}

	p.index += offset

	return
}

func (p *parser) findEnd() (err error) {
	var offset int
	if offset, err = p.findIndexOffset([]byte("\00110=")); err != nil {
		return
	}

	p.index += offset + 3

	if offset, err = p.findIndexOffset([]byte("\001")); err != nil {
		return
	}
	p.index += offset + 1

	return
}

func (p *parser) extractLength() (length int, err error) {
	var offset int
	if offset, err = p.findIndexOffset([]byte("9=")); err != nil {
		return
	}

	p.index += offset + 3

	if offset, err = p.findIndexOffset([]byte("\001")); err != nil {
		return
	}

	length, _ = strconv.Atoi(string(p.buffer[p.index : p.index+offset]))
	p.index += offset

	return
}

func (p *parser) readMessage() (msg []byte, err error) {
	p.startIndex = p.index

	if err = p.findStart(); err != nil {
		return
	}

	msgStart := p.index

	var length int
	if length, err = p.extractLength(); err != nil {
		return
	}

	p.index += length

	if err = p.findEnd(); err != nil {
		return
	}

	msg = p.buffer[msgStart:p.index]

	return
}
