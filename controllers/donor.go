package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (idb *InDB) GetDonors(c *gin.Context) {
	var (
		donors  []structs.Donor
		payload gin.H
	)

	result := idb.DB.
		Set("gorm:auto_preload", true).
		Preload("Article").
		Find(&donors)

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
			"data":    donors,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) GetDonorByID(c *gin.Context) {
	var (
		donor   structs.Donor
		payload gin.H
	)

	id := c.Param("id")
	result := idb.DB.
		Set("gorm:auto_preload", true).
		Where("id = ?", id).
		First(&donor)

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
			"data":    donor,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) GetDonorsByArticleID(c *gin.Context) {
	var (
		donor   []structs.Donor
		payload gin.H
	)

	id := c.Param("id")
	result := idb.DB.
		Set("gorm:auto_preload", true).
		Where("article_id = ? AND status != 0", id).
		Order("id desc").
		Find(&donor)

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
			"data":    donor,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) CreateDonor(c *gin.Context) {
	var (
		donor   structs.Donor
		payload gin.H
	)

	if err := c.BindJSON(&donor); err != nil {
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
		Create(&donor)

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
			"data":    donor,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) UpdateDonorByID(c *gin.Context) {
	var (
		donor   structs.Donor
		payload gin.H
	)

	id := c.Param("id")

	//Make Sure Selection is Available
	result := idb.DB.
		Set("gorm:auto_preload", true).
		Where("id = ?", id).
		First(&donor)

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

	if err := c.BindJSON(&donor); err != nil {
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
		Save(&donor)

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
			"data":    donor,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) DeleteDonorByID(c *gin.Context) {
	var (
		donor   structs.Donor
		payload gin.H
	)

	id := c.Param("id")
	result := idb.DB.
		Set("gorm:auto_preload", true).
		Where("id = ?", id).
		First(&donor)

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
		Delete(&donor); result.Error != nil {

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
