package src

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetUser find incomes by user ID
func GetIncomesByUserId(c *gin.Context) {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		checkErr(err, 400, c)
		return
	}

	selectedUser, db := selectUserID(ID, c)

	var incomes []income
	db.Model(&selectedUser).Related(&incomes)
	selectedUser.Incomes = incomes

	db.Close()

	c.JSON(200, selectedUser)
}

//PostIncome create income
func PostIncome(c *gin.Context) {
	var newIncome income
	if err := c.ShouldBindJSON(&newIncome); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db := DBConnect()
	db.Create(&newIncome)

	db.Close()

	c.JSON(201, newIncome)
}

//DeleteIncome delete income by ID
func DeleteIncome(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		checkErr(err, 400, c)
		return
	}

	income, db := selectIncomeByID(ID, c)
	defer db.Close()
	if income.ID > 0 {
		db.Delete(&income)
	}

	c.JSON(204, nil)
}
