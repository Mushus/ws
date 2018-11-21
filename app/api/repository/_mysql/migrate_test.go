package mysql

import (
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestSchemaTypeName(t *testing.T) {
	s := Column{
		Type: "int(11)",
	}
	got := s.TypeName()
	want := "int"
	if got != want {
		t.Fatalf("schema.TypeName() = %q, want %q", got, want)
	}
}

func TestSchemaTypeSize(t *testing.T) {
	s := schemaRecord{
		Type: "varchar(255)",
	}
	got := s.Size()
	want := 255
	if got != want {
		t.Fatalf("schema.Size() = %d, want %d", got, want)
	}
}

func TestSchemaIsNullTest(t *testing.T) {
	cases := []struct {
		schema schemaRecord
		want   bool
	}{
		{
			schema: schemaRecord{
				Null: "YES",
			},
			want: true,
		},
		{
			schema: schemaRecord{
				Null: "NO",
			},
			want: false,
		},
	}

	for _, c := range cases {
		s := c.schema
		got := s.IsNull()
		if got != c.want {
			t.Fatalf("schma.IsNull() = %v, want %v", got, c.want)
		}
	}
}

// TestToCreateSQL Create文を発行するテスト
func TestToCreateSQL(t *testing.T) {
	in := tableSchemas{
		{
			name: "top",
			columns: columnSchemas{
				{name: "id", typ: "VARCHAR", size: 36, nullable: false, primarykey: true, uniquekey: false, index: false, foeignKey: foeignKey{table: "", column: ""}},
				{name: "title", typ: "VARCHAR", size: 255, nullable: false, primarykey: false, uniquekey: false, index: false, foeignKey: foeignKey{table: "", column: ""}},
				{name: "sentence", typ: "VARCHAR", size: 255, nullable: false, primarykey: false, uniquekey: false, index: false, foeignKey: foeignKey{table: "", column: ""}},
				{name: "header_title", typ: "VARCHAR", size: 255, nullable: false, primarykey: false, uniquekey: false, index: false, foeignKey: foeignKey{table: "", column: ""}},
				{name: "header_description", typ: "VARCHAR", size: 255, nullable: false, primarykey: false, uniquekey: false, index: false, foeignKey: foeignKey{table: "", column: ""}}},
		},
		{
			name: "top_main_prop",
			columns: columnSchemas{
				{name: "top_id", typ: "VARCHAR", size: 36, nullable: false, primarykey: false, uniquekey: false, index: false, foeignKey: foeignKey{table: "top", column: "id"}},
				{name: "order", typ: "BIGINT", size: 0, nullable: false, primarykey: false, uniquekey: false, index: false, foeignKey: foeignKey{table: "", column: ""}},
				{name: "value", typ: "VARCHAR", size: 255, nullable: false, primarykey: false, uniquekey: false, index: false, foeignKey: foeignKey{table: "", column: ""}},
			},
		},
		{
			name: "top_photo_prop",
			columns: columnSchemas{
				{name: "top_id", typ: "VARCHAR", size: 36, nullable: false, primarykey: false, uniquekey: false, index: false, foeignKey: foeignKey{table: "top", column: "id"}},
				{name: "order", typ: "BIGINT", size: 0, nullable: false, primarykey: false, uniquekey: false, index: false, foeignKey: foeignKey{table: "", column: ""}},
				{name: "image", typ: "VARCHAR", size: 255, nullable: false, primarykey: false, uniquekey: false, index: false, foeignKey: foeignKey{table: "", column: ""}},
				{name: "title", typ: "VARCHAR", size: 255, nullable: true, primarykey: false, uniquekey: false, index: false, foeignKey: foeignKey{table: "", column: ""}},
			},
		},
	}
	out := []string{
		"CREATE TABLE `top` (\n\t`id` VARCHAR(36) NOT NULL PRIMARY KEY,\n\t`title` VARCHAR(255) NOT NULL,\n\t`sentence` VARCHAR(255) NOT NULL,\n\t`header_title` VARCHAR(255) NOT NULL,\n\t`header_description` VARCHAR(255) NOT NULL\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;\n",
		"CREATE TABLE `top_main_prop` (\n\t`top_id` VARCHAR(36) NOT NULL,\n\t`order` BIGINT NOT NULL,\n\t`value` VARCHAR(255) NOT NULL\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;\n",
		"CREATE TABLE `top_photo_prop` (\n\t`top_id` VARCHAR(36) NOT NULL,\n\t`order` BIGINT NOT NULL,\n\t`image` VARCHAR(255) NOT NULL,\n\t`title` VARCHAR(255)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;\n",
	}

	got := in.toCreateSQL()
	if reflect.DeepEqual(out, got) {
		t.Errorf("tableSchemas.ToCreateSQL() = %#v, want %#v", got, out)
	}
}
