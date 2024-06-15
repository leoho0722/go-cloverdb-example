package repository

import (
	"leoho.io/go-cloverdb-example/database"
)

// NewCollection 用來建立新的 Collection
func NewCollection(collectionName string) error {
	database.DB.Mutex.Lock()
	defer database.DB.Mutex.Unlock()
	exists, err := database.DB.Context.HasCollection(collectionName)
	if err != nil {
		return err
	}
	if !exists {
		err = database.DB.Context.CreateCollection(collectionName)
		if err != nil {
			return err
		}
	}
	return nil
}

// ListCollections 用來列出所有 Collection
func ListCollections() ([]string, error) {
	database.DB.Mutex.Lock()
	defer database.DB.Mutex.Unlock()
	return database.DB.Context.ListCollections()
}

// ImportCollection 用來匯入 Collection
func ImportCollection(collectionName string, importPath string) error {
	database.DB.Mutex.Lock()
	defer database.DB.Mutex.Unlock()

	return database.DB.Context.ImportCollection(collectionName, importPath)
}

// ExportCollection 用來匯出 Collection
func ExportCollection(collectionName string, exportPath string) error {
	database.DB.Mutex.Lock()
	defer database.DB.Mutex.Unlock()
	return database.DB.Context.ExportCollection(collectionName, exportPath)
}

// DropCollection 用來刪除 Collection
func DropCollection(collectionName string) error {
	database.DB.Mutex.Lock()
	defer database.DB.Mutex.Unlock()
	return database.DB.Context.DropCollection(collectionName)
}
