package user

import (
	"fmt"
	"log"
	"zarbat_data/database"
)

func DbGetUsers() (users []User) {
	db := database.Db
	row, err := db.Query("SELECT * FROM user ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var user User
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&user.Id, &user.Name, &user.Description)
		users = append(users, user)
	}
	return users
}
func DbGetUser(id int) (user User) {
	db := database.Db
	stmt, err := db.Prepare("SELECT id, name, description FROM user WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(id)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&user.Id, &user.Name, &user.Description)
	return user
}
func DbAddUser(user User) User {
	db := database.Db
	insert := "INSERT INTO user (name, description) values ( ?, ?) RETURNING id"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := 0
	err = stmt.QueryRow(user.Name, user.Description).Scan(&id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	user.Id = id
	return user
}
func DbUpdateUser(user User) User {
	db := database.Db
	update := "UPDATE user SET name = ?, description = ? WHERE id = ?"
	stmt, err := db.Prepare(update)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := stmt.Exec(user.Name, user.Description, user.Id)
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(affect)
	return DbGetUser(user.Id)
}
func DbDeleteUser(id int) User {
	db := database.Db
	delete := "DELETE FROM user WHERE id = ?"
	stmt, err := db.Prepare(delete)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := stmt.Exec(id)
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(affect)
	return DbGetUser(id)
}
