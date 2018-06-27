FROM krallin/ubuntu-tini:trusty

RUN mkdir /app

COPY public /app/

COPY logkit.conf /app/

COPY logkit /app/

RUN chmod a+x /app/logkit 

COPY run.sh /app/

RUN chmod a+x /app/run.sh

RUN sed -i -- 's/localhost//g' /app/logkit.conf

EXPOSE 3000

WORKDIR /app

ENTRYPOINT ["/app/run.sh"]
