package repositories

import (
	"context"

	"github.com/Manusiabodoh4/bulk-generic/bin/entity"
	models "github.com/Manusiabodoh4/bulk-generic/bin/modules/users/models/domain"
	"gorm.io/gorm"
)

type UserPostgresImpl struct {
	db    *gorm.DB
	table string
}

func NewUserPostgresImpl(db *gorm.DB) UserPostgres {
	return &UserPostgresImpl{
		db:    db,
		table: "users",
	}
}

func (h *UserPostgresImpl) InsertOne(ctx context.Context, data *models.Users) <-chan entity.TemplateChannelResponse {

	output := make(chan entity.TemplateChannelResponse)

	go func() {
		defer close(output)
		result := h.db.Create(&data)
		if result.Error != nil {
			output <- entity.TemplateChannelResponse{Error: result.Error}
		}
		output <- entity.TemplateChannelResponse{Data: data}
	}()

	return output

}
func (h *UserPostgresImpl) UpdateOneWithID(ctx context.Context, data *models.Users) <-chan entity.TemplateChannelResponse {

	output := make(chan entity.TemplateChannelResponse)

	go func() {
		defer close(output)
		result := h.db.Table(h.table).Where("id = ?", data.ID).Updates(map[string]interface{}{
			"host":     data.Host,
			"username": data.Username,
			"password": data.Password,
			"dbname":   data.Dbname,
			"port":     data.Port,
		})
		if result.Error != nil {
			output <- entity.TemplateChannelResponse{Error: result.Error}
		}
		output <- entity.TemplateChannelResponse{Data: result.RowsAffected}
	}()

	return output

}
func (h *UserPostgresImpl) FindOne(ctx context.Context, query string, parameter map[string]interface{}) <-chan entity.TemplateChannelResponse {

	output := make(chan entity.TemplateChannelResponse)

	go func() {
		defer close(output)

		var data models.Users
		result := h.db.Table(h.table).Model(&data).Where(query, parameter).Find(&data)
		if result.Error != nil {
			output <- entity.TemplateChannelResponse{
				Error: result.Error,
			}
		}
		output <- entity.TemplateChannelResponse{Data: data}
	}()

	return output

}

func (h *UserPostgresImpl) FindAll(ctx context.Context) <-chan entity.TemplateChannelResponse {

	output := make(chan entity.TemplateChannelResponse)

	go func() {
		defer close(output)
		var users []models.Users
		result := h.db.Table(h.table).Find(&users)
		if result.Error != nil {
			output <- entity.TemplateChannelResponse{
				Error: result.Error,
			}
		}
		output <- entity.TemplateChannelResponse{Data: users}
	}()

	return output

}
