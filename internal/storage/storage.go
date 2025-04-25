package storage

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vipos89/storage/internal/file"
)

type Storage struct{
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File, 0),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}
	s.files[newFile.ID] = newFile
	return newFile, nil
}

func (s *Storage) GetById(fileID uuid.UUID) (*file.File, error) {
	f, ok := s.files[fileID]
	if !ok {
		return nil, fmt.Errorf("file with id %s not found", fileID)
	}
	return f, nil
}