FROM scratch
MAINTAINER itswcg

COPY . /Todo-Go
WORKDIR /Todo-Go

EXPOSE 8512
ENTRYPOINT ["./Todo-Go"]