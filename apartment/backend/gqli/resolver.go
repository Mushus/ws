package gqli

import "github.com/graphql-go/graphql"

var resolver Resolver

// SetResolver resolverを設定
func SetResolver(r Resolver) {
	resolver = r
}

// Resolver GraphQLの解決を行う人
type Resolver interface {
	Articles(p graphql.ResolveParams) (interface{}, error)

	Article(p graphql.ResolveParams) (interface{}, error)
	ArticleRooms(p graphql.ResolveParams) (interface{}, error)

	Room(p graphql.ResolveParams) (interface{}, error)
	RoomTenant(p graphql.ResolveParams) (interface{}, error)
	RoomTenantHistory(p graphql.ResolveParams) (interface{}, error)

	Tenant(p graphql.ResolveParams) (interface{}, error)
	TenantBills(p graphql.ResolveParams) (interface{}, error)
}
