package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Files(c *gin.Context) {
	c.HTML(http.StatusOK, "files.html", nil)
}

func Upload(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}
