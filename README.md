# postgres_web_project

PostgresSQL project in progress...

## Requirements 

1. PSQL 9.5 (listening on default port 5432) 
2. Golang Environment ($GOPATH & $GOROOT set)

## Install Instructions

1. Run 'go get' 
2. Run 'sudo su - postgres', cd to the repo and then run 'sh update_db.sh'
3. Regenerate models by running `go generate` (using sqlboiler)
4. Run 'go run coop-cat.go'
