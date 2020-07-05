package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (idb *InDB) GetArticles(c *gin.Context) {
	var (
		articles []structs.Article
		payload  gin.H
	)

	result := idb.DB.
		Set("gorm:auto_preload", true).
		Find(&articles)

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
			"data":    articles,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) GetArticleByID(c *gin.Context) {
	var (
		article structs.Article
		payload gin.H
	)

	id := c.Param("id")
	result := idb.DB.
		Set("gorm:auto_preload", true).
		Where("id = ?", id).
		First(&article)

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
			"data":    article,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) CreateArticle(c *gin.Context) {
	var (
		article structs.Article
		payload gin.H
	)

	if err := c.BindJSON(&article); err != nil {
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
		Create(&article)

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
			"data":    article,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) UpdateArticleByID(c *gin.Context) {
	var (
		article structs.Article
		payload gin.H
	)

	id := c.Param("id")

	//Make Sure Selection is Available
	result := idb.DB.
		Set("gorm:auto_preload", true).
		Where("id = ?", id).
		First(&article)

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

	if err := c.BindJSON(&article); err != nil {
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
		Save(&article)

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
			"data":    article,
		}
	}

	c.JSON(http.StatusOK, payload)
}

func (idb *InDB) DeleteArticleByID(c *gin.Context) {
	var (
		article structs.Article
		payload gin.H
	)

	id := c.Param("id")
	result := idb.DB.
		Set("gorm:auto_preload", true).
		Where("id = ?", id).
		First(&article)

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
		Delete(&article); result.Error != nil {

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
