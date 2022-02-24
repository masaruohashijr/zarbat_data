package context

import (
	"fmt"
	"log"
	"zarbat_data/database"
	p "zarbat_data/features/number"
)

func DbGetPhoneNumbersContext(context Context) Context {
	db := database.Db
	var phoneNumber p.PhoneNumber
	var phoneNumbers []p.PhoneNumber
	stmt, _ := db.Prepare("SELECT b.id, b.phoneNumber, b.sid, b.description, a.alias, COALESCE(a.position,'') " +
		"FROM phoneNumberContext a INNER JOIN \"number\" b ON a.phoneNumberId = b.id WHERE a.contextId = ? ORDER BY position")
	rows, err := stmt.Query(context.Id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&phoneNumber.Id, &phoneNumber.PhoneNumber, &phoneNumber.Sid, &phoneNumber.Description, &phoneNumber.Alias, &phoneNumber.Position)
		if err != nil {
			log.Fatalln(err.Error())
		}
		phoneNumbers = append(phoneNumbers, phoneNumber)
	}
	context.PhoneNumbers = phoneNumbers
	return context
}

func DbAddPhoneNumbersContext(context Context) {
	db := database.Db
	insert := "INSERT INTO phoneNumberContext (contextId, phoneNumberId, alias, position) values ( ?, ?, ?, ?) RETURNING id"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := 0
	for index, phoneNumber := range context.PhoneNumbers {
		err := stmt.QueryRow(context.Id, phoneNumber.Id, phoneNumber.Alias, index).Scan(&id)
		if err != nil {
			log.Fatalln(err.Error())
		}
		println(insert, id, context.Id, phoneNumber.Id, phoneNumber.Alias, phoneNumber.Position)
	}
}

func DbUpdatePhoneNumbersContext(context Context) {
	db := database.Db
	stmt, err := db.Prepare(" SELECT " +
		" a.phoneNumberId, a.contextId, c.phoneNumber, c.sid, c.description, a.alias, a.position " +
		" FROM phoneNumberContext a  " +
		" INNER JOIN context b ON a.contextId = b.id   " +
		" INNER JOIN number c ON a.phoneNumberId = c.id  " +
		" WHERE a.contextId = ? " +
		" ORDER BY position")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(context.Id)
	var phoneNumber p.PhoneNumber
	var phoneNumbersDB []p.PhoneNumber
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&phoneNumber.Id, &phoneNumber.ContextId, &phoneNumber.PhoneNumber, &phoneNumber.Sid, &phoneNumber.Description, &phoneNumber.Alias, &phoneNumber.Position)
		phoneNumbersDB = append(phoneNumbersDB, phoneNumber)
	}
	copyContextPhoneNumbers := make([]p.PhoneNumber, len(context.PhoneNumbers))
	copy(copyContextPhoneNumbers, context.PhoneNumbers)
	phoneNumbersDB, phoneNumbersPage := diffPhoneNumbers(phoneNumbersDB, copyContextPhoneNumbers)
	if len(phoneNumbersDB) > 0 {
		deletePhoneNumbers(context, phoneNumbersDB)
	}
	if len(phoneNumbersPage) > 0 {
		addPhoneNumbers(context, phoneNumbersPage)
	}
	update := "UPDATE phoneNumberContext SET phoneNumberId = ?, alias = ? WHERE contextId = ? AND position = ?"
	stmt, err = db.Prepare(update)
	if err != nil {
		log.Fatalln(err.Error())
	}
	for _, phoneNumber := range context.PhoneNumbers {
		res, err := stmt.Exec(phoneNumber.Id, phoneNumber.Alias, context.Id, phoneNumber.Position)
		affect, err := res.RowsAffected()
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Println("PhoneNumber ", phoneNumber.Id, affect)
	}
}

func diffPhoneNumbers(phoneNumbersDB, phoneNumbersPage []p.PhoneNumber) ([]p.PhoneNumber, []p.PhoneNumber) {
	for indexDB := 0; indexDB < len(phoneNumbersDB); indexDB++ {
		sdb := phoneNumbersDB[indexDB]
		for indexPage := 0; indexPage < len(phoneNumbersPage); indexPage++ {
			sp := phoneNumbersPage[indexPage]
			if sdb.Id == sp.Id && sp.Position == sdb.Position {
				phoneNumbersDB = append(phoneNumbersDB[0:indexDB], phoneNumbersDB[indexDB+1:]...)
				phoneNumbersPage = append(phoneNumbersPage[0:indexPage], phoneNumbersPage[indexPage+1:]...)
				indexDB = -1
				break
			}
		}
	}
	return phoneNumbersDB, phoneNumbersPage
}

func deletePhoneNumbers(context Context, phoneNumbersToBeDeleted []p.PhoneNumber) {
	db := database.Db
	for _, phoneNumber := range phoneNumbersToBeDeleted {
		delete := "DELETE FROM phoneNumberContext WHERE contextId = ? AND phoneNumberId = ? AND position = ?"
		stmt, err := db.Prepare(delete)
		if err != nil {
			log.Fatalln(err.Error())
		}
		_, err = stmt.Exec(context.Id, phoneNumber.Id, phoneNumber.Position)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
}

func addPhoneNumbers(context Context, phoneNumbersToBeAdded []p.PhoneNumber) {
	db := database.Db
	for _, phoneNumber := range phoneNumbersToBeAdded {
		insert := "INSERT INTO phoneNumberContext (contextId, phoneNumberId, alias, position) values ( ?, ?, ?, ?)"
		stmt, err := db.Prepare(insert)
		if err != nil {
			log.Fatalln(err.Error())
		}
		_, err = stmt.Exec(context.Id, phoneNumber.Id, phoneNumber.Alias, phoneNumber.Position)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

}
