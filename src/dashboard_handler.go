package src

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetStatisticsDashboard calculates the statistics of values ​​for the dashboard
func GetStatisticsDashboard(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		checkErr(err, 400, c)
		return
	}
	selectedUser, db := selectUserID(ID, c)
	defer db.Close()

	db.Preload("UsersFamily").First(&selectedUser)
	db.Preload("Incomes").First(&selectedUser)
	db.Preload("Expenses").First(&selectedUser)

	for i := 0; i < len(selectedUser.UsersFamily); i++ {
		relative := selectedUser.UsersFamily[i]
		db.Preload("Incomes").First(&relative)
		db.Preload("Expenses").First(&relative)
	}

	c.JSON(200, selectedUser)
}
