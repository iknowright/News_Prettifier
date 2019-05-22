[![Build Status](https://semaphoreci.com/api/v1/ayoo/go-mux-postgres-rest/branches/master/shields_badge.svg)](https://semaphoreci.com/ayoo/go-mux-postgres-rest)


Simple RESTful API using Golang and Postgres

##### to test locally
$ export TEST_DB_NAME=db_name TEST_DB_USERNAME=username TEST_DB_PASSWORD=password

##### to run the serve locally
$ export APP_DB_NAME=db_name APP_DB_USERNAME=username APP_DB_PASSWORD=password

$ go install

$ $GOPATH/bin/mux-postgres-rest # API server runs on port 8000