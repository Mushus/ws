package gqli

import (
	"github.com/graphql-go/graphql"
)

var articlesField = &graphql.Field{
	Type:        graphql.NewList(articleType),
	Description: "Get articles",
	Resolve:     func(p graphql.ResolveParams) (interface{}, error) { return resolver.Articles(p) },
}

var articleField = &graphql.Field{
	Type:        articleType,
	Description: "Get single article",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) { return resolver.Article(p) },
}

var roomField = &graphql.Field{
	Type:        roomType,
	Description: "Get single room",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) { return resolver.Room(p) },
}

var tenantField = &graphql.Field{
	Type:        tenantType,
	Description: "Get single tenant",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) { return resolver.Tenant(p) },
}

var createArticleField = &graphql.Field{
	Type:        articleType,
	Description: "Create new article",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return nil, nil
	},
}

var loginField = &graphql.Field{
	Type:        loginType,
	Description: "login administrator",
	Args: graphql.FieldConfigArgument{
		"login": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
}
