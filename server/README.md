## Getting Started

The backend is written in [go](https://golang.org/). Download go (version 1.13+) in order to run the backend

## Setup for local dev

These instructions are all under the assumption you're using macOS. Most of these steps should work for linux

### Writing & executing go

Go is a language that has a code formatter baked into the language. For best experience, I recommend using 
[vs code](https://code.visualstudio.com/) as an editor with the 
[vscode-go](https://code.visualstudio.com/docs/languages/go) plugin. In your settings set
`go.autocompleteUnimportedPackages` to `true` as well as `editor.formatOnSave` to `true`.

You should also add the `~/go/bin` to your bash path
    - `export PATH=$PATH:~/go/bin`

This is so you could run `go install` and then run the executables without having to use the full file path

### PostgreSQL

1. Install PostgreSQL 11.x ([macOS download](https://postgresapp.com/),
 [windows download](https://www.enterprisedb.com/downloads/postgres-postgresql-downloads))
2. Make a nsl_dev database `createdb nsl_dev`
    - You may need to switch your user to `postgres` by running `su - postgres`
3. Install the [migrate](https://github.com/golang-migrate/migrate) and migrate the database to the latest state
    1. `go get github.com/golang-migrate/migrate`
    2. `migrate -source file://sql_migrations -database postgres://localhost:5432/nsl_dev up`
4. Seed (ie: fill) the database with test data (WIP)

In order to make changes to the database, use the migrate tool to create new migration files
    - `migrate create <migration_name>`

### Questions?

If you have any questions about any of these steps, please contact Mike

## Running the Program

Running the backend is simple.

1. In the terminal run `go install`.
2. type `server` to start a local webserver. The default local address will be `127.0.0.1:4000`

## Testing

You're expected to write tests for exported functions (functions that start with an uppercase letter)
