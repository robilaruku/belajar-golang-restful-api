package repository

import (
	"context"
	"database/sql"

	"robi/belajar-golang-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
  
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx sql.Tx, category domain.Category) domain.Category {
	panic("not implemented") // TODO: Implement
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx sql.Tx, cateogry domain.Category) domain.Category {
	panic("not implemented") // TODO: Implement
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx sql.Tx, cateogry domain.Category) {
	panic("not implemented") // TODO: Implement
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx sql.Tx, categoryId int) domain.Category {
	panic("not implemented") // TODO: Implement
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx sql.Tx) []domain.Category {
	panic("not implemented") // TODO: Implement
}




