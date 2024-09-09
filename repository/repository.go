package repository

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"test-openapi/generated/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(ctx context.Context, id int) (*models.User, error) {
	u, err := models.FindUser(ctx, r.db, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) Create(ctx context.Context, u *models.User) error {
	return u.Insert(ctx, r.db, boil.Infer())
}

type SpaceRepository struct {
	db *sql.DB
}

func NewSpaceRepository(db *sql.DB) *SpaceRepository {
	return &SpaceRepository{db: db}
}

func (r *SpaceRepository) FindByID(ctx context.Context, id int) (*models.Space, error) {
	s, err := models.FindSpace(ctx, r.db, id)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (r *SpaceRepository) Create(ctx context.Context, s *models.Space) error {
	return s.Insert(ctx, r.db, boil.Infer())
}
