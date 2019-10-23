FROM scratch

MAINTAINER  "david"

WORKDIR .
ADD docker-golang-demo .
ADD test.toml .

EXPOSE 8082
CMD ["./docker-golang-demo","-config=./test.toml"]
