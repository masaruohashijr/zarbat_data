package context

import (
	"fmt"
	"log"
	"zarbat_mock/database"
)

func DbGetContexts() (contexts []Context) {
	db := database.Db
	row, err := db.Query("SELECT * FROM context ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var context Context
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&context.Id, &context.Name, &context.Description, &context.EnvironmentId)
		contexts = append(contexts, context)
	}
	return contexts
}
func DbGetContextsByEnv(environmentId int) (contexts []Context) {
	db := database.Db
	stmt, err := db.Prepare("SELECT id, name, description, environmentId FROM context WHERE environmentId = ?")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(environmentId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var context Context
	for rows.Next() { // Iterate and fetch the records from result cursor
		rows.Scan(&context.Id, &context.Name, &context.Description, &context.EnvironmentId)
		contexts = append(contexts, context)
	}
	return contexts
}
func DbGetContext(id int) (context Context) {
	db := database.Db
	stmt, err := db.Prepare("SELECT id, name, description, environmentId FROM context WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(id)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&context.Id, &context.Name, &context.Description, &context.EnvironmentId)
	return context
}
func DbAddContext(context Context) Context {
	db := database.Db
	insert := "INSERT INTO context (name, description, environmentId) values ( ?, ?, ?) RETURNING id"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := 0
	err = stmt.QueryRow(context.Name, context.Description, context.EnvironmentId).Scan(&id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	context.Id = id
	return context
}
func DbUpdateContext(context Context) Context {
	db := database.Db
	update := "UPDATE context SET name = ?, description = ?, environmentId = ? WHERE id = ?"
	stmt, err := db.Prepare(update)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := stmt.Exec(context.Name, context.Description, context.EnvironmentId, context.Id)
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(affect)
	return DbGetContext(context.Id)
}
func DbDeleteContext(id int) Context {
	db := database.Db
	delete := "DELETE FROM context WHERE id = ?"
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
	return DbGetContext(id)
}
