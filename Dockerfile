FROM golang:1.22.2-alpine3.19 AS Build

WORKDIR /app

COPY . .

RUN go mod download && \
  go build -o ./bin/trigrams


FROM alpine:3.19

ARG USER=appuser
RUN adduser -h /home/${USER} -s /bin/sh -u 1001 ${USER} -D

WORKDIR /
COPY --from=Build --chown=${USER}:${USER} /app/bin/trigrams /usr/local/bin/trigrams

ENTRYPOINT [ "/usr/local/bin/trigrams" ]
