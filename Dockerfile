FROM postgres:14.1


RUN apt-get update && apt-get  install -y postgresql-14-postgis-3  
RUN echo "hello there"

CMD ["/usr/local/bin/docker-entrypoint.sh","postgres"]