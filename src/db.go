package src

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	dbName = "familia_divide"
	dbUser = "admin"
	dbPort = "8010"
	dbHost = "127.0.0.1"
	dbPass = "admin"
)

//DBConnect is the function for open a connection with the database
func DBConnect() *gorm.DB {
	urlConn := fmt.Sprintf("host = %s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPass)

	db, err := gorm.Open("postgres", urlConn)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	if err != nil {
		fmt.Println("Error on open Connection with Database")
		panic(err)
	}

	return db
}

/* USER */
func selectAll(table interface{}, c *gin.Context) interface{} {
	db := DBConnect()
	if err := db.Find(reflect.ValueOf(table).Interface()).Error; err != nil {
		checkErr(err, 500, c)
	}

	return table
}

func selectUserID(ID int, c *gin.Context) (user, *gorm.DB) {
	db := DBConnect()
	var userFind user
	if err := db.Where("id = ?", ID).First(&userFind).Error; err != nil {
		checkErr(err, 500, c)
	}

	return userFind, db
}

func selectUserByCpf(cpf string, c *gin.Context) (user, *gorm.DB) {
	db := DBConnect()
	var userFind user
	if err := db.Where("cpf = ?", cpf).First(&userFind).Error; err != nil {
		checkErr(err, 500, c)
	}

	return userFind, db
}

func selectUserByEmail(email string, c *gin.Context) (user, *gorm.DB) {
	db := DBConnect()
	var userFind user
	if err := db.Where("email = ?", email).First(&userFind).Error; err != nil {
		checkErr(err, 500, c)
	}

	return userFind, db
}

/* INCOME */

func selectIncomeByID(ID int, c *gin.Context) (income, *gorm.DB) {
	db := DBConnect()
	var incomeFind income
	if err := db.Where("id = ?", ID).First(&incomeFind).Error; err != nil {
		checkErr(err, 500, c)
	}

	return incomeFind, db
}

/* EXPENSE */
func selectExpenseByID(ID int, c *gin.Context) (expense, *gorm.DB) {
	db := DBConnect()
	var expenseFind expense
	if err := db.Where("id = ?", ID).First(&expenseFind).Error; err != nil {
		checkErr(err, 500, c)
	}

	return expenseFind, db
}

func checkErr(err error, statusCode int, c *gin.Context) {
	if err != nil {
		c.AbortWithError(statusCode, err)
		return
	}
}
