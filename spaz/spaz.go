package spaz

import (
	"io"
	"math/rand"
	"time"
)

type SpazReader struct {
	buf []byte
	rs  *rand.Rand
}

// NewSpazReader creates a strange
func NewReader(buf []byte) io.Reader {
	return &SpazReader{buf, rand.New(rand.NewSource(time.Now().Unix()))}
}

// Read reads a random number of bytes into p. The number
// of bytes read is a randomly-selected n between 0 and remainder
// of bytes in the underlying buffer or the length of p.
//
func (sr *SpazReader) Read(p []byte) (n int, err error) {
	k := len(sr.buf)
	l := len(p)
	m := min(k, l)

	if k == 0 {
		return 0, io.EOF
	}

	r := int(sr.rs.Int31n(int32(m)))
	n = copy(p, sr.buf[:r])
	sr.buf = sr.buf[n:]

	return n, nil
}

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}