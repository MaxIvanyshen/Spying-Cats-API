package main

import (
	"fmt"
    "net/http"
    "io"
    "bytes"

    "spyingCats/db"
    "spyingCats/controllers"
    "spyingCats/logger"

	"github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

func LoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        requestBody, err := c.GetRawData()
        if err != nil {
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
            return
        }

        logger.Log.WithFields(logrus.Fields{
            "method": c.Request.Method,
            "path":   c.Request.URL.Path,
            "query":  c.Request.URL.RawQuery,
            "body":   string(requestBody),
        }).Info("Request received")

        c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))

        responseBody := &bytes.Buffer{}
        recordedWriter := &logger.ResponseWriterCapture{ResponseWriter: c.Writer, Body: responseBody}
        c.Writer = recordedWriter

        c.Next()

        logger.Log.WithFields(logrus.Fields{
            "status": c.Writer.Status(),
            "body":   responseBody.String(),
        }).Info("Response sent")

        c.Writer = recordedWriter.ResponseWriter
    }
}

func main() {
    db.InitDB();

    logger.InitLogger();

    router := gin.Default()
    router.Use(LoggingMiddleware())

    router.GET("/cats", controllers.GetAllCats)
    router.POST("/cats", controllers.NewCat)
    router.GET("/cats/:id", controllers.GetCatById)
    router.DELETE("/cats/:id", controllers.DeleteCatById)
    router.PATCH("/cats/:id", controllers.UpdateCatSalary)

    router.GET("/missions", controllers.GetAllMissions)
    router.POST("/missions", controllers.CreateMission)
    router.GET("/missions/:id", controllers.GetMissionById)
    router.PATCH("/missions/:id", controllers.CompleteMission)
    router.PATCH("/missions/assign/:missionId/:catId", controllers.AssignCatToMission)
    router.DELETE("/missions/:id", controllers.DeleteMissionById)

    router.POST("/targets/:missionId/", controllers.AddTargetToMission)
    router.PATCH("/targets/:id/", controllers.CompleteTarget)
    router.PATCH("/targets/notes/:id", controllers.UpdateNotes)
    router.DELETE("/targets/:missionId/:targetId", controllers.DeleteTargetFromMission)

    err := router.Run("localhost:8080")
    if err != nil {
        fmt.Printf("Couldn't start the server: %v", err)
    }
}
