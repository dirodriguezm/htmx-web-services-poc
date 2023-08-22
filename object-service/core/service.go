package core

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetObjectById(id string,
	handle_success func(result Object),
	handle_error func(err error),
	db *pgxpool.Pool,
) {
	conn, err := db.Acquire(context.Background())
	if err != nil {
		handle_error(err)
	}
	defer conn.Release()
	row := conn.QueryRow(context.Background(), "select  oid, corrected, stellar, ndet, meanra, meandec, firstmjd, lastmjd from object where oid=$1", id)
	var object Object
	err = row.Scan(&object.Oid, &object.Corrected, &object.Stellar, &object.Ndet, &object.Meanra, &object.Meandec, &object.Firstmjd, &object.Lastmjd)
	if err != nil {
		handle_error(err)
	} else {
		handle_success(object)
	}
}
