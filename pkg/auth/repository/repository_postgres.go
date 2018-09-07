package repository

import (
	"github.com/FernandoCagale/serverless-auth/pkg/datastore"
	"github.com/FernandoCagale/serverless-auth/pkg/entity"
)

//GormRepository in memory repo
type GormRepository struct {
	connection string
}

//NewGormRepository create new repository
func NewGormRepository(connection string) *GormRepository {
	return &GormRepository{connection}
}

//Login auth
func (r *GormRepository) Login(e *entity.Auth) error {
	db, err := datastore.NewPostgres(r.connection)
	if err != nil {
		return err
	}

	defer db.Close()

	auth := new(entity.Auth)
	if err := db.Where("username = ?", e.Username).First(&auth).Error; err != nil {
		return err
	}

	if err := auth.ValidatePassword(e.Password); err != nil {
		return err
	}
	return nil
}
