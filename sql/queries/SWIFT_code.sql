-- name: CreateSWIFTCodeEntry :one
INSERT INTO SWIFT_code (id, countryCode, swiftCode, codeType, name, address, townName, timeZone, countryName, isHeadquarter, associatedWith)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: GetEntryBySWIFTCode :one
SELECT * FROM SWIFT_code WHERE swiftCode = $1;

-- name: CheckSWIFTCodeExists :one
SELECT EXISTS (SELECT 1 FROM SWIFT_code WHERE swiftCode = $1);

-- name: GetBranchesByAssociatedWith :many
SELECT * FROM SWIFT_code WHERE associatedWith = $1;

-- name: DeleteSWIFTCodeEntry :exec
DELETE FROM SWIFT_code WHERE swiftCode = $1;

-- name: GetEntryByCountryCode :many
SELECT * FROM SWIFT_code WHERE countryCode = $1;

-- name: GetAllSWIFTCodes :many
SELECT swiftCode FROM SWIFT_code;