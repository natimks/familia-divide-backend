package src

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

//PutUser update user add new relative
func PutAddRelative(c *gin.Context) {
	cpf_admin := c.Param("cpf")
	email_relative := c.Param("email")

	adminUser, db := selectUserByCpf(cpf_admin, c)
	relativeUser, _ := selectUserByEmail(email_relative, c)
	defer db.Close()

	db.Model(&adminUser).Association("UsersFamily").Append(&relativeUser)
	db.Save(&adminUser)

	c.JSON(200, relativeUser)
}

//DeleteRelative delete relative from user
func DeleteRelative(c *gin.Context) {
	id_admin, _ := strconv.Atoi(c.Param("id"))
	id_relative, _ := strconv.Atoi(c.Param("id_relative"))

	adminUser, db := selectUserID(id_admin, c)
	relativeUser, _ := selectUserID(id_relative, c)
	defer db.Close()

	db.Model(&adminUser).Association("UsersFamily").Delete(&relativeUser)
	db.Save(&adminUser)

	c.JSON(200, adminUser)
}
