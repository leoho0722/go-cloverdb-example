package repository

import (
	documentv2 "github.com/ostafen/clover/v2/document"
	queryv2 "github.com/ostafen/clover/v2/query"
	"leoho.io/go-cloverdb-example/database"
)

// NewDocument 用來建立新的 Document
func NewDocument(collectionName string, document interface{}) (docId string, err error) {
	database.DB.Mutex.Lock()
	defer database.DB.Mutex.Unlock()

	doc := documentv2.NewDocumentOf(document)
	return database.DB.Context.InsertOne(collectionName, doc)
}

// GetDocuments 使用指定 Query 來查詢 Documents
func GetDocuments(query *queryv2.Query) (docs []*documentv2.Document, err error) {
	database.DB.Mutex.Lock()
	defer database.DB.Mutex.Unlock()

	return database.DB.Context.FindAll(query)
}

// UpdateDocument 使用指定 Query 來更新 Document
func UpdateDocument(query *queryv2.Query, updates map[string]interface{}) error {
	database.DB.Mutex.Lock()
	defer database.DB.Mutex.Unlock()

	return database.DB.Context.Update(query, updates)
}

// DeleteDocument 使用指定 Query 來刪除 Document
func DeleteDocument(query *queryv2.Query) error {
	database.DB.Mutex.Lock()
	defer database.DB.Mutex.Unlock()

	return database.DB.Context.Delete(query)
}
