FROM centos:7

RUN yum install -y epel-release ; yum install -y zeromq;yum clean all; rm -rf /var/cache/yum
WORKDIR /
COPY server /
CMD /server $ENDPOINT $SERVERID
