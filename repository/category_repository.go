package repository

import (
	"context"
	"database/sql"

	"github.com/juandaantoniusapakpahan/go-restful-api/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(context context.Context, tx *sql.Tx, category domain.Category) domain.Category
	FindById(context context.Context, tx *sql.Tx, id int) (domain.Category, error)
	Delete(context context.Context, tx *sql.Tx, category domain.Category)
	FindAll(context context.Context, tx *sql.Tx) []domain.Category
}
