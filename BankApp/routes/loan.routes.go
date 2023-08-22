package routes

import (
	"bankDemo/controllers"

	"github.com/gin-gonic/gin"
)

func LoanRoute(router *gin.Engine, controller controllers.LoanController) {
	router.POST("/api/profile/createloan", controller.CreateLoan)
	router.GET("/api/profile/getloan/:id", controller.GetLoanById)
	router.PUT("/api/profile/updateloan/:id", controller.UpdateLoanById)
	router.DELETE("/api/profile/deleteloan/:id", controller.DeleteLoanById)
	//router.POST("/api/profile/createMany", controller.Createl)

}