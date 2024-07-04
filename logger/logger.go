package logger

import (
    "github.com/sirupsen/logrus"
    "bytes"
    "github.com/gin-gonic/gin"
    "os"
)

var Log = logrus.New()

func InitLogger() {
    Log.SetFormatter(&logrus.JSONFormatter{})
    Log.SetOutput(os.Stdout)
    Log.SetLevel(logrus.InfoLevel)
}

type ResponseWriterCapture struct {
    gin.ResponseWriter
    Body *bytes.Buffer
}

func (rw *ResponseWriterCapture) Write(b []byte) (int, error) {
    if rw.Body != nil {
        rw.Body.Write(b)
    }
    return rw.ResponseWriter.Write(b)
}


