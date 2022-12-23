FROM golang:1.19.4-alpine3.17

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /usr/src/app ./...
CMD [ "/usr/src/app/imdb-app"]
