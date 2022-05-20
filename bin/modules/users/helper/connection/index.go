package helper

import (
	"context"

	models "github.com/Manusiabodoh4/bulk-generic/bin/modules/users/models/domain"
	"gorm.io/gorm"
)

var listFunc map[string]interface{}

//add all function connection in helper
func NewCreateConnectionHelper() {
	listFunc["pgsql"] = CreateConnectionPostgre
}

func CreateConnectionSQL(ctx context.Context, data *models.RegisterRequest) *gorm.DB {
	return listFunc[data.Tipe].(func(context.Context, *models.RegisterRequest) *gorm.DB)(ctx, data)
}
