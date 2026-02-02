FROM golang:1.25 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/url-short-url cmd/app/main.go

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/url-short-url /app/url-short-url

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/url-short-url"]