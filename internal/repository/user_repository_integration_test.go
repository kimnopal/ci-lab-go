//go:build integration

package repository

import (
	"context"
	"os"
	"testing"

	"github.com/kimnopal/ci-lab-go/internal/database"
)

func TestRepositoryCreateAndList(t *testing.T) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL is not set")
	}

	db, err := database.Open(dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()
	if err := database.EnsureSchema(ctx, db); err != nil {
		t.Fatal(err)
	}

	if _, err := db.ExecContext(ctx, "TRUNCATE TABLE users RESTART IDENTITY"); err != nil {
		t.Fatal(err)
	}

	repo := NewRepository(db)
	created, err := repo.Create(ctx, "Naufal", "naufal@example.com")
	if err != nil {
		t.Fatal(err)
	}

	users, err := repo.List(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(users) != 1 || users[0].ID != created.ID {
		t.Fatalf("unexpected users: %+v", users)
	}
}
