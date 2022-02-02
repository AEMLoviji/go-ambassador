FROM golang:1.17

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

## DEV: we use https://github.com/cosmtrek/air package to support live reload
## in docker-compose.yml file we have mounted all our application to container 
#
## binary will be $(go env GOPATH)/bin/air
## RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
#RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
#RUN cp ./bin/air /bin/air
#
## air command automatically will run main.go file and compile app if anything in file system changes
#CMD [ "air" ]


# PROD: Uncomment below lines to build and run application as an executable
RUN go build -o /ambassador-backend

EXPOSE 8000

CMD [ "/ambassador-backend" ]