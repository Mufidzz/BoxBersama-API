package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (idb *InDB) GetBankAccounts(c *gin.Context) {
	var (
		bankAccounts []structs.BankAccount
		payload      gin.H
	)

	result := idb.DB.
		Set("gorm:auto_preload", true).
		Find(&bankAccounts)

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
			"data":    bankAccounts,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) GetBankAccountByID(c *gin.Context) {
	var (
		bankAccount structs.BankAccount
		payload     gin.H
	)

	id := c.Param("id")
	result := idb.DB.
		Set("gorm:auto_preload", true).
		Where("id = ?", id).
		First(&bankAccount)

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
			"data":    bankAccount,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) CreateBankAccount(c *gin.Context) {
	var (
		bankAccount structs.BankAccount
		payload     gin.H
	)

	if err := c.BindJSON(&bankAccount); err != nil {
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
		Create(&bankAccount)

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
			"data":    bankAccount,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) UpdateBankAccountByID(c *gin.Context) {
	var (
		bankAccount structs.BankAccount
		payload     gin.H
	)

	id := c.Param("id")

	//Make Sure Selection is Available
	result := idb.DB.
		Set("gorm:auto_preload", true).
		Where("id = ?", id).
		First(&bankAccount)

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

	if err := c.BindJSON(&bankAccount); err != nil {
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
		Save(&bankAccount)

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
			"data":    bankAccount,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) DeleteBankAccountByID(c *gin.Context) {
	var (
		bankAccount structs.BankAccount
		payload     gin.H
	)

	id := c.Param("id")
	result := idb.DB.
		Set("gorm:auto_preload", true).
		Where("id = ?", id).
		First(&bankAccount)

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
		Delete(&bankAccount); result.Error != nil {

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
