FROM golang:1.19-alpine as build-base
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o ./out/assessment .

FROM alpine:3.16.2
COPY --from=build-base /app/out/assessment /app/assessment
CMD ["/app/assessment"]
