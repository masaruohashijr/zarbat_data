package context

import (
	"fmt"
	"log"
	"zarbat_data/database"
	p "zarbat_data/features/parameter"
)

func DbGetParametersContext(context Context) Context {
	db := database.Db
	var parameter p.Parameter
	var parameters []p.Parameter
	stmt, _ := db.Prepare("SELECT b.id, b.name, b.description, a.value, COALESCE(a.position,'') " +
		"FROM parameterContext a INNER JOIN parameter b ON a.parameterId = b.id WHERE a.contextId = ? ORDER BY position")
	rows, err := stmt.Query(context.Id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&parameter.Id, &parameter.Name, &parameter.Description, &parameter.Value, &parameter.Position)
		if err != nil {
			log.Fatalln(err.Error())
		}
		parameters = append(parameters, parameter)
	}
	context.Parameters = parameters
	return context
}

func DbAddParametersContext(context Context) {
	db := database.Db
	insert := "INSERT INTO parameterContext (contextId, parameterId, value, position) values ( ?, ?, ?, ?) RETURNING id"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := 0
	for index, parameter := range context.Parameters {
		err := stmt.QueryRow(context.Id, parameter.Id, parameter.Value, index).Scan(&id)
		if err != nil {
			log.Fatalln(err.Error())
		}
		println(insert, id, context.Id, parameter.Id, parameter.Value, parameter.Position)
	}
}

func DbUpdateParametersContext(context Context) {
	db := database.Db
	stmt, err := db.Prepare(" SELECT " +
		" a.parameterId, a.contextId, c.name, c.description, a.value, a.position " +
		" FROM parameterContext a  " +
		" INNER JOIN context b ON a.contextId = b.id   " +
		" INNER JOIN parameter c ON a.parameterId = c.id  " +
		" WHERE a.contextId = ? " +
		" ORDER BY position")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(context.Id)
	var parameter p.Parameter
	var parametersDB []p.Parameter
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&parameter.Id, &parameter.ContextId, &parameter.Name, &parameter.Description, &parameter.Value, &parameter.Position)
		parametersDB = append(parametersDB, parameter)
	}
	copyContextParameters := make([]p.Parameter, len(context.Parameters))
	copy(copyContextParameters, context.Parameters)
	parametersDB, parametersPage := diffParameters(parametersDB, copyContextParameters)
	if len(parametersDB) > 0 {
		deleteParameters(context, parametersDB)
	}
	if len(parametersPage) > 0 {
		addParameters(context, parametersPage)
	}
	update := "UPDATE parameterContext SET parameterId = ?, value = ? WHERE contextId = ? AND position = ?"
	stmt, err = db.Prepare(update)
	if err != nil {
		log.Fatalln(err.Error())
	}
	for _, parameter := range context.Parameters {
		res, err := stmt.Exec(parameter.Id, parameter.Value, context.Id, parameter.Position)
		affect, err := res.RowsAffected()
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Println("Parameter ", parameter.Id, affect)
	}
}

func diffParameters(parametersDB, parametersPage []p.Parameter) ([]p.Parameter, []p.Parameter) {
	for indexDB := 0; indexDB < len(parametersDB); indexDB++ {
		sdb := parametersDB[indexDB]
		for indexPage := 0; indexPage < len(parametersPage); indexPage++ {
			sp := parametersPage[indexPage]
			if sdb.Id == sp.Id && sp.Position == sdb.Position {
				parametersDB = append(parametersDB[0:indexDB], parametersDB[indexDB+1:]...)
				parametersPage = append(parametersPage[0:indexPage], parametersPage[indexPage+1:]...)
				indexDB = -1
				break
			}
		}
	}
	return parametersDB, parametersPage
}

func deleteParameters(context Context, parametersToBeDeleted []p.Parameter) {
	db := database.Db
	for _, parameter := range parametersToBeDeleted {
		delete := "DELETE FROM parameterContext WHERE contextId = ? AND parameterId = ? AND position = ?"
		stmt, err := db.Prepare(delete)
		if err != nil {
			log.Fatalln(err.Error())
		}
		_, err = stmt.Exec(context.Id, parameter.Id, parameter.Position)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
}

func addParameters(context Context, parametersToBeAdded []p.Parameter) {
	db := database.Db
	for _, parameter := range parametersToBeAdded {
		insert := "INSERT INTO parameterContext (contextId, parameterId, value, position) values ( ?, ?, ?, ?)"
		stmt, err := db.Prepare(insert)
		if err != nil {
			log.Fatalln(err.Error())
		}
		_, err = stmt.Exec(context.Id, parameter.Id, parameter.Value, parameter.Position)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

}
