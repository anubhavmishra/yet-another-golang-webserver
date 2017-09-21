FROM alpine:3.1
MAINTAINER Anubhav Mishra <anubhavmishra@me.com>
ADD build/linux/amd64/yet-another-golang-webserver /usr/bin/yet-another-golang-webserver
ENTRYPOINT ["yet-another-golang-webserver"]