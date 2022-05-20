package entity

import "gorm.io/gorm"

type TemplateResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type TemplateChannelResponse struct {
	Data  interface{}
	Error error
}

type BulkConnection struct {
	Sql map[string]*gorm.DB
}
