# GGT Stack Template (Gin, Gorm, TypeScript)

Template project for exposing Http endpoints with CRUD functionality.
The CRUD layer's functionality works generically with a given Go struct.
To enhance fullstack rapid prototyping, a converter tool is provided to convert said Go structs to TypeScript models. Generated TypeScript models are outputted to the `models` directory.

Example includes Get, GetAll, Create, Update, and Delete endpoints for an employee.

Communication is handled by [Gin](https://github.com/gin-gonic/gin), CRUD by [GORM](https://github.com/go-gorm/gorm), and Go struct to TS model conversion by [Typescriptify-golang-structs](https://github.com/tkrajina/typescriptify-golang-structs).

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
curl -X PUT localhost:8664/api/v1/employee/1 -d '{"Name":  "jasonnolonger"}'
curl -X DELETE localhost:8664/api/v1/employee/
```

Todo: 
- Auth alongside a user model

