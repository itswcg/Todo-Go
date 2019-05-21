FROM alpine
MAINTAINER itswcg

COPY . /todo
WORKDIR /todo

EXPOSE 8000
ENTRYPOINT ["./Todo-Go"]