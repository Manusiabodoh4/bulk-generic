package databases

import (
	"github.com/Manusiabodoh4/bulk-generic/bin/entity"
	"gorm.io/gorm"
)

type BulkConnectionPkg struct {
	connection entity.BulkConnection
}

var obj BulkConnectionPkg

func InitBulkConnectionPkg() *BulkConnectionPkg {
	if obj.connection.Sql == nil {
		obj = BulkConnectionPkg{}
		obj.connection.Sql = make(map[string]*gorm.DB)
	}
	return &obj
}

func (h *BulkConnectionPkg) AddBulkConnection(prop string, value *gorm.DB) bool {
	if _, ok := h.connection.Sql[prop]; ok {
		return false
	}
	h.connection.Sql[prop] = value
	return true
}

func (h *BulkConnectionPkg) GetBulkConnection(prop string) (bool, *gorm.DB) {
	val, ok := h.connection.Sql[prop]
	return ok, val
}
