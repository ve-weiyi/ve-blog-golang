#!/bin/bash
if [ ! -d "/var/lib/mysql/blog" ]; then
    mysqld --initialize-insecure --user=mysql --datadir=/var/lib/mysql
    mysqld --daemonize --user=mysql
    sleep 5s
    mysql -uroot -e "create database blog default charset 'utf8' collate 'utf8_bin'; grant all on blog.* to 'root'@'127.0.0.1' identified by '123456'; flush privileges;"
else
    mysqld --daemonize --user=mysql
fi
redis-server &
if [ "$1" = "actions" ]; then
    cd /opt/blog/server && go run main.go &
    cd /opt/blog/web/ && yarn serve &
else
    /usr/sbin/nginx &
    cd /usr/share/nginx/html/ && ./server &
fi
echo "blog ALL start!!!"
tail -f /dev/null
