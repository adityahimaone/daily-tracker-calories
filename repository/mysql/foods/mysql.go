package foods

import (
	"daily-tracker-calories/bussiness/foods"
	"gorm.io/gorm"
)

type repositoryFoods struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) foods.Repository {
	return &repositoryFoods{
		DB: db,
	}
}

func (repository repositoryFoods) GetFoodByName(name string) (*foods.Domain, error) {
	panic("implement me")
}

func (repository repositoryFoods) Insert(food *foods.Domain) (*foods.Domain, error) {
	recordFood := fromDomain(*food)
	if err := repository.DB.Create(&recordFood).Error; err != nil {
		return &foods.Domain{}, err
	}
	result := toDomain(recordFood)
	return &result, nil
}
