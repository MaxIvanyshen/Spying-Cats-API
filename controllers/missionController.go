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

func GetAllMissions(c *gin.Context) {
    missionsRepo := db.MissionRepo{}
    missions, err := missionsRepo.GetAllMissions()
    if err != nil {
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error("can't get missions")
        c.JSON(http.StatusInternalServerError, gin.H{"error": err})
        return
    }

    logger.Log.WithFields(logrus.Fields{
        "missions": missions,
    }).Info("all missions fetched")

    c.JSON(http.StatusOK, missions)
}

func CreateMission(c *gin.Context) {
    var mission models.Mission    
    if err := c.ShouldBindJSON(&mission); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "wrong request payload"})
        return
    }

    missionsRepo := db.MissionRepo{}
    err := missionsRepo.Create(&mission)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, mission)
}

func GetMissionById(c *gin.Context) {
    id := checkId(c)

    missionsRepo := db.MissionRepo{}
    mission,err := missionsRepo.GetById(id)
    if err != nil {
        errMsg := fmt.Sprintf("couldn't get mission with id: %d", id)
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error(errMsg)
        c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
        return
    }

    logger.Log.WithFields(logrus.Fields{
        "mission": mission,
    }).Info("mission fetched")

    c.JSON(http.StatusOK, mission)
}

func DeleteMissionById(c *gin.Context) {
    id := checkId(c)

    missionsRepo := db.MissionRepo{}
    err := missionsRepo.DeleteById(id)
    if err != nil {
        errMsg := fmt.Sprintf("couldn't delete mission with id: %d: %v", id, err)
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error(errMsg)
        c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
        return
    }
}

func AssignCatToMission(c *gin.Context) {
    missionIdParam := c.Param("missionId")
    missionId, err := strconv.Atoi(missionIdParam)
    if err != nil {
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error("invalid mission ID")
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mission ID"})
        return
    }

    catIdParam := c.Param("catId")
    catId, err := strconv.Atoi(catIdParam)
    if err != nil {
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error("invalid cat ID")
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cat ID"})
        return
    }

    missionsRepo := db.MissionRepo{}
    updatedMission, err := missionsRepo.AssignCat(missionId, catId)
    if err != nil {
        errMsg := fmt.Sprintf("couldn't assign cat to mission: %v", err)
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error(errMsg)
        c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
        return
    }

    c.JSON(http.StatusOK, updatedMission)
}

func CompleteMission(c *gin.Context) {
    id := checkId(c)
    missionsRepo := db.MissionRepo{}
    updatedMission, err := missionsRepo.CompleteMission(id)
    if err != nil {
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error("couldn't set mission status to completed")
        c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't set mission status to completed"})
        return
    }

    c.JSON(http.StatusOK, updatedMission)
}

