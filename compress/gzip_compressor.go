package compress

import (
	"bytes"
	"compress/gzip"
	"io"
	"sync"
)

type GzipCompressor struct {
	readerPool sync.Pool
	writerPool sync.Pool
}

func (c *GzipCompressor) Compress(data []byte) ([]byte, error) {
	if data == nil || len(data) == 0 {
		return data, nil
	}

	buffer := bytes.NewBuffer(nil)
	writer, ok := c.writerPool.Get().(*gzip.Writer)
	if !ok {
		writer = gzip.NewWriter(buffer)
	} else {
		writer.Reset(buffer)
	}
	defer c.writerPool.Put(writer)

	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (c *GzipCompressor) Decompress(data []byte) ([]byte, error) {
	if data == nil || len(data) == 0 {
		return data, nil
	}

	reader := bytes.NewReader(data)
	gr, ok := c.readerPool.Get().(*gzip.Reader)
	if !ok {
		newReader, err := gzip.NewReader(reader)
		if err != nil {
			return nil, err
		}
		gr = newReader
	} else {
		if err := gr.Reset(reader); err != nil {
			return nil, err
		}
	}

	defer func() {
		if gr != nil {
			c.readerPool.Put(gr)
		}
	}()

	out, err := io.ReadAll(gr)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return nil, err
	}
	return out, nil
}
