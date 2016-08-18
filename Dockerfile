FROM debian:wheezy

EXPOSE 8000

ADD ./bin/ipam /bin/ipam

ENTRYPOINT /bin/ipam --mongo mongodb:27017
