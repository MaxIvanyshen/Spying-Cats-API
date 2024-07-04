package controllers

import (
    "fmt"
    "net/http"
    "strconv"

    "spyingCats/logger"
    "spyingCats/db"
    "spyingCats/models"

    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

func AddTargetToMission(c *gin.Context) {
    idParam := c.Param("missionId")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error("invalid ID")
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
    }
    
    var target *models.Target
    if err := c.ShouldBindJSON(&target); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "wrong request payload"})
        return
    }

    missionsRepo := db.MissionRepo{}
    updatedMission, err := missionsRepo.AddTarget(id, target)
    if err != nil {
        errMsg := fmt.Sprintf("couldn't add target to mission: %v", err)
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error(errMsg)
        c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
        return
    }

    c.JSON(http.StatusOK, updatedMission)
}

func DeleteTargetFromMission(c *gin.Context) {
    missionIdParam := c.Param("missionId")
    missionId, err := strconv.Atoi(missionIdParam)
    if err != nil {
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error("invalid mission ID")
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mission ID"})
        return
    }

    targetIdParam := c.Param("targetId")
    targetId, err := strconv.Atoi(targetIdParam)
    if err != nil {
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error("invalid cat ID")
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cat ID"})
        return
    }

    missionsRepo := db.MissionRepo{}
    updatedMission, err := missionsRepo.RemoveTarget(missionId, targetId)
    if err != nil {
        errMsg := fmt.Sprintf("couldn't remove target from mission: %v", err)
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error(errMsg)
        c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
    }

    c.JSON(http.StatusOK, updatedMission)
}

func CompleteTarget(c *gin.Context) {
    id := checkId(c)

    targetRepo := db.TargetRepo{}
    updatedTarget, err := targetRepo.Complete(id)
    if err != nil {
        errMsg := fmt.Sprintf("couldn't mark tarket as complete: %v", err)
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error(errMsg)
        c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
        return
    }

    c.JSON(http.StatusOK, updatedTarget)
}

func UpdateNotes(c *gin.Context) {
    id := checkId(c)
    targetRepo := db.TargetRepo{}
    var newNotesObj struct {
        Notes string `json:"notes"`
    }
    if err := c.ShouldBindJSON(&newNotesObj); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "wrong request payload"})
        return
    }
    updatedTarget, err := targetRepo.UpdateNotes(id, newNotesObj.Notes)
    if err != nil {
        errMsg := fmt.Sprintf("couldn't mark tarket as complete: %v", err)
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error(errMsg)
        c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
        return
    }

    c.JSON(http.StatusOK, updatedTarget)
}

