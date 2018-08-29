start_local_db:
	pg_ctl -D /usr/local/var/postgres start

process_sql:
	go run ./generate_sql/main.go

import_sql: process_sql
	echo "You may need to create the database first"
	echo "psql -U postgres <<< \"CREATE DATABASE gophersearch;\""
	psql gophersearch < ./database.sql

build:
	CGO_ENABLED=0 GOOS=linux buffalo build --static -o gophersearch

release: build
	goreleaser

run:
	buffalo dev

build_docker:
	docker build -t nicholasjackson/gophersearch:latest .
