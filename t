#!/bin/bash

function print_help() {
    echo " $0 [command]
Testbox
Available Commands:
  dc	db connect
  dg	db get
  dr	db run
  l	logs
  rl	run local
  rd	run with docker
" >&2
}

#CREATE TABLE bustime(
#  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
#  stationId INT NOT NULL,
#  dateTime DATETIME NOT NULL
#);

case "${1:-}" in
dc)
  docker exec -it monbus_mysql mysql -u root -p1234
;;
dg)
  docker exec -it monbus_mysql mysql -u root -p1234 -e "select * from mysql.bustime;"
;;
dr)
  docker run --name monbus_mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=1234 -d mysql:8.0.21
;;
l)
  docker logs monbus
;;
rl)
  go run cmd/monbus/main.go -logtostderr=true -v 2
;;
rd)
  CGO_ENABLED=0 go build -o bin/monbus cmd/monbus/main.go
  docker build -t quay.io/woohhan/monbus:latest .
  docker push quay.io/woohhan/monbus:latest
  docker rm --force monbus || true
  docker run --name monbus --network=host -d --restart always -e "TZ=Asia/Seoul" quay.io/woohhan/monbus:latest
;;
*)
  print_help
;;
esac
