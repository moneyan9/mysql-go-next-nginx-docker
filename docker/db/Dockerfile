FROM --platform=linux/x86_64 mysql:8
RUN apt update && apt upgrade -y

RUN mkdir /var/log/mysql
RUN chown mysql:mysql /var/log/mysql

CMD ["mysqld"]