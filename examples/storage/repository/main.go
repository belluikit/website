// A custom storage repository that keeps "mem://" files in memory.
package main

import (
	"bytes"
	"errors"
	"io"
	"sync"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/storage"
	"github.com/belluikit/bellui/storage/repository"
	"github.com/belluikit/bellui/widget"
)

//docs:start repo
// memRepo stores file contents in a map, keyed by the URI string.
type memRepo struct {
	mu    sync.Mutex
	files map[string][]byte
}

// The four methods every repository needs.

func (m *memRepo) Exists(u bellui.URI) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.files[u.String()]
	return ok, nil
}

func (m *memRepo) Reader(u bellui.URI) (bellui.URIReadCloser, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	data, ok := m.files[u.String()]
	if !ok {
		return nil, errors.New("not found: " + u.String())
	}
	return &memReader{Reader: bytes.NewReader(data), uri: u}, nil
}

func (m *memRepo) CanRead(u bellui.URI) (bool, error) { return m.Exists(u) }

func (m *memRepo) Destroy(string) {}

// Writer, CanWrite and Delete make it a WritableRepository.

func (m *memRepo) Writer(u bellui.URI) (bellui.URIWriteCloser, error) {
	return &memWriter{repo: m, uri: u}, nil
}

func (m *memRepo) CanWrite(bellui.URI) (bool, error) { return true, nil }

func (m *memRepo) Delete(u bellui.URI) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.files, u.String())
	return nil
}

//docs:end repo

type memReader struct {
	*bytes.Reader
	uri bellui.URI
}

func (r *memReader) Close() error     { return nil }
func (r *memReader) URI() bellui.URI { return r.uri }

// memWriter buffers what is written and stores it on Close.
type memWriter struct {
	bytes.Buffer
	repo *memRepo
	uri  bellui.URI
}

func (w *memWriter) Close() error {
	w.repo.mu.Lock()
	defer w.repo.mu.Unlock()
	w.repo.files[w.uri.String()] = w.Bytes()
	return nil
}

func (w *memWriter) URI() bellui.URI { return w.uri }

func main() {
	a := app.New()
	win := a.NewWindow("Custom repository")

	//docs:start register
	repository.Register("mem", &memRepo{files: map[string][]byte{}})

	u, err := storage.ParseURI("mem://notes/hello.txt")
	if err != nil {
		panic(err)
	}
	//docs:end register

	//docs:start use
	wr, _ := storage.Writer(u)
	_, _ = wr.Write([]byte("kept in memory"))
	_ = wr.Close()

	r, _ := storage.Reader(u)
	data, _ := io.ReadAll(r)
	_ = r.Close()

	_, err = storage.List(u) // memRepo is not a ListableRepository
	listErr := errors.Is(err, repository.ErrOperationNotSupported)
	//docs:end use

	msg := "Read back: " + string(data)
	if listErr {
		msg += "\nList: not supported, as expected"
	}
	win.SetContent(container.NewVBox(widget.NewLabel(msg)))
	bellui.ShowAndRun(win)
}
