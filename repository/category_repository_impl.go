package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/juandaantoniusapakpahan/go-restful-api/helper"
	"github.com/juandaantoniusapakpahan/go-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	queryScript := "insert into category(name) values(?)"
	resutl, err := tx.ExecContext(ctx, queryScript, category.Name)

	helper.ErrorHandle(err)

	id, err := resutl.LastInsertId()
	helper.ErrorHandle(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	queryScript := "update category set name = ?  where id = ?"
	_, err := tx.ExecContext(ctx, queryScript, category.Name, category.Id)
	helper.ErrorHandle(err)
	return category
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error) {
	queryScript := "SELECT id, name from category where id = ?"
	row, err := tx.QueryContext(ctx, queryScript, id)
	helper.ErrorHandle(err)
	cty := domain.Category{}

	if row.Next() {
		err := row.Scan(&cty.Id, &cty.Name)
		helper.ErrorHandle(err)
		return cty, nil
	} else {
		return cty, errors.New("category is not found!")
	}
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	queryScript := "DELETE FROM  category WHERE id = ?"
	_, err := tx.ExecContext(ctx, queryScript, category.Id)
	helper.ErrorHandle(err)

}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	queryScript := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, queryScript)
	helper.ErrorHandle(err)
	categories := []domain.Category{}
	for rows.Next() {
		category := domain.Category{}
		err = rows.Scan(&category.Id, &category.Name)
		categories = append(categories, category)
	}
	return categories
}

// Delete(context context.Context, tx *sql.Tx, category domain.Category) error
// FindAll(context context.Context, tx *sql.Tx) ([]domain.Category, error)
