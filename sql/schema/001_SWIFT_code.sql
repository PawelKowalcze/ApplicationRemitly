-- +goose Up
CREATE TABLE SWIFT_code (
    id UUID PRIMARY KEY,
    countryCode TEXT NOT NULL,
    swiftCode TEXT NOT NULL,
    codeType TEXT NOT NULL,
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    townName TEXT NOT NULL,
    timeZone TEXT NOT NULL,
    countryName TEXT NOT NULL,
    isHeadquarter BOOLEAN NOT NULL,
    associatedWith INTEGER NOT NULL
);

-- +goose Down
DROP TABLE SWIFT_code;
