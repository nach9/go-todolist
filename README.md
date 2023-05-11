# TODO APP

## How To Run

You can run by using Docker run by including below environment
```
MYSQL_HOST
MYSQL_PORT
MYSQL_USER
MYSQL_PASSWORD
MYSQL_DBNAME
```

example
```
docker run --network="host" -e MYSQL_HOST=127.0.0.1 -e MYSQL_PORT=3306 -e MYSQL_USER=user -e MYSQL_PASSWORD=secret -e MYSQL_DBNAME=memberapp nach9/todoapp:1.0
```