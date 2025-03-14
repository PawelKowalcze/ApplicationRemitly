# ApplicationRemitly

install postgresql
install pgadmin4
in pgadmin4 create a database called ApplicationRemitly

sql\schema> goose postgres postgres://pawel:pass@localhost:5432/ApplicationRemitly up
C:\Users\kowal\OneDrive\Pulpit\Go\ApplicationRemitly> sqlc generate


### GET request to example server
GET http://localhost:8080/v1/swift-codes
Code: SWIFTCode ALBPPLP1BMW
###

### POST request to example server
POST http://localhost:8080/v1/swift-codes

{
"address": "Cracow",
"bankName": "ING",
"countryISO2": "PL",
"countryName": "Poland",
"isHeadquarter": true,
"swiftCode": "INGOPLPWXXX"
}