FROM golang:stretch

RUN apt-get update && apt-get install -y postgresql-client-9.6

# if left blank app will run with dev settings
# to build production image run:
# $ docker build ./api --build-args app_env=production
ARG app_env
ENV APP_ENV $app_env

# it is okay to leave user/GoDoRP as long as you do not want to share code with other libraries
COPY . /go/src/apigo/back/
WORKDIR /go/src/apigo/back/

RUN go get ./database

RUN go get ./
RUN go build

COPY ./entrypt.sh /usr/local/bin/docker-entrypoint
RUN chmod +x /usr/local/bin/docker-entrypoint

ENTRYPOINT [ "docker-entrypoint" ]
	
EXPOSE 5000
