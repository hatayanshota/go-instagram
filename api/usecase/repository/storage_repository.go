package repository

import "io"

type StrageRepository interface {
	UploadFile(imagefile io.Reader, id string, contenType string) error
}
