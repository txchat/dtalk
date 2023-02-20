package dao

import (
	"os"
	"testing"
)

var (
	mysqlRootPassword string
)

func TestMain(m *testing.M) {
	mysqlRootPassword = os.Getenv("MYSQL_ROOT_PASSWORD")

	os.Exit(m.Run())
}
