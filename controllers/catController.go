package controllers

import (
    "fmt"
    "net/http"
    "strconv"

    "spyingCats/validation"
    "spyingCats/logger"
    "spyingCats/db"
    "spyingCats/models"

    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

func GetAllCats(c *gin.Context) {
    catsRepo := db.CatsRepo{}
    cats, err := catsRepo.GetAllCats()
    if err != nil {
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error("can't get cats")
        c.JSON(http.StatusInternalServerError, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, cats)
}

func NewCat(c *gin.Context) {
    var cat models.Cat    
    if err := c.ShouldBindJSON(&cat); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "wrong request payload"})
        return
    }

    valid, err := validation.IsValidBreed(cat.Breed) 
    if err != nil {
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error("error while validating breed")
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("couldn't validate breed: %v", err)})
        return
    }
    if !valid {
        logger.Log.WithFields(logrus.Fields {
            "cat": cat,
        }).Error("invalid cat breed")
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cat breed"})
        return
    }

    catsRepo := db.CatsRepo{}
    err = catsRepo.Create(&cat)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, cat)
}

func GetCatById(c *gin.Context) {
    id := checkId(c)

    catsRepo := db.CatsRepo{}
    cat, err := catsRepo.GetById(id)
    if err != nil {
        errMsg := fmt.Sprintf("couldn't get cat with id: %d", id)
        logger.Log.WithFields(logrus.Fields{
            "error": err.Error(),
        }).Error(errMsg)
        c.JSON(http.StatusNotFound, gin.H{"error": errMsg})
        return
    }

    c.JSON(http.StatusOK, cat)
}

func DeleteCatById(c *gin.Context) {
    id := checkId(c)

    catsRepo := db.CatsRepo{}
    err := catsRepo.DeleteById(id)
    if err != nil {
        errMsg := fmt.Sprintf("couldn't delete cat with id: %d", id)
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error(errMsg)
        c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
        return
    }

    logger.Log.Info(fmt.Sprintf("deleted cat with id %d", id))
}

func UpdateCatSalary(c *gin.Context) {
    id := checkId(c)


    var newSalaryObj struct {
        Salary int `json:"salary"`
    }
    if err := c.ShouldBindJSON(&newSalaryObj); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "wrong request payload"})
        return
    }

    catsRepo := db.CatsRepo{}
    updatedCat, err := catsRepo.UpdateSalary(id, newSalaryObj.Salary)
    if err != nil {
        errMsg := fmt.Sprintf("couldn't update cat with id: %d", id)
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error(errMsg)
        c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
        return
    }

    c.JSON(http.StatusOK, updatedCat)
}


func checkId(c *gin.Context) int {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        logger.Log.WithFields(logrus.Fields {
            "error": err.Error(),
        }).Error("invalid ID")
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
        return -1
    }
    return id
}
