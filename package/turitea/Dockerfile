FROM ubuntu
VOLUME /var/local
COPY main /var/
RUN chmod 777 /var/main
EXPOSE 80
WORKDIR /var/local
CMD ["../main"]
