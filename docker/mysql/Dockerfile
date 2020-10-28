FROM mysql:8.0.22

EXPOSE 3306

# MySQL設定ファイルをイメージ内にコピー
COPY ./my.cnf /etc/mysql/conf.d/my.cnf

CMD ["mysqld"]