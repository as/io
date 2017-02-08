package min

type Reader struct {
	Min int
	r io.Reader
}

type Writer struct {
	Min int
	w io.Writer
}

func NewReader(r io.Reader, m int) *Reader{
	return &Reader{m,r}
}

func NewWriter(w io.Writer, m int) *Writer{
	return Writer{m, w}
}

func (m Reader) Read(p []byte) (n int, err error) {
	return io.ReadAtLeast(m.r, p, m.Min)
}

func (m Writer) Write(p []byte) (n int, err error) {
	return io.Copy(m.w, NewReader(bytes.NewReader(p), m.Min))
}