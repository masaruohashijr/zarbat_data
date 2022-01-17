package feature

import (
	"fmt"
	"log"
	"zarbat_mock/database"
)

func DbGetFeatures() (features []Feature) {
	db := database.Db
	row, err := db.Query("SELECT * FROM feature ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var feature Feature
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&feature.Id, &feature.Name, &feature.Description)
		features = append(features, feature)
	}
	return features
}
func DbGetFeature(id int) (feature Feature) {
	db := database.Db
	stmt, err := db.Prepare("SELECT id, name, description FROM feature WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(id)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&feature.Id, &feature.Name, &feature.Description)
	return feature
}
func DbAddFeature(feature Feature) Feature {
	db := database.Db
	insert := "INSERT INTO feature (name, description) values ( ?, ?) RETURNING id"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := 0
	err = stmt.QueryRow(feature.Name, feature.Description).Scan(&id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	feature.Id = id
	return feature
}
func DbUpdateFeature(feature Feature) Feature {
	db := database.Db
	update := "UPDATE feature SET name = ?, description = ? WHERE id = ?"
	stmt, err := db.Prepare(update)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := stmt.Exec(feature.Name, feature.Description, feature.Id)
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(affect)
	return DbGetFeature(feature.Id)
}
func DbDeleteFeature(id int) Feature {
	db := database.Db
	delete := "DELETE FROM feature WHERE id = ?"
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
	return DbGetFeature(id)
}
