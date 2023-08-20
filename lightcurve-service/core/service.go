package core

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDetections(
	oid string,
	handle_success func(result []Detection),
	handle_error func(err error),
	db *pgxpool.Pool,
) {
	conn, err := db.Acquire(context.Background())
	if err != nil {
		handle_error(err)
	}
	defer conn.Release()
	rows, _ := conn.Query(context.Background(), "select candid, oid, mjd, magpsf, sigmapsf, fid from detection where oid=$1", oid)
	detections, err := pgx.CollectRows(rows, pgx.RowToStructByName[Detection])
	if err != nil {
		handle_error(err)
	} else {
		handle_success(detections)
	}
}

func GetNonDetections(
	oid string,
	handle_success func(result []NonDetection),
	handle_error func(err error),
	db *pgxpool.Pool,
) {
	conn, err := db.Acquire(context.Background())
	if err != nil {
		handle_error(err)
	}
	defer conn.Release()
	rows, _ := conn.Query(context.Background(), "select * from non_detection where oid=$1", oid)
	detections, err := pgx.CollectRows(rows, pgx.RowToStructByName[NonDetection])
	if err != nil {
		handle_error(err)
	} else {
		handle_success(detections)
	}
}
