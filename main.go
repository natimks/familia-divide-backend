package main

import (
	"fmt"
	"go-workshop/src"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Use(cors.Default())

	health := engine.Group("/")

	usersRoute := engine.Group("/users")
	familyRoute := engine.Group("/family")
	incomesRoute := engine.Group("/incomes")
	expensesRoute := engine.Group("/expenses")
	dashboardRoute := engine.Group("/dashboard")

	//user
	usersRoute.GET("/", src.GetUsers)
	usersRoute.GET("/:cpf", src.GetUserByCpf)
	usersRoute.POST("/", src.PostUser)
	usersRoute.PUT("/:id", src.PutUser)
	usersRoute.DELETE("/:id", src.DeleteUser)

	//family
	familyRoute.PUT("/:cpf/:email", src.PutAddRelative)
	familyRoute.DELETE("/:id/:id_relative", src.DeleteRelative)

	//incomes
	incomesRoute.GET("/:id", src.GetIncomesByUserId)
	incomesRoute.POST("/", src.PostIncome)
	incomesRoute.DELETE("/:id", src.DeleteIncome)

	//expenses
	expensesRoute.POST("/", src.PostExpense)
	expensesRoute.DELETE("/:id", src.DeleteExpense)

	//dashboard
	dashboardRoute.GET("/:id", src.GetStatisticsDashboard)

	src.AutoMigration()

	health.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Go healthy!",
		})
	})

	engine.Run(fmt.Sprintf(":%v", 8088))
}
