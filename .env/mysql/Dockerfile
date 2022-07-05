FROM mysql:8.0.29

RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 467B942D3A79BD29
RUN apt-get update \
    && apt-get install -y locales \
    && locale-gen ja_JP.UTF-8 \
    && echo "export LANG=ja_JP.UTF-8" >> ~/.bashrc \
    && localedef -f UTF-8 -i ja_JP ja_JP.UTF-8

ENV MYSQL_ROOT_PASSWORD rootpass
ENV MYSQL_DATABASE test
ENV MYSQL_USER testuser
ENV MYSQL_PASSWORD testpass

COPY config.toml /etc/mysql/conf.d/my.cnf

EXPOSE 3306

# 参考サイト
# https://zenn.dev/jojojo/articles/f1223bb06cf5be