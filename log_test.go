package log_test

import (
	"io"
	"io/ioutil"
	"sync"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"go.ideatocode.tech/log"
)

func TestColorLogDefault(t *testing.T) {
	r, w := io.Pipe()

	log.TimestampFunc = func() time.Time {
		tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
		return tm
	}
	x := log.NewColor(w)

	wg := sync.WaitGroup{}
	wg.Add(1)

	var buf []byte
	go func() {
		buf, _ = ioutil.ReadAll(r)
		// fmt.Fprintf(os.Stderr, string(buf))
		wg.Done()
	}()

	x.Info("testing")
	x.Error("testing")
	w.Close()

	wg.Wait()

	assert.Equal(t, "2006-01-02 15:04:05+07:00 \x1b[94m[Info]\x1b[0m testing\n"+
		"2006-01-02 15:04:05+07:00 \x1b[31m[Error]\x1b[0m testing\n", string(buf))
}
func TestNoLogDefault(t *testing.T) {
	r, w := io.Pipe()

	log.TimestampFunc = func() time.Time {
		tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
		return tm
	}

	log.DefaultLogger = log.NewColor(w)

	wg := sync.WaitGroup{}
	wg.Add(1)

	var buf []byte
	go func() {
		buf, _ = ioutil.ReadAll(r)
		// fmt.Fprintf(os.Stderr, string(buf))
		wg.Done()
	}()

	log.Info("testing")
	log.Error("testing")
	w.Close()

	wg.Wait()

	assert.Equal(t, "2006-01-02 15:04:05+07:00 \x1b[94m[Info]\x1b[0m testing\n"+
		"2006-01-02 15:04:05+07:00 \x1b[31m[Error]\x1b[0m testing\n", string(buf))
}

func TestStdLogDefault(t *testing.T) {
	r, w := io.Pipe()

	log.TimestampFunc = func() time.Time {
		tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
		return tm
	}
	x := log.NewStd(w)

	wg := sync.WaitGroup{}
	wg.Add(1)

	var buf []byte
	go func() {
		buf, _ = ioutil.ReadAll(r)
		// fmt.Fprintf(os.Stderr, string(buf))
		wg.Done()
	}()

	x.Info("testing")
	x.Debug("testing")
	w.Close()

	wg.Wait()

	assert.Equal(t, "2006-01-02 15:04:05+07:00 [Info] testing\n"+
		"2006-01-02 15:04:05+07:00 [Debug] testing\n", string(buf))
}
func TestZeroLogDefault(t *testing.T) {
	r, w := io.Pipe()
	x := log.NewZero(w)

	zerolog.TimestampFunc = func() time.Time {
		tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
		return tm
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	var buf []byte
	go func() {
		buf, _ = ioutil.ReadAll(r)
		wg.Done()
	}()

	x.Info("testing")
	x.Debug("testing")
	w.Close()

	wg.Wait()

	assert.Equal(t, `{"level":"info","time":"2006-01-02 15:04:05+07:00","message":"testing"}`+"\n"+
		`{"level":"debug","time":"2006-01-02 15:04:05+07:00","message":"testing"}`+"\n", string(buf))
}
func TestZerologPretty(t *testing.T) {
	r, w := io.Pipe()
	x := log.NewZero(w).Pretty()

	zerolog.TimestampFunc = func() time.Time {
		tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
		return tm
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	var buf []byte
	go func() {
		buf, _ = ioutil.ReadAll(r)
		wg.Done()
		// fmt.Fprintf(os.Stderr, string(buf))
	}()

	x.Info("testing")
	x.Debug("testing")
	w.Close()

	wg.Wait()

	assert.Equal(t, "\x1b[90m2006-01-02 15:04:05+07:00\x1b[0m \x1b[32mINF\x1b[0m testing\n"+
		"\x1b[90m2006-01-02 15:04:05+07:00\x1b[0m \x1b[33mDBG\x1b[0m testing\n", string(buf))
}
