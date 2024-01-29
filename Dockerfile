FROM golang:1.21.2

WORKDIR /rest_project

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

RUN go build -o ./bin/app.exe -v ./cmd/main.go

EXPOSE 4060

CMD [ "./bin/app.exe" ]