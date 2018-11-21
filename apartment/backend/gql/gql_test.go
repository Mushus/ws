package gql_test

import (
	"encoding/json"
	"testing"

	"github.com/Mushus/apartment/backend/gql"
	"github.com/Mushus/apartment/backend/gqli"
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const sql = `
INSERT INTO articles (name) VALUES ('hoge');
INSERT INTO articles (name) VALUES ('fuga');
INSERT INTO rooms (article_id, 'index', name, rent) VALUES (1, 1, '101', 10000);
INSERT INTO rooms (article_id, 'index', name, rent) VALUES (1, 2, '201', 10000);
INSERT INTO rooms (article_id, 'index', name, rent) VALUES (2, 1, '101', 10000);
`

func initDB() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	db.MustExec(gql.Schema)
	db.MustExec(sql)
	gqli.SetResolver(&gql.Resolver{DB: db})
	return db
}

type testTableColumn struct {
	query  string
	result string
}

var testTable = []testTableColumn{
	{
		`{ articles { id name } }`,
		`{"data":{"articles":[{"id":1,"name":"hoge"}]}}`,
	},
}

func TestGraphQL(t *testing.T) {
	initDB()
	for _, column := range testTable {
		params := graphql.Params{Schema: gqli.Schema(), RequestString: column.query}
		r := graphql.Do(params)
		if len(r.Errors) > 0 {
			t.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
		}
		bjson, _ := json.Marshal(r)
		json := string(bjson)
		if json != column.result {
			t.Fatalf("invalid result,\nhave %s\nwant %s\n", json, column.result)
		}
	}
}

