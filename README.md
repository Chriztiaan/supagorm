This is a PoC Gin/GORM microservice.
It provides an HTTP API interface via Gin and has ORM functionality via GORM.

Install
```
go mod tidy
```

Run:
```
go run .
```

Example HTTP calls:
```
curl localhost:8664  
curl localhost:8664/api/v1/employee/1
curl -X POST localhost:8664/api/v1/employee -d '{"name": "jason"}'
curl -X PUT localhost:8664/api/v1/employee/1 -d '{"Name":  "asd"}'
curl -X DELETE localhost:8664/api/v1/employee/
```

Todo: 
[] Auth
[] POC gRPC API
