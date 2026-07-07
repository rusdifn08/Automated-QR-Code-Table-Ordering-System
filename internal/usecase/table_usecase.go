package usecase

import (
	"myapp/internal/domain"
)

type tableUsecase struct {
	tableRepo domain.TableRepository
}

func NewTableUsecase(repo domain.TableRepository) domain.TableUsecase {
	return &tableUsecase{
		tableRepo: repo,
	}
}

func (t *tableUsecase) GetTableByNumber(tableNumber int) (*domain.Table, error) {
	return t.tableRepo.GetByNumber(tableNumber)
}
