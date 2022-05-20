package databases

import (
	"github.com/Manusiabodoh4/bulk-generic/bin/entity"
	"gorm.io/gorm"
)

type BulkConnectionPkg struct {
	Connection entity.BulkConnection
}

var obj BulkConnectionPkg

func InitBulkConnectionPkg() *BulkConnectionPkg {
	if obj.Connection.Sql == nil {
		obj = BulkConnectionPkg{}
		obj.Connection.Sql = make(map[string]*gorm.DB)
	}
	return &obj
}

func (h *BulkConnectionPkg) AddBulkConnection(prop string, value *gorm.DB) bool {
	if _, ok := h.Connection.Sql[prop]; ok {
		return false
	}
	h.Connection.Sql[prop] = value
	return true
}

func (h *BulkConnectionPkg) GetBulkConnection(prop string) (bool, *gorm.DB) {
	val, ok := h.Connection.Sql[prop]
	return ok, val
}
