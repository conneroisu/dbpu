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
package main

import (
    "fmt"
    "os"
    "github.com/conneroisu/dbpu"
    "github.com/google/uuid"
)

func main() {
    orgToken := os.Getenv("ORG_TOKEN")
    orgName := os.Getenv("ORG_NAME")
    dbName := uuid.New().String()
    db, err := dbpu.CreateDatabase(orgToken, orgName, dbName, "default")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(db)
}
```
