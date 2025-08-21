FROM golang:1.23-alpine as builder

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go mod download
RUN go build -v -o consoleapp ./cmd/.

#############################################

FROM alpine:latest

ARG vc_user
ARG vc_pass

ENV vc_user=$vc_user
ENV vc_pass=$vc_pass
# Variable de entorno para manejar la hora de Colombia
ENV TZ=America/Bogota

RUN mkdir /app
RUN apk update && apk add --no-cache tzdata

COPY --from=builder /app/consoleapp /app/consoleapp

EXPOSE 80

CMD [ "/app/consoleapp" ]