FROM rhaps1071/golang-1.14-alpine-git AS build
WORKDIR /
COPY . .
RUN CGO_ENABLED=0 go get -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o ./app

FROM scratch
COPY --from=build /go/bin/dlv /dlv
COPY --from=build /app /app
ENTRYPOINT [ "/dlv" ]