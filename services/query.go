package services

import (
	queryv2 "github.com/ostafen/clover/v2/query"

	"leoho.io/go-cloverdb-example/repository"
)

// NewBasicQuery 用來建立新的 Query
func NewBasicQuery(collectionName string) *queryv2.Query {
	return repository.NewBasicQuery(collectionName)
}
