FROM golang:alpine AS build-env
RUN apk add --no-cache git
WORKDIR $GOPATH/src/github.com/kenkoii/go-kube
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/* && apk add curl
WORKDIR /app
COPY --from=build-env /app /app

EXPOSE 3000
ENTRYPOINT ./app