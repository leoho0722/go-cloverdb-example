package database

import (
	cloverv2 "github.com/ostafen/clover/v2"
	documentv2 "github.com/ostafen/clover/v2/document"
	queryv2 "github.com/ostafen/clover/v2/query"
)

// ConnectDB 用來連結 CloverDB
func ConnectDB() (db *cloverv2.DB, err error) {
	const (
		dbDir = "./"
	)
	db, err = cloverv2.Open(dbDir)
	return db, err
}

// NewCollection 用來建立新的 Collection
func NewCollection(db *cloverv2.DB, collectionName string) error {
	exists, err := db.HasCollection(collectionName)
	if err != nil {
		return err
	}
	if !exists {
		err = db.CreateCollection(collectionName)
		if err != nil {
			return err
		}
	}
	return nil
}

// ExportCollection 用來匯出 Collection
func ExportCollection(db *cloverv2.DB, collectionName string, exportPath string) error {
	return db.ExportCollection(collectionName, exportPath)
}

// NewDocument 用來建立新的 Document
func NewDocument(
	db *cloverv2.DB,
	collectionName string,
	document map[string]interface{},
) (docId string, err error) {
	doc := documentv2.NewDocumentOf(document)
	docId, err = db.InsertOne(collectionName, doc)
	return
}

// NewBasicQuery 用來建立新的 Query
func NewBasicQuery(collectionName string) *queryv2.Query {
	return queryv2.NewQuery(collectionName)
}

// UpdateDocument 使用指定 Query 來更新 Document
func UpdateDocument(
	db *cloverv2.DB,
	query *queryv2.Query,
	updates map[string]interface{},
) error {
	return db.Update(query, updates)
}

// DeleteDocument 使用指定 Query 來刪除 Document
func DeleteDocument(db *cloverv2.DB, query *queryv2.Query) error {
	return db.Delete(query)
}
