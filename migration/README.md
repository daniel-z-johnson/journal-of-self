# Migration
This folder contains the migrations for journal-of-self

## How to run
This uses [goose](https://github.com/pressly/goose) refere to the README.md from that project if anything here is unclear
1. get goose locally `go get -u github.com/pressly/goose/cmd/goose`
1. run in this folder and not project root `goose postgres "dbstring" up`
    - `dbstring` should be replaced with a db string like `user=postgres dbname=postgres sslmode=disable`, journal-of-self uses postgres
1. cheers thats it, [goose](https://github.com/pressly/goose) contains a lot more info about goose's funtionallity that you should checkout
