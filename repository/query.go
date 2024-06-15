package repository

import (
	queryv2 "github.com/ostafen/clover/v2/query"
)

// NewBasicQuery 用來建立新的 Query
func NewBasicQuery(collectionName string) *queryv2.Query {
	return queryv2.NewQuery(collectionName)
}
