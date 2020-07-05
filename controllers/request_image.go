package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (idb *InDB) GetRequestImages(c *gin.Context) {
	var (
		requestImages []structs.RequestImage
		payload       gin.H
	)

	result := idb.DB.
		Set("gorm:auto_preload", true).
		Find(&requestImages)

	if result.Error != nil {
		if result.RecordNotFound() {
			payload = gin.H{
				"message": "Record Not Found",
				"data":    nil,
			}
		} else {
			payload = gin.H{
				"message": "Request Had Some Error",
				"errors":  result.GetErrors(),
			}
		}
	} else {
		payload = gin.H{
			"message": "Success",
			"data":    requestImages,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) GetRequestImageByID(c *gin.Context) {
	var (
		requestImage structs.RequestImage
		payload      gin.H
	)

	id := c.Param("id")
	result := idb.DB.
		Set("gorm:auto_preload", true).
		Where("id = ?", id).
		First(&requestImage)

	if result.Error != nil {
		if result.RecordNotFound() {
			payload = gin.H{
				"message": "Record Not Found",
				"data":    nil,
			}
		} else {
			payload = gin.H{
				"message": "Request Had Some Error",
				"errors":  result.GetErrors(),
			}
		}
	} else {
		payload = gin.H{
			"message": "Success",
			"data":    requestImage,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) CreateRequestImage(c *gin.Context) {
	var (
		requestImage structs.RequestImage
		payload      gin.H
	)

	if err := c.BindJSON(&requestImage); err != nil {
		payload = gin.H{
			"message":   "Request Had Some Error",
			"reference": "JSON Binding Failed",
			"errors":    err.Error(),
		}

		c.JSON(http.StatusBadRequest, payload)
		return
	}

	result := idb.DB.
		Set("gorm:auto_preload", true).
		Create(&requestImage)

	if result.Error != nil {
		if result.RecordNotFound() {
			payload = gin.H{
				"message": "Record Not Found",
				"data":    nil,
			}
		} else {
			payload = gin.H{
				"message": "Request Had Some Error",
				"errors":  result.GetErrors(),
			}
		}
	} else {
		payload = gin.H{
			"message": "Success",
			"data":    requestImage,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) UpdateRequestImageByID(c *gin.Context) {
	var (
		requestImage structs.RequestImage
		payload      gin.H
	)

	id := c.Param("id")

	//Make Sure Selection is Available
	result := idb.DB.
		Set("gorm:auto_preload", true).
		Where("id = ?", id).
		First(&requestImage)

	if result.Error != nil {
		if result.RecordNotFound() {
			payload = gin.H{
				"message": "Record Not Found",
				"data":    nil,
			}
		} else {
			payload = gin.H{
				"message": "Request Had Some Error",
				"errors":  result.GetErrors(),
			}
		}
		c.JSON(http.StatusOK, payload)
		return
	}

	if err := c.BindJSON(&requestImage); err != nil {
		payload = gin.H{
			"message":   "Request Had Some Error",
			"reference": "JSON Binding Failed",
			"errors":    err.Error(),
		}

		c.JSON(http.StatusBadRequest, payload)
		return
	}

	result = idb.DB.
		Set("gorm:auto_preload", true).
		Save(&requestImage)

	if result.Error != nil {
		if result.RecordNotFound() {
			payload = gin.H{
				"message": "Record Not Found",
				"data":    nil,
			}
		} else {
			payload = gin.H{
				"message": "Request Had Some Error",
				"errors":  result.GetErrors(),
			}
		}
	} else {
		payload = gin.H{
			"message": "Success",
			"data":    requestImage,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) DeleteRequestImageByID(c *gin.Context) {
	var (
		requestImage structs.RequestImage
		payload      gin.H
	)

	id := c.Param("id")
	result := idb.DB.
		Set("gorm:auto_preload", true).
		Where("id = ?", id).
		First(&requestImage)

	if result.Error != nil {
		if result.RecordNotFound() {
			payload = gin.H{
				"message": "Record Not Found",
				"data":    nil,
			}
		} else {
			payload = gin.H{
				"message": "Request Had Some Error",
				"errors":  result.GetErrors(),
			}
		}

		c.JSON(http.StatusOK, payload)
		return
	}

	if result := idb.DB.
		Set("gorm:auto_preload", true).
		Delete(&requestImage); result.Error != nil {

		payload = gin.H{
			"message": "Request Had Some Error",
			"errors":  result.GetErrors(),
		}
	} else {
		payload = gin.H{
			"message": "Success",
			"data":    nil,
		}
	}

	c.JSON(http.StatusOK, payload)
}
