package main

import (
	"context"
	"e-commerce-api-golang/internal/env"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found")
	}

	ctx := context.Background()

	cgf := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "user=postgres password=postgres dbname=ecommerce-golang sslmode=disable"),
		},
	}

	conn, err := pgx.Connect(ctx, cgf.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	api := application{
		config: cgf,
		db: conn,
	}

	logger.Info("Connected to Database", "dsn", cgf.db.dsn)

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server has failed to start", "error", err)
		os.Exit(1)
	}
}
