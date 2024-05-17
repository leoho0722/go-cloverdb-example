package main

import (
	"fmt"

	queryv2 "github.com/ostafen/clover/v2/query"

	"leoho.io/go-cloverdb-example/database"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = database.NewCollection(db, "users")
	if err != nil {
		panic(err)
	}
	collections, err := db.ListCollections()
	if err != nil {
		panic(err)
	}
	fmt.Println(collections)

	document := make(map[string]interface{})
	document["name"] = "Leo Ho"

	// CREATE
	docId, err := database.NewDocument(db, "users", document)
	if err != nil {
		panic(err)
	}
	fmt.Println(docId)

	// READ
	query := database.NewBasicQuery("users")
	docs, err := db.FindAll(query)
	if err != nil {
		panic(err)
	}
	for _, doc := range docs {
		fmt.Println(doc)
	}

	// UPDATE
	query1 := database.NewBasicQuery("users")
	query1 = query1.Where(queryv2.Field("_id").Eq("66a3052b-7233-4ebf-acab-375cb457cbdd"))
	updates := make(map[string]interface{})
	updates["name"] = "IU"
	err = database.UpdateDocument(db, query1, updates)
	if err != nil {
		panic(err)
	}

	// DELETE
	deleteQuery := database.NewBasicQuery("users")
	deleteQuery = deleteQuery.Where(queryv2.Field("_id").Eq("fc91ec20-f074-4714-980b-7a601fc72b20"))
	err = database.DeleteDocument(db, deleteQuery)
	if err != nil {
		panic(err)
	}

	// EXPORT
	err = database.ExportCollection(db, "users", "./users.json")
	if err != nil {
		panic(err)
	}
}
