####################
#     BUILD GO     #
####################
FROM golang:1.14-alpine as go-builder

WORKDIR "/app"

COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

####################
# BUILD JAVASCRIPT #
####################
FROM node:12.16-alpine as js-builder

WORKDIR "/app/client"

COPY client/package*.json client/yarn.lock ./
RUN yarn
COPY client ./

RUN yarn build

###################
#    SERVE APP    #
###################
FROM alpine
LABEL maintainer="Yiğit Sözer <sozer.work@gmail.com>"

WORKDIR "/app"

COPY --from=go-builder /app/main ./
COPY --from=js-builder /app/client/build ./client/build

EXPOSE 8000
CMD ["./main", "dockermode"]