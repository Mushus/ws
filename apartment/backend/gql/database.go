package gql

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
)

// Resolver resolve on database
type Resolver struct {
	DB *sqlx.DB
}

// Articles Query.Articles
func (r *Resolver) Articles(p graphql.ResolveParams) (interface{}, error) {
	sql, args, err := sq.Select("id", "name").From("articles").ToSql()
	if err != nil {
		return nil, err
	}

	articles := []Article{}
	r.DB.Select(&articles, sql, args...)
	return articles, nil
}

// Article Query.Article
func (r *Resolver) Article(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"]

	sql, args, err := sq.Select("id", "name").From("articles").Where("id = ?", id).Limit(1).ToSql()
	if err != nil {
		return nil, err
	}

	articles := []Article{}
	r.DB.Select(&articles, sql, args...)
	if len(articles) > 0 {
		return articles[0], nil
	}
	return nil, nil
}

// ArticleRooms ...Article.Rooms
func (r *Resolver) ArticleRooms(p graphql.ResolveParams) (interface{}, error) {
	articleID := p.Source.(Article).ID

	now, err := getNow(p.Args)
	if err != nil {
		return nil, err
	}

	qb := sq.
		Select(
			"r.`index` AS `index`",
			"r.id AS id",
			"r.name AS name",
			"r.rent AS rent",
			"r.article_id AS article_id",
		).
		From("rooms AS r").
		Where(sq.Eq{
			"r.article_id": articleID,
		}).
		OrderBy("r.`index` ASC")

	if status, ok := p.Args["status"]; ok {
		qb = qb.Join("tenants AS t ON t.room_id = r.id").GroupBy("r.id")
		switch status {
		case StatusLiving:
			qb = qb.Where("? BETWEEN t.since AND t.until", now)
		case StatusEmpty:
			qb = qb.Where("NOT (? BETWEEN t.since AND t.until)", now)
		case StatusReserved:
			qb = qb.Where("t.since > ?", now)
		}
	}

	sql, arg, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	rooms := []Room{}
	if err := r.DB.Select(&rooms, sql, arg...); err != nil {
		return nil, err
	}
	return rooms, nil
}

// Room Query.Room
func (r *Resolver) Room(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"]
	sql, args, err := sq.
		Select("id", "article_id", "`index`", "name", "rent").
		From("rooms").
		Where("id = ?", id).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, err
	}

	rooms := []Room{}
	r.DB.Select(&rooms, sql, args...)
	if len(rooms) > 0 {
		return rooms[0], nil
	}
	return nil, nil
}

// RoomTenant ...Room.Tenant
func (r *Resolver) RoomTenant(p graphql.ResolveParams) (interface{}, error) {
	roomID := p.Source.(Room).ID

	now, err := getNow(p.Args)
	if err != nil {
		return nil, err
	}

	sql, arg, err := sq.
		Select(
			"id",
			"name",
			"rent",
			"CAST(since AS TEXT) AS since",
			"CAST(until AS TEXT) AS until",
		).
		From("tenants").
		Where(
			sq.Expr("room_id = ?", roomID),
			sq.Expr("? BETWEEN since AND until", now),
		).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, err
	}
	tenant := new(Tenant)
	if err := r.DB.Get(tenant, sql, arg...); err != nil {
		return nil, err
	}
	return tenant, nil
}

// Tenant Query.Tenant
func (r *Resolver) Tenant(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"]

	sql, args, err := sq.
		Select(
			"id",
			"name",
			"rent",
			"CAST(since AS TEXT) AS since",
			"CAST(until AS TEXT) AS until",
		).
		From("tenants").
		Where("id = ?", id).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, err
	}

	tenant := new(Tenant)
	r.DB.Get(&tenant, sql, args...)
	return tenant, nil
}

// TenantBills ...Tenant.Bills
func (r *Resolver) TenantBills(p graphql.ResolveParams) (interface{}, error) {
	qb := sq.
		Select(
			"id",
			"tenant_id",
			"billing_term_id",
			"rent",
			"CAST(since AS TEXT) AS since",
			"CAST(until AS TEXT) AS until",
		).
		From("bills")
	if since, ok := p.Args["since"]; ok {
		if _, err := time.Parse(DateFormat, since.(string)); err != nil {
			return nil, err
		}
		qb = qb.Where("? <= since", since)
	}

	if until, ok := p.Args["until"]; ok {
		if _, err := time.Parse(DateFormat, until.(string)); err != nil {
			return nil, err
		}
		qb = qb.Where("until <= ?")
	}

	sql, arg, err := qb.
		ToSql()
	if err != nil {
		return nil, err
	}

	bills := []Bill{}
	if err := r.DB.Select(&bills, sql, arg...); err != nil {
		return nil, err
	}

	return bills, nil
}

// RoomTenantHistory ...Room.TenantHistory
func (r *Resolver) RoomTenantHistory(p graphql.ResolveParams) (interface{}, error) {
	qb := sq.
		Select(
			"id",
			"name",
			"rent",
			"CAST(since AS TEXT) AS since",
			"CAST(until AS TEXT) AS until",
		).
		From("tenants")

	if since, ok := p.Args["since"]; ok {
		if _, err := time.Parse(DateFormat, since.(string)); err != nil {
			return nil, err
		}
		qb = qb.Where("? <= since", since)
	}

	if until, ok := p.Args["until"]; ok {
		if _, err := time.Parse(DateFormat, until.(string)); err != nil {
			return nil, err
		}
		qb = qb.Where("until <= ?", until)
	}

	sql, arg, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	tenants := []Tenant{}
	if err := r.DB.Select(&tenants, sql, arg...); err != nil {
		return nil, err
	}
	return tenants, nil
}
