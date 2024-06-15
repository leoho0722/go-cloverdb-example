package database

import (
	"sync"

	cloverv2 "github.com/ostafen/clover/v2"
	documentv2 "github.com/ostafen/clover/v2/document"
	queryv2 "github.com/ostafen/clover/v2/query"
)

type Database struct {
	Context *cloverv2.DB
	Mutex   sync.Mutex
}

var DB *Database

// ConnectDB 用來連結 CloverDB
func ConnectDB() error {
	const (
		dbDir = "./"
	)
	db, err := cloverv2.Open(dbDir)
	if err != nil {
		return err
	}
	DB = &Database{Context: db}
	return nil
}

type DatabaseOperation interface {
	// NewCollection 用來建立新的 Collection
	NewCollection(collectionName string) error
	// NewCollections 用來建立多個 Collection
	NewCollections(collectionNames ...string) error
	// ListCollections 用來列出所有 Collection
	ListCollections() ([]string, error)
	// ImportCollection 用來匯入 Collection
	ImportCollection(collectionName string, importPath string) error
	// ImportCollections 用來匯入多個 Collection
	ImportCollections(importPath string, collectionNames ...string) error
	// ExportCollection 用來匯出 Collection
	ExportCollection(collectionName string, exportPath string) error
	// ExportCollections 用來匯出多個 Collection
	ExportCollections(exportPath string, collectionNames ...string) error
	// DropCollection 用來刪除 Collection
	DropCollection(collectionName string) error
	// DropCollections 用來刪除多個 Collection
	DropCollections(collectionNames ...string) error

	// NewDocument 用來建立新的 Document
	NewDocument(collectionName string, document interface{}) (docId string, err error)
	// GetDocument 使用指定 Query 來查詢 Document
	GetDocument(query *queryv2.Query) (doc *documentv2.Document, err error)
	// GetDocuments 使用指定 Query 來查詢 Documents
	GetDocuments(query *queryv2.Query) (docs []*documentv2.Document, err error)
	// UpdateDocument 使用指定 Query 來更新 Document
	UpdateDocument(query *queryv2.Query, updates map[string]interface{}) error
	// DeleteDocument 使用指定 Query 來刪除 Document
	DeleteDocument(query *queryv2.Query) error
}
