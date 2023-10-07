FROM golang:1.21.1-alpine AS build
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build --ldflags '-extldflags "-static"' -o documentkeeper

FROM scratch
WORKDIR /
COPY --from=build /build/documentkeeper /

ENTRYPOINT ["/documentkeeper"]
