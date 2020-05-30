# FROM alpine:latest
# MAINTAINER janakh

# ENV LOG_LEVEL=info
# ENV ENV=production
# ENV REPO_URL=github.com/janakhpon/gopherscom
# ENV GOPATH=/app
# ENV APP_PATH=gopherscom
# ENV WORKPATH=$APP_PATH/src
# ENV PORT 8080
# EXPOSE 8080
# RUN mkdir app
# COPY ../gopherscom /app/gopherscom
# RUN chmod +x /app/gopherscom
# WORKDIR /app/gopherscom
# RUN chmod +x /app/gopherscom
# RUN apk add --no-cache git
# RUN apk --no-cache add ca-certificates

# CMD ["go", "run", "main.go"]
# ENTRYPOINT /app/gopherscom
# RUN mkdir /app
# ADD . /app
# WORKDIR /app
# RUN go build -o main .
# CMD ["/app/main"]


# FROM golang:1.14-alpine AS build

# WORKDIR /src/
# RUN mkdir /app
# COPY ./gopherscom /app
# RUN CGO_ENABLED=0 go build -o /bin/demo

# FROM scratch
# COPY --from=build /bin/demo /bin/demo
# ENTRYPOINT ["/bin/demo"]


# RUN apk update && \
#     apk add \
#         bash \
#         build-base \
#         curl \
#         make \
#         git \
#         && rm -rf /var/cache/apk/*

# COPY . /app/gopherscom
# RUN chmod +x /app/gopherscom
# RUN chmod a+x /app/gopherscom
# RUN chmod 700 /app/gopherscom

# ENV PORT 8080
# EXPOSE 8080

# ENTRYPOINT /app/gopherscom

FROM golang:1.13-alpine AS build
MAINTAINER janakh


WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o gopherscom .

FROM alpine:3.9.6 AS server
LABEL maintainer='janakh <jnovaxer@gmail.com>'

WORKDIR /app
COPY --from=build /app/gopherscom .
RUN export DBUSER=testuser
RUN export PASSWORD=testpassword
RUN export DB=testDB
RUN export HOST=localhost:5432
RUN export PORT=5000
RUN export MODE=debug
RUN export COST=16
RUN export SECRET=thisisasecretbetweenmeandgopherscom
RUN export REDIS_SECRET="oops!thisisalittlesecrecybetweenmeandredis"
RUN export REDIS_HOST=localhost:6379
RUN export REDIS_PASSWORD=

CMD ./gopherscom

