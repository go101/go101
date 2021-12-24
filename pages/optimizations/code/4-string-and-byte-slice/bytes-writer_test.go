package bytes

import "testing"
import "io"
import "bufio"

type BytesWriter struct {
	io.Writer
	buf []byte
}

func NewBytesWriter(w io.Writer, bufLen int) *BytesWriter {
	return &BytesWriter{
		Writer: w,
		buf:    make([]byte, bufLen),
	}
}

func (sw *BytesWriter) WriteString(s string) (int, error) {
	var sum int
	var err error
	for len(s) > 0 {
		n := copy(sw.buf, s)
		n, err = sw.Write(sw.buf[:n])
		sum += n
		if err != nil || n == 0 {
			break
		}
		s = s[n:]
	}
	if err == nil && len(s) > 0 {
		err = io.ErrShortWrite
	}
	return sum, err
}

type DummyWriter struct{}

func (dw DummyWriter) Write(bs []byte) (int, error) {
	return len(bs), nil
}

var s = string(make([]byte, 500))
var w io.Writer = DummyWriter{}
var bytesw = NewBytesWriter(DummyWriter{}, 512)
var bufw = bufio.NewWriterSize(DummyWriter{}, 512)

func Benchmark_BytesWriter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesw.WriteString(s)
	}
}

func Benchmark_GeneralWriter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w.Write([]byte(s))
	}
}

func Benchmark_BufWriter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bufw.WriteString(s)
	}
}
