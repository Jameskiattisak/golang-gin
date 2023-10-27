package service

import (
	"net/http"

	"itmx-test/database"
	"itmx-test/model"

	"github.com/gin-gonic/gin"
)

func GetAllCustomers(c *gin.Context) {
	var customers []model.Customers
	database.DB.Find(&customers)
	c.JSON(http.StatusOK, customers)
}

func GetACustomers(c *gin.Context) {
	id := c.Param("id")
	var customers []model.Customers
	database.DB.First(&customers, id)
	c.JSON(http.StatusOK, customers)
}

func CreateCustomers(c *gin.Context) {
	var customers model.Customers
	if err := c.ShouldBindJSON(&customers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := database.DB.Create(&customers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, customers)
}

func UpdateCustomers(c *gin.Context) {
	id := c.Param("id")
	var customer model.Customers

	if result := database.DB.First(&customer, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func DeleteCustomers(c *gin.Context) {
	id := c.Param("id")
	var customers []model.Customers
	database.DB.First(&customers, id)
	database.DB.Delete(&customers)
	c.JSON(http.StatusOK, customers)
}
