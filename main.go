package main

import (
	"fmt"

	queryv2 "github.com/ostafen/clover/v2/query"

	"leoho.io/go-cloverdb-example/database"
	"leoho.io/go-cloverdb-example/services"
)

func main() {
	if err := database.ConnectDB(); err != nil {
		panic(err)
	}
	defer database.DB.Context.Close()

	// Import
	// importCollection()

	// Import All
	// importCollections()

	// Create Collection
	newCollection()

	// Create Collections
	// newCollections()

	// List Collections
	listCollections()

	// Create
	// newDocument()

	// Create All
	newDocuments()

	// Read
	// readDocument()

	// Read All
	readDocuments()

	// Update
	// updateDocument()

	// Delete
	// deleteDocument()

	// Export
	exportCollection()

	// Export All
	// exportCollections()

	// Drop
	// dropCollection()
}

/* Collection Operations */

func newCollection() {
	err := services.NewCollection("users")
	if err != nil {
		panic(err)
	}
}
func newCollections() {
	err := services.NewCollections("posts", "comments")
	if err != nil {
		panic(err)
	}
}
func listCollections() {
	collections, err := services.ListCollections()
	if err != nil {
		panic(err)
	}
	fmt.Println(collections)
}
func importCollection() {
	err := services.ImportCollection("users", "./users.json")
	if err != nil {
		panic(err)
	}
}
func importCollections() {
	err := services.ImportCollections("./", "users", "posts", "comments")
	if err != nil {
		panic(err)
	}
}
func exportCollection() {
	err := services.ExportCollection("users", "./users.json")
	if err != nil {
		panic(err)
	}
}
func exportCollections() {
	err := services.ExportCollections("./")
	if err != nil {
		panic(err)
	}
}
func dropCollection() {
	err := services.DropCollection("users")
	if err != nil {
		panic(err)
	}
}

/* Document Operations */

func newDocument() {
	document := map[string]interface{}{
		"name": "David",
	}
	docId, err := services.NewDocument("users", document)
	if err != nil {
		panic(err)
	}
	fmt.Println(docId)
}
func newDocuments() {
	documents := []map[string]interface{}{
		{"name": "David"},
		{"nickname": "Beckham"},
	}
	var docInterfaces []interface{}
	for _, doc := range documents {
		docInterfaces = append(docInterfaces, doc)
	}
	docIds, err := services.NewDocuments("users", docInterfaces)
	if err != nil {
		panic(err)
	}
	fmt.Println(docIds)
}
func readDocument() {
	readQuery := services.NewBasicQuery("users")
	readQuery = readQuery.Where(queryv2.Field("_id").Eq("501f9b9d-7c21-4a23-9018-0cce404639fa"))
	doc, err := services.GetDocument(readQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println(doc)
}
func readDocuments() {
	readAllQuery := services.NewBasicQuery("users")
	docs, err := services.GetDocuments(readAllQuery)
	if err != nil {
		panic(err)
	}
	for _, doc := range docs {
		fmt.Println(doc)
	}
}
func updateDocument() {
	updateQuery := services.NewBasicQuery("users")
	updateQuery = updateQuery.Where(queryv2.Field("_id").Eq("66a3052b-7233-4ebf-acab-375cb457cbdd"))
	updates := map[string]interface{}{
		"name": "David Beckham",
	}
	err := services.UpdateDocument(updateQuery, updates)
	if err != nil {
		panic(err)
	}
}
func deleteDocument() {
	deleteQuery := services.NewBasicQuery("users")
	deleteQuery = deleteQuery.Where(queryv2.Field("_id").Eq("fc91ec20-f074-4714-980b-7a601fc72b20"))
	err := services.DeleteDocument(deleteQuery)
	if err != nil {
		panic(err)
	}
}