func TestGetArticles(t *testing.T) {
	db := initDB()
	db.MustExec("INSERT INTO articles (name) VALUES ('hoge')")

	query := `{
		articles {
			id
			name
		}
	}`
	params := graphql.Params{Schema: gqli.Schema(), RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		t.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	bjson, _ := json.Marshal(r)
	json := string(bjson)
	want := `{"data":{"articles":[{"id":1,"name":"hoge"}]}}`
	if json != want {
		t.Fatalf("invalid result,\nhave %s\nwant %s\n", json, want)
	}
}

func TestGetConditionArticles(t *testing.T) {
	db := initDB()
	db.MustExec("INSERT INTO articles (name) VALUES ('hoge')")
	db.MustExec("INSERT INTO articles (name) VALUES ('fuga')")

	query := `{
		article(id: 2) {
			id
			name
		}
	}`
	params := graphql.Params{Schema: gqli.Schema(), RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		t.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	bjson, _ := json.Marshal(r)
	json := string(bjson)
	want := `{"data":{"article":{"id":2,"name":"fuga"}}}`
	if json != want {
		t.Fatalf("invalid result,\nhave %s\nwant %s\n", json, want)
	}
}

func TestGetArticleWithRoom(t *testing.T) {
	db := initDB()
	db.MustExec("INSERT INTO articles (name) VALUES ('hoge')")
	db.MustExec("INSERT INTO articles (name) VALUES ('fuga')")
	db.MustExec("INSERT INTO rooms (article_id, 'index', name, rent) VALUES (1, 1, '101', 10000)")
	db.MustExec("INSERT INTO rooms (article_id, 'index', name, rent) VALUES (1, 2, '201', 10000)")
	db.MustExec("INSERT INTO rooms (article_id, 'index', name, rent) VALUES (2, 1, '101', 10000)")

	query := `{
		article(id: 1) {
			id
			rooms {
				id
				name
			}
		}
	}`
	params := graphql.Params{Schema: gqli.Schema(), RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		t.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	bjson, _ := json.Marshal(r)
	json := string(bjson)
	want := `{"data":{"article":{"id":1,"rooms":[{"id":1,"name":"101"},{"id":2,"name":"201"}]}}}`
	if json != want {
		t.Fatalf("invalid result,\nhave %s\nwant %s\n", json, want)
	}
}

func TestGetArticleWithRooms(t *testing.T) {
	db := initDB()
	db.MustExec("INSERT INTO articles (name) VALUES ('hoge')")
	db.MustExec("INSERT INTO articles (name) VALUES ('fuga')")
	db.MustExec("INSERT INTO rooms (article_id, 'index', name, rent) VALUES (1, 1, '101', 10000)")
	db.MustExec("INSERT INTO rooms (article_id, 'index', name, rent) VALUES (1, 2, '201', 10000)")
	db.MustExec("INSERT INTO rooms (article_id, 'index', name, rent) VALUES (2, 1, '101', 10000)")
	db.MustExec("INSERT INTO tenants (room_id, name, rent, since, until) VALUES (1, 'Bob', 10000, '1990-01-01', '2010-01-01')")
	db.MustExec("INSERT INTO tenants (room_id, name, rent, since, until) VALUES (1, 'Jane', 10000, '2011-01-01', '2018-01-01')")

	query := `{
		article(id: 1){
			id
			rooms(now: "2000-01-01", status: Living) {
				id
				name
			}
		}
	}`
	params := graphql.Params{Schema: gqli.Schema(), RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		t.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	bjson, _ := json.Marshal(r)
	json := string(bjson)
	want := `{"data":{"article":{"id":1,"rooms":[{"id":1,"name":"101"}]}}}`
	if json != want {
		t.Fatalf("invalid result,\nhave %s\nwant %s\n", json, want)
	}
}
func TestGetRoomWithTenant(t *testing.T) {
	db := initDB()
	db.MustExec("INSERT INTO articles (name) VALUES ('hoge')")
	db.MustExec("INSERT INTO articles (name) VALUES ('fuga')")
	db.MustExec("INSERT INTO rooms (article_id, 'index', name, rent) VALUES (1, 1, '101', 10000)")
	db.MustExec("INSERT INTO rooms (article_id, 'index', name, rent) VALUES (1, 2, '201', 10000)")
	db.MustExec("INSERT INTO rooms (article_id, 'index', name, rent) VALUES (2, 1, '101', 10000)")
	db.MustExec("INSERT INTO tenants (room_id, name, rent, since, until) VALUES (1, 'Bob', 10000, '1990-01-01', '2010-01-01')")

	query := `{
		room(id: 1) {
			id
			name
			tenant(now: "2000-01-01") {
				id
				name
				rent
				since
				until
			}
		}
	}`
	params := graphql.Params{Schema: gqli.Schema(), RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		t.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	bjson, _ := json.Marshal(r)
	json := string(bjson)
	want := `{"data":{"room":{"id":1,"name":"101","tenant":{"id":1,"name":"Bob","rent":10000,"since":"1990-01-01","until":"2010-01-01"}}}}`
	if json != want {
		t.Fatalf("invalid result,\nhave %s\nwant %s\n", json, want)
	}
}

func TestGetRoomWithTenantHistory(t *testing.T) {
	db := initDB()
	db.MustExec("INSERT INTO articles (name) VALUES ('hoge')")
	db.MustExec("INSERT INTO rooms (article_id, 'index', name, rent) VALUES (1, 1, '101', 10000)")
	db.MustExec("INSERT INTO tenants (room_id, name, rent, since, until) VALUES (1, 'Bob', 10000, '1990-01-01', '2000-01-01')")
	db.MustExec("INSERT INTO tenants (room_id, name, rent, since, until) VALUES (1, 'Jane', 10000, '2000-01-01', '2010-01-01')")
	db.MustExec("INSERT INTO tenants (room_id, name, rent, since, until) VALUES (1, 'Tom', 10000, '2010-01-01', '2020-01-01')")

	query := `{
		room(id: 1) {
			id
			name
			tenantHistory(since: "1990-01-01", until: "2010-01-01") {
				id
				name
			}
		}
	}`
	params := graphql.Params{Schema: gqli.Schema(), RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		t.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	bjson, _ := json.Marshal(r)
	json := string(bjson)
	want := `{"data":{"room":{"id":1,"name":"101","tenantHistory":[{"id":1,"name":"Bob"},{"id":2,"name":"Jane"}]}}}`
	if json != want {
		t.Fatalf("invalid result,\nhave %s\nwant %s\n", json, want)
	}
}

func TestGetTenantWithBills(t *testing.T) {
	db := initDB()
	db.MustExec("INSERT INTO articles (name) VALUES ('hoge')")
	db.MustExec("INSERT INTO rooms (article_id, 'index', name, rent) VALUES (1, 1, '101', 10000)")
	db.MustExec("INSERT INTO tenants (room_id, name, rent, since, until) VALUES (1, 'Bob', 10000, '1990-01-01', '2000-01-01')")
	db.MustExec("INSERT INTO billing_terms (name, since, until) VALUES ('January', '2010-01-01', '2010-01-31')")
	db.MustExec("INSERT INTO bills (billing_term_id, tenant_id, until, since, rent) VALUES (1, 1, '2010-01-01', '2010-01-31', 10000)")

	query := `{
		tenant(id: 1) {
			id
			name
			bills {
				until
				since
				rent
			}
		}
	}`
	params := graphql.Params{Schema: gqli.Schema(), RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		t.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	bjson, _ := json.Marshal(r)
	json := string(bjson)
	want := `{"data":{"tenant":{"bills":[{"rent":10000,"since":"2010-01-31","until":"2010-01-01"}],"id":1,"name":"Bob"}}}`
	if json != want {
		t.Fatalf("invalid result,\nhave %s\nwant %s\n", json, want)
	}
}
