FROM golang:1.22


RUN mkdir /app
WORKDIR /app

COPY app /app

RUN go mod download
RUN go build -v -o server

EXPOSE 8801
# Run the web service on container startup.
CMD ["/app/server"]