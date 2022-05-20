package usecases

import (
	"context"

	"github.com/Manusiabodoh4/bulk-generic/bin/entity"
	models "github.com/Manusiabodoh4/bulk-generic/bin/modules/users/models/domain"
)

type UsecaseUsers interface {
	SetupConnection(ctx context.Context) entity.TemplateChannelResponse
	Register(ctx context.Context, data *models.RegisterRequest) entity.TemplateChannelResponse
	GetDetailWithID(ctx context.Context, data *models.GetDetailRequest) entity.TemplateChannelResponse
	UpdateUsers(ctx context.Context, data *models.UpdateUsersRequest) entity.TemplateChannelResponse
}
