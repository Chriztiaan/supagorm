# Supagorm (with TS) template

Schema/Migration management for Supabase projects, using [GORM](https://github.com/go-gorm/gorm).
TypeScript types for tables are generated via [Typescriptify-golang-structs](https://github.com/tkrajina/typescriptify-golang-structs).

Generated TS models are stored under the `models` dir.

Install
```
go mod tidy
```

Run:
```
go run .
```