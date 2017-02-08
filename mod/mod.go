package mod

import (
	"bytes"
	"io"
)

type Reader struct {
	mod int
	r   io.Reader
	buf *bytes.Buffer
}

func NewReader(r io.Reader, m int) *Reader {
	return &Reader{
		m,
		r,
		new(bytes.Buffer),
	}
}
func (m Reader) Read(p []byte) (int, error) {
	maxread := int64(len(p))
	lr := io.LimitReader(m.r, maxread)
	n, _ := m.buf.ReadFrom(lr)
	if n == 0 {
		return 0, io.EOF
	}
	align := m.buf.Len()
	if int(n) > m.mod {
		align -= align % m.mod
	}
	return io.ReadFull(m.buf, p[:align])

}
