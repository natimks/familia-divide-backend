package src

//AutoMigration is the func for update the database schema
func AutoMigration() {
	db := DBConnect()
	defer db.Close()

	db.AutoMigrate(user{}, income{}, expense{})

	db.Model(&income{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&expense{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
