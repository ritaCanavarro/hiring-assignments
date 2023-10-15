FROM golang:1.21.1-alpine AS build

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go test -o Documentkeeper
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o Documentkeeper

FROM scratch

WORKDIR /
COPY --from=build /build/Documentkeeper /

ENTRYPOINT [ "/Documentkeeper" ]