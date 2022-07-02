FROM golang:1.18 AS app
WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app

FROM scratch
COPY --from=app /app/app /app/app
EXPOSE 1117
CMD ["/app/app"]