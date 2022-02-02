FROM golang:1.17

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

#COPY *.go ./
COPY . ./
RUN go build -o /ambassador-backend

EXPOSE 8000

CMD [ "/ambassador-backend" ]