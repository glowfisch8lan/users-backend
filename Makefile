include .env

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)

export POSTGRESQL_URL='postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable'
#export VENDOR=${ROOT_DIR}/packages
#export GOPATH=${VENDOR}
#export GOCACHE=${ROOT_DIR}/.cache

#export GO111MODULE=auto
#export GOROOT=${VENDOR}/go
#export BIN=/usr/lib/go-1.19/bin
#export GOENV=${VENDOR}/go/go.env
#export BIN=${VENDOR}/go/bin

#install/go:
#	rm -rf ${VENDOR}/tmp; \
#	mkdir -p ${VENDOR}/tmp; \
#	cd ${VENDOR}/tmp;\
#	wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz; \
#	tar zxvf go1.21.6.linux-amd64.tar.gz; \
#	mv ${VENDOR}/tmp/go ${VENDOR}/go; \
#	rm -rf ${VENDOR}/tmp;

run/dev/watch:
	${AIR_BIN}/air -c air.toml
run/dev:
	go run ./cmd/users

migrate/install:
	wget https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.deb -P /tmp/;\
	dpkg -i /tmp/migrate.linux-amd64.deb; \
	rm -rf /tmp/migrate.linux-amd64.deb;

migrate/create:
	migrate create -ext sql -dir migrations -seq $(RUN_ARGS)

migrate/up:
	migrate -database ${POSTGRESQL_URL} -path migrations up

migrate/force:
	migrate -database ${POSTGRESQL_URL} -path migrations force $(RUN_ARGS)

migrate/drop:
	migrate -database ${POSTGRESQL_URL} -path migrations drop
