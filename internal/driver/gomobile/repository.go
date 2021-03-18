// +build ios android

package gomobile

import (
	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/storage/repository"
)

// declare conformance with repository types
var _ repository.Repository = (*mobileFileRepo)(nil)
var _ repository.ListableRepository = (*mobileFileRepo)(nil)
var _ repository.WritableRepository = (*mobileFileRepo)(nil)

type mobileFileRepo struct {
	driver *mobileDriver
}

func (m *mobileFileRepo) Exists(u fyne.URI) (bool, error) {
	return true, nil // TODO check a file exists
}

func (m *mobileFileRepo) Reader(u fyne.URI) (fyne.URIReadCloser, error) {
	return fileReaderForURI(u)
}

func (m *mobileFileRepo) CanRead(u fyne.URI) (bool, error) {
	return true, nil // TODO check a file can be read
}

func (m *mobileFileRepo) Destroy(string) {
}

func (m *mobileFileRepo) CanWrite(u fyne.URI) (bool, error) {
	return true, nil // TODO check a file can be written
}

func (m *mobileFileRepo) Writer(u fyne.URI) (fyne.URIWriteCloser, error) {
	return fileWriterForURI(u)
}

func (m *mobileFileRepo) Delete(u fyne.URI) error {
	// TODO: implement this
	return repository.ErrOperationNotSupported
}

func (m *mobileFileRepo) CanList(u fyne.URI) (bool, error) {
	return canListURI(u), nil
}

func (m *mobileFileRepo) List(u fyne.URI) ([]fyne.URI, error) {
	return listURI(u)
}

func (m *mobileFileRepo) CreateListable(u fyne.URI) error {
	// TODO: implement this
	return repository.ErrOperationNotSupported
}