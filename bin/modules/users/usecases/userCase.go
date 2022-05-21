package usecases

import (
	"context"
	"errors"

	"github.com/Manusiabodoh4/bulk-generic/bin/entity"
	helper "github.com/Manusiabodoh4/bulk-generic/bin/modules/users/helper/connection"
	models "github.com/Manusiabodoh4/bulk-generic/bin/modules/users/models/domain"
	"github.com/Manusiabodoh4/bulk-generic/bin/modules/users/repositories"
	"github.com/Manusiabodoh4/bulk-generic/bin/pkg/databases"
	"github.com/google/uuid"
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
	queryRes := <-h.userRepo.FindAll(ctx)
	if queryRes.Error != nil {
		result.Error = queryRes.Error
		return result
	}
	resUsers := queryRes.Data.([]models.Users)
	if total := len(resUsers); total <= 0 {
		return result
	}
	for _, itemUser := range resUsers {
		h.bulkConnection.AddBulkConnection(itemUser.ID, helper.CreateConnectionSQL(ctx, &itemUser))
	}
	result.Data = map[string]interface{}{
		"users":      resUsers,
		"total":      len(resUsers),
		"connection": h.bulkConnection.Connection.Sql,
	}
	return result
}

func (h *UsecaseUsersImpl) Register(ctx context.Context, data *models.RegisterRequest) entity.TemplateChannelResponse {
	var (
		result       entity.TemplateChannelResponse
		tmpModelUser models.Users
		userID       string
	)
	userID = uuid.New().String()
	tmpModelUser = models.Users{
		ID:       userID,
		Tipe:     data.Tipe,
		Host:     data.Host,
		Username: data.Username,
		Password: data.Password,
		Dbname:   data.Dbname,
		Port:     data.Port,
	}
	queryRes := <-h.userRepo.InsertOne(ctx, &tmpModelUser)
	if queryRes.Error != nil {
		result.Error = queryRes.Error
		return result
	}
	sta := h.bulkConnection.AddBulkConnection(tmpModelUser.ID, helper.CreateConnectionSQL(ctx, &tmpModelUser))
	if !sta {
		result.Error = errors.New("cannot add connection to bulk manager connection")
		return result
	}
	result.Data = queryRes.Data
	return result
}

func (h *UsecaseUsersImpl) GetDetailWithID(ctx context.Context, data *models.GetDetailRequest) entity.TemplateChannelResponse {
	var result entity.TemplateChannelResponse
	query := "id = @id"
	parameter := map[string]interface{}{
		"id": data.ID,
	}
	queryRes := <-h.userRepo.FindOne(ctx, query, parameter)
	if queryRes.Error != nil {
		result.Error = queryRes.Error
		return result
	}
	result.Data = queryRes.Data
	return result
}

func (h *UsecaseUsersImpl) UpdateUsers(ctx context.Context, id string, data *models.UpdateUsersRequest) entity.TemplateChannelResponse {
	var (
		result    entity.TemplateChannelResponse
		userModel models.Users
	)

	userModel = models.Users{
		ID:       id,
		Tipe:     data.Tipe,
		Host:     data.Host,
		Username: data.Username,
		Password: data.Password,
		Dbname:   data.Dbname,
		Port:     data.Port,
	}

	queryRes := <-h.userRepo.UpdateOneWithID(ctx, &userModel)
	if queryRes.Error != nil {
		result.Error = queryRes.Error
	}
	result.Data = queryRes.Data
	return result
}
