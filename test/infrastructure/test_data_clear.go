package infrastructure

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/context"
)

func TruncateTestData(ctx context.Context, dbPool *pgxpool.Pool) {
	_, truncateResultErr := dbPool.Exec(ctx, "TRUNCATE product RESTART IDENTITY")
	if truncateResultErr != nil {
		log.Error("Not product exists on database")
	} else {
		log.Info("Products table truncated")
	}
}
