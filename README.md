# DrakoHacks
## pgx-pgjsontogo

This is a new example to execute a query and result convert in JSON in database server, so using [pgxv4](https://github.com/jackc/pgx) driver just extract a json in Go.

### Execute the example

Just clone project and change database connection settings like this: 

```
$ git clone https://github.com/DrakoRod/pgx-pgjsontogo
$ cd pgx-rowstojson/
```

Change database connections settings in code

```
urlExample := "postgres://postgres:s3cr3t@localhost:5432/postgres"

```
Execute main.go to test

```
$ go run main.go
JSON Result::>  [{"datname":"template1","datowner":"postgres"},{"datname":"template0","datowner":"postgres"},{"datname":"explain","datowner":"drakorod"}{"datname":"drakorod","datowner":"drakorod"},{"datname":"go_project","datowner":"corvus_user"}]
```

#### Enjoy! :)
