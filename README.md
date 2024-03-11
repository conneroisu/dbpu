# go-dbpu

Golang database per user for turso db.

<img src='./assets/dbpu.png' width='50'>

## Installation

```bash
go get github.com/conneroisu/dbpu
```

## Usage

See examples for more details.

### Create a new database

The following example creates a new database.

```go
// func CreateDatabase(orgToken string, orgName string, name string, group string) (Db, error) {
db, err := dbpu.CreateDatabase("orgToken", "orgName", "dbName", "group")
if err != nil {
    log.Fatal(err)
}
```
