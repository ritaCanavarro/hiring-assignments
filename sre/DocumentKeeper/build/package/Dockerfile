FROM golang:1.21.1-alpine AS build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go test -o documentkeeper
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o documentkeeper

FROM scratch
WORKDIR /
COPY --from=build /build/documentkeeper /

ENTRYPOINT [ "/documentkeeper" ]