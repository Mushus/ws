package gqli

import (
	"github.com/graphql-go/graphql"
)

var billType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Bill",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"rent": &graphql.Field{
			Type: graphql.Int,
		},
		"since": &graphql.Field{
			Type: graphql.String,
		},
		"until": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var tenantType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Tenant",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"rent": &graphql.Field{
			Type: graphql.Int,
		},
		"since": &graphql.Field{
			Type: graphql.String,
		},
		"until": &graphql.Field{
			Type: graphql.String,
		},
		"bills": &graphql.Field{
			Type:        graphql.NewList(billType),
			Description: "getBills",
			Args: graphql.FieldConfigArgument{
				"since": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"until": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) { return resolver.TenantBills(p) },
		},
	},
})

var roomType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Room",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"rent": &graphql.Field{
			Type: graphql.Int,
		},
		"tenant": &graphql.Field{
			Type:        tenantType,
			Description: "Get tenant moving-in",
			Args: graphql.FieldConfigArgument{
				"now": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) { return resolver.RoomTenant(p) },
		},
		"tenantHistory": &graphql.Field{
			Type:        graphql.NewList(tenantType),
			Description: "List of past tenants",
			Args: graphql.FieldConfigArgument{
				"since": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"until": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) { return resolver.RoomTenantHistory(p) },
		},
	},
})

var articleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Article",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"rooms": &graphql.Field{
			Type:        graphql.NewList(roomType),
			Description: "Get rooms",
			Args: graphql.FieldConfigArgument{
				"now": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"status": &graphql.ArgumentConfig{
					Type: graphql.NewEnum(graphql.EnumConfig{
						Name:        "status",
						Description: "room status",
						Values: graphql.EnumValueConfigMap{
							"Living": &graphql.EnumValueConfig{
								Description: "this room have people living in",
								Value:       0,
							},
							"Empty": &graphql.EnumValueConfig{
								Description: "this room is empty",
								Value:       1,
							},
							"Reserved": &graphql.EnumValueConfig{
								Description: "this room is scheduled to move in",
								Value:       2,
							},
						},
					}),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) { return resolver.ArticleRooms(p) },
		},
	},
})

var loginType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Login",
	Fields: graphql.Fields{
		"token": &graphql.Field{
			Type: graphql.String,
		},
	},
})
