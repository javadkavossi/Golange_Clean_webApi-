package usecase

import (
	"context"

	"github.com/javadkavossi/Golange_Clean_webApi/src/config"
	"github.com/javadkavossi/Golange_Clean_webApi/src/domain/filter"
	model "github.com/javadkavossi/Golange_Clean_webApi/src/domain/model"
	"github.com/javadkavossi/Golange_Clean_webApi/src/domain/repository"
	"github.com/javadkavossi/Golange_Clean_webApi/src/usecase/dto"
)

type CarTypeUsecase struct {
	base *BaseUsecase[model.CarType, dto.Name, dto.Name, dto.IdName]
}

func NewCarTypeUsecase(cfg *config.Config, repository repository.CarTypeRepository) *CarTypeUsecase {
	return &CarTypeUsecase{
		base: NewBaseUsecase[model.CarType, dto.Name, dto.Name, dto.IdName](cfg, repository),
	}
}

// Create
func (u *CarTypeUsecase) Create(ctx context.Context, req dto.Name) (dto.IdName, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *CarTypeUsecase) Update(ctx context.Context, id int, req dto.Name) (dto.IdName, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarTypeUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarTypeUsecase) GetById(ctx context.Context, id int) (dto.IdName, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarTypeUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.IdName], error) {
	return s.base.GetByFilter(ctx, req)
}
