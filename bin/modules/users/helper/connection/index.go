package helper

import (
	"context"

	models "github.com/Manusiabodoh4/bulk-generic/bin/modules/users/models/domain"
	"gorm.io/gorm"
)

var listFunc map[string]interface{}

//add all function connection in helper
func NewCreateConnectionHelper() {
	if listFunc == nil {
		listFunc = make(map[string]interface{})
	}
	listFunc["pgsql"] = CreateConnectionPostgre
}

func CreateConnectionSQL(ctx context.Context, data *models.Users) *gorm.DB {
	return listFunc[data.Tipe].(func(context.Context, *models.Users) *gorm.DB)(ctx, data)
}
