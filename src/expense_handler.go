package src

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetUser find expenses by user ID
func GetExpensesByUserId(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		checkErr(err, 400, c)
		return
	}

	selectedUser, db := selectUserID(ID, c)
	defer db.Close()

	var expenses []expense
	db.Model(&selectedUser).Related(&expenses)
	selectedUser.Expenses = expenses

	c.JSON(200, selectedUser)
}

//PostExpense create expense
func PostExpense(c *gin.Context) {

	var newExpense expense
	if err := c.ShouldBindJSON(&newExpense); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := DBConnect()
	defer db.Close()

	db.Create(&newExpense)

	c.JSON(201, newExpense)
}

//DeleteExpense delete expense by ID
func DeleteExpense(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		checkErr(err, 400, c)
		return
	}

	expense, db := selectExpenseByID(ID, c)
	defer db.Close()

	if expense.ID > 0 {
		db.Delete(&expense)
	}

	c.JSON(204, nil)
}
