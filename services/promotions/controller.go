package promotions

import (
	"fmt"
	"net/http"
	"promotions-app/common/helpers"

	"github.com/gin-gonic/gin"
)

var service *Service

func init() {
	service = new(Service)
}

type Controller struct{}

func (_ Controller) Find(c *gin.Context) {

	id := c.Param("id")
	res, ok := service.Find(id)
	if ok {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"code": "RECORD_NOT_FOUND", "message": "Record not found"})
		c.Abort()
	}
	return
}

func (_ Controller) Upload(c *gin.Context) {

	file, err := c.FormFile("file")

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"code": "INVALID_FILE", "message": "Invalid file"})
		c.Abort()
		return
	}

	filePath := helpers.UniqueTmpPath(file.Filename)

	c.SaveUploadedFile(file, filePath)

	ok := service.Import(filePath)

	if ok {
		c.JSON(http.StatusOK, gin.H{"success": 1})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "INTERNAL_SERVER_ERROR", "message": "Internal server error"})
		c.Abort()
	}
	return
}
