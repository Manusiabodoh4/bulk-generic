package usecases

import (
	"context"

	"github.com/Manusiabodoh4/bulk-generic/bin/entity"
	models "github.com/Manusiabodoh4/bulk-generic/bin/modules/users/models/domain"
	"github.com/Manusiabodoh4/bulk-generic/bin/modules/users/repositories"
	"github.com/Manusiabodoh4/bulk-generic/bin/pkg/databases"
)

type UsecaseUsersImpl struct {
	userRepo       repositories.UserPostgres
	bulkConnection *databases.BulkConnectionPkg
}

func NewUsecaseUsers(userRepo repositories.UserPostgres, bulkConnection databases.BulkConnectionPkg) UsecaseUsers {
	return &UsecaseUsersImpl{
		userRepo:       userRepo,
		bulkConnection: &bulkConnection,
	}
}

func (h *UsecaseUsersImpl) SetupConnection(ctx context.Context) entity.TemplateChannelResponse {
	var result entity.TemplateChannelResponse

	return result
}

func (h *UsecaseUsersImpl) Register(ctx context.Context, data *models.RegisterRequest) entity.TemplateChannelResponse {
	var result entity.TemplateChannelResponse
	return result
}

func (h *UsecaseUsersImpl) GetDetailWithID(ctx context.Context, data *models.GetDetailRequest) entity.TemplateChannelResponse {
	var result entity.TemplateChannelResponse
	return result
}

func (h *UsecaseUsersImpl) UpdateUsers(ctx context.Context, data *models.UpdateUsersRequest) entity.TemplateChannelResponse {
	var result entity.TemplateChannelResponse
	return result
}
