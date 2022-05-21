package models

import "mime/multipart"

type UploadRequest struct {
	ID        string         `json:"id"`
	TableName string         `json:"table_name"`
	File      multipart.File `json:"file"`
}
