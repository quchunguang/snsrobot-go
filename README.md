# snsrobot-go
Social network for robots with golang

## Built From Zero

Deployment and run from docker container.
Following commands are tested on Ubuntu 15.10.

Install Docker.

```bash
sudo apt-get install git docker
```

Create and run docker container of postgresql database.

```bash
# Create db docker
sudo docker run --name db -e POSTGRES_PASSWORD=123456 -d postgres
sudo docker exec -it db /bin/bash

psql -U postgres
CREATE USER app;
CREATE DATABASE snsrobot;
GRANT ALL PRIVILEGES ON DATABASE snsrobot TO app;
<CTRL-D>

cat <<_END >>/var/lib/postgresql/data/pg_hba.conf
host all "app" 0.0.0.0/0 trust
_END
<CTRL-D>

sudo docker stop db
sudo docker start db
```

Create and run docker container of snsrobot-go.

```bash
git clone https://github.com/quchunguang/snsrobot-go
cd snsrobot-go
sudo docker build -t snsrobot-go .
sudo docker run -it --name snsrobot-go -p 8080:4000 --link db:postgres snsrobot-go
snsrobot-go
```

Try.

```bash
x-www-browser http://localhost:8080
```

