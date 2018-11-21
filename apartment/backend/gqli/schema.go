package gqli

import (
	"github.com/graphql-go/graphql"
)

// Schema get Schema object
func Schema() graphql.Schema {
	return schema
}

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	},
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"articles": articlesField,
		"article":  articleField,
		"room":     roomField,
		"tenant":   tenantField,
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"login":         loginField,
		"createArticle": createArticleField,
	},
})
