ARG KASPAD_IMAGE
ARG KASPAMINER_IMAGE

FROM ${KASPAD_IMAGE} as karlsend
FROM ${KASPAMINER_IMAGE} as karlsenminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/karlsend/PYVERT/testfork/karlsend/

WORKDIR /go/src/github.com/karlsend/PYVERT/testfork/karlsend/

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=karlsend /app/ /app/
COPY --from=karlsenminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/karlsend/PYVERT/testfork/karlsend/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]
