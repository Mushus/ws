package mysql

import (
	"reflect"
	"testing"
)

func TestSchema(t *testing.T) {
	d := DataStore{
		models: models,
	}

	want := []string{
		"CREATE TABLE `top` (\n" +
			"\t`id` VARCHAR(24) NOT NULL PRIMARY KEY,\n" +
			"\t`title` VARCHAR(255) NOT NULL,\n" +
			"\t`sentence` VARCHAR(255) NOT NULL,\n" +
			"\t`header_title` VARCHAR(255) NOT NULL,\n" +
			"\t`header_description` VARCHAR(255) NOT NULL\n" +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;\n",
		"CREATE TABLE `top_main_prop` (\n" +
			"\t`id` VARCHAR(24) NOT NULL PRIMARY KEY,\n" +
			"\t`top_id` VARCHAR(24) NOT NULL,\n" +
			"\t`order` BIGINT NOT NULL,\n" +
			"\t`value` VARCHAR(255) NOT NULL\n" +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;\n",
		"CREATE TABLE `top_photo_prop` (\n" +
			"\t`id` VARCHAR(24) NOT NULL PRIMARY KEY,\n" +
			"\t`top_id` VARCHAR(24) NOT NULL,\n" +
			"\t`order` BIGINT NOT NULL,\n" +
			"\t`image` VARCHAR(255) NOT NULL,\n" +
			"\t`title` VARCHAR(255)\n" +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;\n",
	}

	got := d.Schemas()
	if reflect.DeepEqual(want, got) {
		t.Fatalf("DataStore.Schema() = %q, want %q", got, want)
	}
}
