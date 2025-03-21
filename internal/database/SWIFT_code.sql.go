// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: SWIFT_code.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const checkSWIFTCodeExists = `-- name: CheckSWIFTCodeExists :one
SELECT EXISTS (SELECT 1 FROM SWIFT_code WHERE swiftCode = $1)
`

func (q *Queries) CheckSWIFTCodeExists(ctx context.Context, swiftcode string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkSWIFTCodeExists, swiftcode)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createSWIFTCodeEntry = `-- name: CreateSWIFTCodeEntry :one
INSERT INTO SWIFT_code (id, countryCode, swiftCode, codeType, name, address, townName, timeZone, countryName, isHeadquarter, associatedWith)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING id, countrycode, swiftcode, codetype, name, address, townname, timezone, countryname, isheadquarter, associatedwith
`

type CreateSWIFTCodeEntryParams struct {
	ID             uuid.UUID
	Countrycode    string
	Swiftcode      string
	Codetype       string
	Name           string
	Address        string
	Townname       string
	Timezone       string
	Countryname    string
	Isheadquarter  bool
	Associatedwith int32
}

func (q *Queries) CreateSWIFTCodeEntry(ctx context.Context, arg CreateSWIFTCodeEntryParams) (SwiftCode, error) {
	row := q.db.QueryRowContext(ctx, createSWIFTCodeEntry,
		arg.ID,
		arg.Countrycode,
		arg.Swiftcode,
		arg.Codetype,
		arg.Name,
		arg.Address,
		arg.Townname,
		arg.Timezone,
		arg.Countryname,
		arg.Isheadquarter,
		arg.Associatedwith,
	)
	var i SwiftCode
	err := row.Scan(
		&i.ID,
		&i.Countrycode,
		&i.Swiftcode,
		&i.Codetype,
		&i.Name,
		&i.Address,
		&i.Townname,
		&i.Timezone,
		&i.Countryname,
		&i.Isheadquarter,
		&i.Associatedwith,
	)
	return i, err
}

const deleteSWIFTCodeEntry = `-- name: DeleteSWIFTCodeEntry :exec
DELETE FROM SWIFT_code WHERE swiftCode = $1
`

func (q *Queries) DeleteSWIFTCodeEntry(ctx context.Context, swiftcode string) error {
	_, err := q.db.ExecContext(ctx, deleteSWIFTCodeEntry, swiftcode)
	return err
}

const getAllSWIFTCodes = `-- name: GetAllSWIFTCodes :many
SELECT swiftCode FROM SWIFT_code
`

func (q *Queries) GetAllSWIFTCodes(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getAllSWIFTCodes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var swiftcode string
		if err := rows.Scan(&swiftcode); err != nil {
			return nil, err
		}
		items = append(items, swiftcode)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBranchesByAssociatedWith = `-- name: GetBranchesByAssociatedWith :many
SELECT id, countrycode, swiftcode, codetype, name, address, townname, timezone, countryname, isheadquarter, associatedwith FROM SWIFT_code WHERE associatedWith = $1
`

func (q *Queries) GetBranchesByAssociatedWith(ctx context.Context, associatedwith int32) ([]SwiftCode, error) {
	rows, err := q.db.QueryContext(ctx, getBranchesByAssociatedWith, associatedwith)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SwiftCode
	for rows.Next() {
		var i SwiftCode
		if err := rows.Scan(
			&i.ID,
			&i.Countrycode,
			&i.Swiftcode,
			&i.Codetype,
			&i.Name,
			&i.Address,
			&i.Townname,
			&i.Timezone,
			&i.Countryname,
			&i.Isheadquarter,
			&i.Associatedwith,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEntryByCountryCode = `-- name: GetEntryByCountryCode :many
SELECT id, countrycode, swiftcode, codetype, name, address, townname, timezone, countryname, isheadquarter, associatedwith FROM SWIFT_code WHERE countryCode = $1
`

func (q *Queries) GetEntryByCountryCode(ctx context.Context, countrycode string) ([]SwiftCode, error) {
	rows, err := q.db.QueryContext(ctx, getEntryByCountryCode, countrycode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SwiftCode
	for rows.Next() {
		var i SwiftCode
		if err := rows.Scan(
			&i.ID,
			&i.Countrycode,
			&i.Swiftcode,
			&i.Codetype,
			&i.Name,
			&i.Address,
			&i.Townname,
			&i.Timezone,
			&i.Countryname,
			&i.Isheadquarter,
			&i.Associatedwith,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEntryBySWIFTCode = `-- name: GetEntryBySWIFTCode :one
SELECT id, countrycode, swiftcode, codetype, name, address, townname, timezone, countryname, isheadquarter, associatedwith FROM SWIFT_code WHERE swiftCode = $1
`

func (q *Queries) GetEntryBySWIFTCode(ctx context.Context, swiftcode string) (SwiftCode, error) {
	row := q.db.QueryRowContext(ctx, getEntryBySWIFTCode, swiftcode)
	var i SwiftCode
	err := row.Scan(
		&i.ID,
		&i.Countrycode,
		&i.Swiftcode,
		&i.Codetype,
		&i.Name,
		&i.Address,
		&i.Townname,
		&i.Timezone,
		&i.Countryname,
		&i.Isheadquarter,
		&i.Associatedwith,
	)
	return i, err
}
