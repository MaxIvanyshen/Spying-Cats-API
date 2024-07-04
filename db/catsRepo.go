package db

import (
	"fmt"
	"spyingCats/models"
)

type CatsRepo struct {}

func (repo CatsRepo) GetAllCats() ([]*models.Cat, error) {
    var cats []*models.Cat
    err := DB.Find(&cats).Error
    return cats, err
}

func (repo CatsRepo) Create(cat *models.Cat) error {
    return DB.Create(cat).Error
}

func (repo CatsRepo) GetById(id int) (*models.Cat, error) {
    var cat models.Cat
    err := DB.First(&cat, id).Error
    return &cat, err
}

func (repo CatsRepo) DeleteById(id int) (error) {
    return DB.Delete(&models.Cat{}, id).Error
}

func (repo CatsRepo) UpdateSalary(id int, salary int) (*models.Cat, error) {
    cat, err := repo.GetById(id)
    if err != nil {
        return nil, fmt.Errorf("couldn't update cat: %v", err)
    }

    cat.Salary = salary
    return cat, DB.Save(cat).Error
}
