package postgres

import (
	"context"
	"os"
	"testing"

	"github.com/Hack-Hack-geek-Vol10/services/cmd/config"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestConnection(t *testing.T) {
	config.LoadEnv()

	conn := NewConnection()
	db, err := conn.Connection()
	if err != nil {
		t.Fatalf("failed to connect to database :%v", err)
	}

	t.Log("success to connect database")

	if db.Ping() != nil {
		t.Fatalf("failed to ping database :%v", err)
	}

	t.Log("success to ping database")

	if conn.Close(context.Background()) != nil {
		t.Fatalf("failed to close database :%v", err)
	}

	t.Log("success to close database")
}
