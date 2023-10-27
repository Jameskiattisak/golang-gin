package controller

import (
	"fmt"
	"itmx-test/service"

	"github.com/gin-gonic/gin"
)

func Api() {
	r := gin.Default()

	r.GET("/customers", service.GetAllCustomers)
	r.GET("/customers/:id", service.GetACustomers)
	r.POST("/customers", service.CreateCustomers)
	r.PUT("/customers/:id", service.UpdateCustomers)
	r.DELETE("/customers/:id", service.DeleteCustomers)

	err := r.Run(":8080")
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
	fmt.Println("Port 8080")
}
