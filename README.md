# ApplicationRemitly
Hello, this is my application that allows you to manage SWIFT codes.\
In order to start this app run this commands in terminal in ApplicationRemitly directory.
```
docker-compose build
docker-compose up
```
In different terminal
```
docker-compose run --service-ports api bash
go build
./ApplicationRemitly
```

Use pg4admin to connect to the database\
Regiser server:
 - general:
   name: docker
 - connection:
   hostname: localhost\
   port: 5432\
   username: postgres\
   password: password

In query tool:
- SELECT * FROM swift_code

to see all the data in the database


In http-tests directory there are files with http requests that can be used to test the application.

### Endpoint 1: GET request Return SWIFT code details for a specific SWIFT code
GET http://localhost:8080/v1/swift-codes/ABIEBGS1XXX

### Endpoint 2: GET request Return all SWIFT codes with details for a specific country (both headquarters and branches)
GET http://localhost:8080/v1/swift-codes/country/PL
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

### Endpoint 4: DELETE request to example server
DELETE http://localhost:8080/v1/swift-codes/ABIEBGS1XXX



