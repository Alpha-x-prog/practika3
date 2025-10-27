FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM gcr.io/distroless/base-debian12
ENV PORT=5002
EXPOSE 5002
COPY --from=build /src/app /app
USER nonroot:nonroot
ENTRYPOINT ["/app"]

