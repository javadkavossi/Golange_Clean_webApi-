package usecase

import (
	"context"

	"github.com/javadkavossi/Golange_Clean_webApi/src/config"
	"github.com/javadkavossi/Golange_Clean_webApi/src/domain/filter"
	model "github.com/javadkavossi/Golange_Clean_webApi/src/domain/model"
	"github.com/javadkavossi/Golange_Clean_webApi/src/domain/repository"
	"github.com/javadkavossi/Golange_Clean_webApi/src/usecase/dto"
)

type CountryUsecase struct {
	base *BaseUsecase[model.Country, dto.Name, dto.Name, dto.Country]
}

func NewCountryUsecase(cfg *config.Config, repository repository.CountryRepository) *CountryUsecase {
	return &CountryUsecase{
		base: NewBaseUsecase[model.Country, dto.Name, dto.Name, dto.Country](cfg, repository),
	}
}

// Create
func (u *CountryUsecase) Create(ctx context.Context, req dto.Name) (dto.Country, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *CountryUsecase) Update(ctx context.Context, id int, req dto.Name) (dto.Country, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CountryUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CountryUsecase) GetById(ctx context.Context, id int) (dto.Country, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CountryUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.Country], error) {
	return s.base.GetByFilter(ctx, req)
}
