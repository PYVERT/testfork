# -- multistage docker build: stage #1: build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /go/src/github.com/karlsend/PYVERT/testfork/karlsend/

WORKDIR /go/src/github.com/karlsend/PYVERT/testfork/karlsend/

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache karlsend dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o karlsend .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/karlsend/PYVERT/testfork/karlsend/karlsend /app/
COPY --from=build /go/src/github.com/karlsend/PYVERT/testfork/karlsend/infrastructure/config/sample-karlsend.conf /app/

USER root
ENTRYPOINT [ "/sbin/tini", "--" ]

