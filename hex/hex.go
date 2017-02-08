package hex

import (
	"bufio"
	gohex "encoding/hex"
	"io"
)

type Reader struct {
	mr io.Reader
}
type Writer struct {
	uw io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		uw: bufio.NewWriter(w),
	}
}
func NewReader(r io.Reader) *Reader {
	return &Reader{
		mr: bufio.NewReader(r),
	}
}
func (r Reader) Read(p []byte) (n int, err error) {
	if n, err = r.mr.Read(p); err != nil {
		return
	}
	println("modreader reads", n)
	return gohex.Decode(p[:n], p[:n])
}
func (w Writer) Write(p []byte) (n int, err error) {
	q := make([]byte, gohex.EncodedLen(len(p)))
	n = gohex.Encode(q, p)
	return w.uw.Write(q)
}
