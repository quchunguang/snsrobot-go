# snsrobot-go
Social network for robots with golang.

## Deploy SNS Client

```bash
go get -u github.com/quchunguang/snsrobot-go/snsrobot
snsroboot --help

# Or, access the HTTP interface.
x-www-browser http://localhost:8080
```

## Deploy SNS Server with Docker

```bash
sudo docker run -d --name snsrobotdb -e POSTGRES_PASSWORD=123456 postgres
sudo docker exec -it snsrobotdb /usr/bin/createdb -U postgres snsrobot

git clone https://github.com/quchunguang/snsrobot-go
cd snsrobot-go/snsrobotd
sudo docker build -t snsrobotd .
sudo docker run -d --name snsrobotd -p 8080:8080 --link snsrobotdb:postgres snsrobotd
```

## Build Develop Environment

```bash
# run `db` as postgresql container
sudo docker run --name db -e POSTGRES_PASSWORD=123456 -d postgres

# create database `snsrobot` for `snsrobotd`
sudo docker exec -it db /usr/bin/createdb -U postgres snsrobot

# clone source in right GOPATH
go get github.com/quchunguang/snsrobot-go/snsrobot
go get github.com/quchunguang/snsrobot-go/snsrobotd

# run snsrobotd
cd $GOPATH/src/ithub.com/quchunguang/snsrobot-go/snsrobotd/
bee run -gendoc=true -downdoc=true

# run snsrobot
cd $GOPATH/src/ithub.com/quchunguang/snsrobot-go/snsrobot/
go build
./snsrobot --help
```

## Auto-generated api source

This will automatically generate the `snsrobotd` folder.
If the `snsrobotd` folder already exist with source in it. Rename first and merge.

```bash
# get IP address of `db` container
sudo docker inspect db | grep IPAddress

# generate code, 172.17.0.2 is my IP address of `db` container
bee api snsrobotd -driver="postgres" -conn="postgres://postgres:123456@172.17.0.2:5432/snsrobot?sslmode=disable"
cd snsrobotd

# run service with watch by default.
# -gendoc=true  generate document automatically
# -downdoc=true download webtool `swagger` automatically
bee run -gendoc=true -downdoc=true

# HAVE FUN CIUD API!
x-www-browser http://localhost:8080/swagger
```

## Reference

1. [postgresql doc](http://www.postgresql.org/docs/current/static/)
1. [beego doc](http://beego.me/docs/intro/)
1. [bee api](https://github.com/beego/bee#bee-api)
