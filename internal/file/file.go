package file

import (
	"fmt"

	"github.com/google/uuid"
)

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(name string, data []byte) (*File, error) {
	fileID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &File{
		ID:   fileID,
		Name: name,
		Data: data,
	}, nil
}

func (f *File) String() string {
	return fmt.Sprintf("File{ID: %s, Name: %s, Data: %s}", f.ID, f.Name, f.Data)
}
