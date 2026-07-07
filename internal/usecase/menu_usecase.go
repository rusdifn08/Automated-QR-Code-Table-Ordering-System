package usecase

import (
	"myapp/internal/domain"
)

type menuUsecase struct {
	menuRepo domain.MenuRepository
}

func NewMenuUsecase(repo domain.MenuRepository) domain.MenuUsecase {
	return &menuUsecase{
		menuRepo: repo,
	}
}

func (m *menuUsecase) FetchAll() ([]domain.Menu, error) {
	return m.menuRepo.GetAll()
}
