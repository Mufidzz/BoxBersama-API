package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (idb *InDB) GetBanks(c *gin.Context) {
	var (
		banks   []structs.Bank
		payload gin.H
	)

	result := idb.DB.
		Set("gorm:auto_preload", true).
		Find(&banks)

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
			"data":    banks,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) GetBankByID(c *gin.Context) {
	var (
		bank    structs.Bank
		payload gin.H
	)

	id := c.Param("id")
	result := idb.DB.
		Set("gorm:auto_preload", true).
		Where("id = ?", id).
		First(&bank)

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
			"data":    bank,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) CreateBank(c *gin.Context) {
	var (
		bank    structs.Bank
		payload gin.H
	)

	if err := c.BindJSON(&bank); err != nil {
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
		Create(&bank)

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
			"data":    bank,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) UpdateBankByID(c *gin.Context) {
	var (
		bank    structs.Bank
		payload gin.H
	)

	id := c.Param("id")

	//Make Sure Selection is Available
	result := idb.DB.
		Set("gorm:auto_preload", true).
		Where("id = ?", id).
		First(&bank)

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

	if err := c.BindJSON(&bank); err != nil {
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
		Save(&bank)

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
			"data":    bank,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) DeleteBankByID(c *gin.Context) {
	var (
		bank    structs.Bank
		payload gin.H
	)

	id := c.Param("id")
	result := idb.DB.
		Set("gorm:auto_preload", true).
		Where("id = ?", id).
		First(&bank)

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
		Delete(&bank); result.Error != nil {

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
