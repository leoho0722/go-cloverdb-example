package services

import (
	"fmt"

	"leoho.io/go-cloverdb-example/repository"
)

// NewCollection 用來建立新的 Collection
func NewCollection(collectionName string) error {
	return repository.NewCollection(collectionName)
}

// NewCollections 用來建立多個 Collection
func NewCollections(collectionNames ...string) error {
	for _, collectionName := range collectionNames {
		if err := NewCollection(collectionName); err != nil {
			return err
		}
	}
	return nil
}

// ListCollections 用來列出所有 Collection
func ListCollections() ([]string, error) {
	return repository.ListCollections()
}

// ImportCollection 用來匯入 Collection
func ImportCollection(collectionName string, importPath string) error {
	return repository.ImportCollection(collectionName, importPath)
}

// ImportCollections 用來匯入多個 Collection
func ImportCollections(importPath string, collectionNames ...string) error {
	for _, collectionName := range collectionNames {
		path := fmt.Sprintf("%s/%s.json", importPath, collectionName)
		if err := ImportCollection(collectionName, path); err != nil {
			return err
		}
	}
	return nil
}

// ExportCollection 用來匯出 Collection
func ExportCollection(collectionName string, exportPath string) error {
	return repository.ExportCollection(collectionName, exportPath)
}

// ExportCollections 用來匯出多個 Collection
func ExportCollections(exportPath string, collectionNames ...string) error {
	var _collectionNames []string
	if len(collectionNames) > 0 {
		_collectionNames = collectionNames
	} else {
		collectionNames, err := ListCollections()
		if err != nil {
			return err
		}
		_collectionNames = collectionNames
	}
	for _, collectionName := range _collectionNames {
		path := fmt.Sprintf("%s/%s.json", exportPath, collectionName)
		if err := ExportCollection(collectionName, path); err != nil {
			return err
		}
	}
	return nil
}

// DropCollection 用來刪除 Collection
func DropCollection(collectionName string) error {
	return repository.DropCollection(collectionName)
}

// DropCollections 用來刪除多個 Collection
func DropCollections(collectionNames ...string) error {
	for _, collectionName := range collectionNames {
		if err := DropCollection(collectionName); err != nil {
			return err
		}
	}
	return nil
}
