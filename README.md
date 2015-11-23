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
sudo docker run -d --name snsrobotd -p 8080:4000 --link snsrobotdb:postgres snsrobotd
```

## Create Develop Environment

```bash
sudo docker run --name db -e POSTGRES_PASSWORD=123456 -d postgres
sudo docker exec -it db /usr/bin/createdb -U postgres snsrobot
git clone https://github.com/quchunguang/snsrobot-go
cd snsrobot-go
sudo docker run -it --name snsrobotd -v "$PWD":/go/snsrobot-go -p 8080:4000 --link db:postgres golang
```

All we need of the docker container `db`, instance of postgresql server,
is create a database named `snsrobot` in it, so it can be shared with other projects.
Now we got the terminal of docker container.
Change source with fun in $PWD of host machine.
And then build/run it in the terminal of the docker container by,

```bash
# install dependences of both sides
go get github.com/lib/pq


# server side build in container
cd snsrobot-go/snsrobotd
go build
./snsrobotd

# client side build in container
cd snsrobot-go/snsrobot
go build
./snsrobot --help
```
