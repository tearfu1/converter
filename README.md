api server which convert currencies
language: go 1.22
database: mysql server version 8.3.0
server, app for filling database from site and database are in docker containers

launch:
run: docker-compose build && docker-compose up -d
to interact with server use: curl -X "method" http://localhost:8080/"url"/"currencyName"(optional) '{}'(optional)
to check database use: 
docker exec -it temp_mysql_1 /bin/bash
mysql -u root -p (password: example_password)

apis:
curl -X GET localhost:8080/getAll

curl -X GET localhost:8080/currencies/"example"

curl -X POST \                                              
  http://localhost:8080/create \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "aaa",
    "rate": 1111.23
}'

curl -X DELETE \
  http://localhost:8080/delete/USD

curl -X PUT \
  http://localhost:8080/update/USD \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "USD",
    "rate": 1.23
}'

curl -X GET  http://localhost:8080/converter -d   '{
    "from": "JPY",
    "to": "USD",
    "amount": 200.0
}'

