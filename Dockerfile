FROM golang
RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee
EXPOSE 4000
