# ApplicationRemitly

install postgresql
install pgadmin4
in pgadmin4 create a database called ApplicationRemitly

sql\schema> goose postgres postgres://pawel:pass@localhost:5432/ApplicationRemitly up
C:\Users\kowal\OneDrive\Pulpit\Go\ApplicationRemitly> sqlc generate


### Endpoint 1: GET request Return all SWIFT codes with details for a specific country (both headquarters and branches)
GET http://localhost:8080/v1/swift-codes/ALBPPLP1BMW
###

### Endpoint 3: POST request  Adds new SWIFT code entries to the database for a specific country
POST http://localhost:8080/v1/swift-codes

{
"address": "Cracow",
"bankName": "ING",
"countryISO2": "PL",
"countryName": "Poland",
"isHeadquarter": true,
"swiftCode": "INGOPLPWXXX"
}


docker-compose up
docker-compose run --service-ports web bash