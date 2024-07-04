package db

import (
	"fmt"
	"spyingCats/models"
)

type TargetRepo struct {}

func (repo TargetRepo) GetById(id int) (*models.Target, error) {
    var target models.Target
    err := DB.Preload("Notes").First(&target, id).Error
    return &target, err
}

func (repo TargetRepo) DeleteById(id int) (error) {
    target, err := repo.GetById(id)
    if err != nil {
        return fmt.Errorf("coulnd't delete target: %v", err)
    }

    for _, note := range target.Notes {
        if err := DB.Delete(note).Error; err != nil {
            return fmt.Errorf("coulnd't delete target: can't clear notes from target: %v", err)
        }
    }
    
    return DB.Delete(target).Error
}

func (repo TargetRepo) Complete(id int) (*models.Target, error) {
    target, err := repo.GetById(id)
    if err != nil {
        return nil, err
    }

    target.Complete = true

    return target, DB.Save(target).Error
}

func (repo TargetRepo) UpdateNotes(id int, notes string) (*models.Target, error) {
    target, err := repo.GetById(id)
    if err != nil {
        return nil, err
    }

    missionRepo := MissionRepo{}
    mission, err := missionRepo.GetById(target.MissionID)
    if err != nil {
        return nil, err
    }

    if mission.Complete {
        return nil, fmt.Errorf("couldn't add notes to target since mission is complete")
    }

    if target.Complete {
        return nil, fmt.Errorf("couldn't add notes to target since target is complete")
    }

    newNote := &models.Note {
        Content: notes,
    }

    target.Notes = append(target.Notes, newNote)

    return target, DB.Save(target).Error
}
