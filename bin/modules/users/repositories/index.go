package repositories

import (
	"context"

	"github.com/Manusiabodoh4/bulk-generic/bin/entity"
	models "github.com/Manusiabodoh4/bulk-generic/bin/modules/users/models/domain"
)

type UserPostgres interface {
	InsertOne(ctx context.Context, data *models.Users) <-chan entity.TemplateChannelResponse
	UpdateOneWithID(ctx context.Context, data *models.Users) <-chan entity.TemplateChannelResponse
	FindOne(ctx context.Context, query string, parameter map[string]interface{}) <-chan entity.TemplateChannelResponse
	FindAll(ctx context.Context) <-chan entity.TemplateChannelResponse
}
