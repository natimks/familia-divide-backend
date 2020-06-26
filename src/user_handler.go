package src

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetUsers find all users
func GetUsers(c *gin.Context) {
	var users []user
	allUsers := selectAll(&users, c)

	c.JSON(200, allUsers)
}

//GetUser find a user by ID
func GetUser(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		checkErr(err, 400, c)
		return
	}
	selectedUser, _ := selectUserID(ID, c)
	c.JSON(200, selectedUser)
}

//GetUser find a user by CPF
func GetUserByCpf(c *gin.Context) {
	cpf := c.Param("cpf")

	selectedUser, db := selectUserByCpf(cpf, c)
	defer db.Close()

	db.Preload("UsersFamily").First(&selectedUser)
	db.Preload("Incomes").First(&selectedUser)
	db.Preload("Expenses").First(&selectedUser)

	c.JSON(200, selectedUser)
}

//PostUser create user
func PostUser(c *gin.Context) {
	var newUser user
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := DBConnect()
	defer db.Close()
	db.Create(&newUser)

	c.JSON(201, newUser)
}

//PutUser update user by ID
func PutUser(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		checkErr(err, 400, c)
		return
	}

	var updateUser user
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, db := selectUserID(ID, c)
	defer db.Close()

	user.CPF = updateUser.CPF
	user.Name = updateUser.Name
	user.Email = updateUser.Email
	user.Phone = updateUser.Phone

	db.Save(&user)

	c.JSON(200, user)
}

//DeleteUser delete a user by ID
func DeleteUser(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		checkErr(err, 400, c)
		return
	}

	user, db := selectUserID(ID, c)
	defer db.Close()

	if user.ID > 0 {
		db.Delete(&user)
	}

	c.JSON(204, nil)
}
