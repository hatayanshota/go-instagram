package model

type UploadImage struct {
	ACL:         *string
	ContentType: *string
	Bucket:      *string
	Key:         *string
	Body:        io.Reader
}