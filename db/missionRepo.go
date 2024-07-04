package db

import (
	"fmt"
	"spyingCats/models"
)

type MissionRepo struct {}

func (repo MissionRepo) GetAllMissions() ([]*models.Mission, error) {
    var missions []*models.Mission
    err := DB.Preload("Cat").Preload("Targets").Preload("Targets.Notes").Find(&missions).Error
    return missions, err
}

func (repo MissionRepo) Create(mission *models.Mission) error {
    return DB.Create(mission).Error
}

func (repo MissionRepo) GetById(id int) (*models.Mission, error) {
    var mission models.Mission
    err := DB.Preload("Cat").Preload("Targets").Preload("Targets.Notes").First(&mission, id).Error
    fmt.Printf("is cat nil? %v\n", mission.Cat == nil)
    return &mission, err
}

func (repo MissionRepo) DeleteById(id int) (error) {
    mission, err := repo.GetById(id)
    if err != nil {
        return fmt.Errorf("coulnd't delete mission: %v", err)
    }
    if mission.Cat != nil {
        return fmt.Errorf("coulnd't delete mission since it's already assigned to a cat with id: %d", mission.Cat.Id)
    }
    return DB.Delete(mission).Error
}

func (repo MissionRepo) AssignCat(missionId, catId int) (*models.Mission, error) {
    mission, err := repo.GetById(missionId)
    if err != nil {
        return nil, fmt.Errorf("can't find mission with id: %d", missionId)
    }

    catsRepo := CatsRepo{}
    cat, err := catsRepo.GetById(catId)
    if err != nil {
        return nil, fmt.Errorf("can't find cat with id: %d", catId)
    }

    mission.Cat = cat
    
    return mission, DB.Save(mission).Error
}

func (repo MissionRepo) CompleteMission(id int) (*models.Mission, error) {
    mission, err := repo.GetById(id)
    if err != nil {
        return nil, fmt.Errorf("can't find mission with id: %d", id)
    }

    mission.Complete = true

    return mission, DB.Save(mission).Error
}

func (repo MissionRepo) AddTarget(id int, target *models.Target) (*models.Mission, error) {
    mission, err := repo.GetById(id)
    if err != nil {
        return nil, fmt.Errorf("can't find mission with id: %d", id)
    }

    if mission.Complete {
        return nil, fmt.Errorf("couldn't add target to mission since it's already completed")
    }

    mission.Targets = append(mission.Targets, target)

    return mission, DB.Save(mission).Error
}

func (repo MissionRepo) RemoveTarget(missionId, targetId int) (*models.Mission, error) {
    mission, err := repo.GetById(missionId)
    if err != nil {
        return nil, fmt.Errorf("couldn't find mission with id: %d", missionId)
    }

    targetRepo := TargetRepo{}
    target, err := targetRepo.GetById(targetId)
    if err != nil {
        return nil, fmt.Errorf("couldn't find target with id: %d", targetId)
    }

    if target.Complete {
        return nil, fmt.Errorf("couldn't remove target from mission since it's already completed")
    }

    for i, target := range mission.Targets {
        if target.Id == targetId {
            mission.Targets = append(mission.Targets[:i], mission.Targets[i+1:]...)
        }
    }

    err = targetRepo.DeleteById(targetId)
    if err != nil {
        return nil, fmt.Errorf("couldn't delete target with id: %d: %v", targetId, err)
    }

    return mission, DB.Save(mission).Error
}

