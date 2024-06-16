package services

import (
	documentv2 "github.com/ostafen/clover/v2/document"
	queryv2 "github.com/ostafen/clover/v2/query"
	"leoho.io/go-cloverdb-example/repository"
)

// NewDocument 用來建立新的 Document
func NewDocument(collectionName string, document interface{}) (docId string, err error) {
	return repository.NewDocument(collectionName, document)
}

// NewDocuments 用來建立新的 Documents
func NewDocuments(collectionName string, documents []interface{}) (docIds []string, err error) {
	for _, document := range documents {
		docId, err := NewDocument(collectionName, document)
		if err != nil {
			return nil, err
		}
		docIds = append(docIds, docId)
	}
	return docIds, nil
}

// GetDocument 使用指定 Query 來查詢 Document
func GetDocument(query *queryv2.Query) (doc *documentv2.Document, err error) {
	return repository.GetDocument(query)
}

// GetDocuments 使用指定 Query 來查詢 Documents
func GetDocuments(query *queryv2.Query) (docs []*documentv2.Document, err error) {
	return repository.GetDocuments(query)
}

// UpdateDocument 使用指定 Query 來更新 Document
func UpdateDocument(query *queryv2.Query, updates map[string]interface{}) error {
	return repository.UpdateDocument(query, updates)
}

// DeleteDocument 使用指定 Query 來刪除 Document
func DeleteDocument(query *queryv2.Query) error {
	return repository.DeleteDocument(query)
}
