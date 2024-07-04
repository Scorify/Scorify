package structs

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type FileType string

const (
	FileTypeInject     FileType = "inject"
	FileTypeSubmission FileType = "submission"
)

type File struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (file *File) Validate() error {
	if strings.Contains(file.Name, "/") || strings.Contains(file.Name, "\\") || strings.Contains(file.Name, "..") {
		return fmt.Errorf("invalid file name: %q", file.Name)
	}
	return nil
}

func (file *File) FilePath(fileType FileType, parentID uuid.UUID) (string, error) {
	err := file.Validate()
	if err != nil {
		return "", err
	}

	return filepath.Join("./files/", string(fileType), parentID.String(), file.ID.String(), file.Name), nil
}

func (file *File) APIPath(fileType FileType, parentID uuid.UUID) (string, error) {
	err := file.Validate()
	if err != nil {
		return "", err
	}

	return filepath.Join("/api/files/", string(fileType), parentID.String(), file.ID.String(), file.Name), nil
}

func RemoveInject(parentID uuid.UUID) error {
	return os.RemoveAll(filepath.Join("./files/", string(FileTypeInject), parentID.String()))
}

func (file *File) WriteFile(fileType FileType, parentID uuid.UUID, reader io.ReadSeeker) error {
	filePath, err := file.FilePath(fileType, parentID)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}

	fileHandle, err := os.Create(filePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(fileHandle, reader)
	return err
}

func (file *File) DeleteFile(fileType FileType, parentID uuid.UUID) error {
	filePath, err := file.FilePath(fileType, parentID)
	if err != nil {
		return err
	}

	return os.RemoveAll(filepath.Dir(filePath))
}
