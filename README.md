# uptime_monitor

uptime monitor service written in golang

## To build Docker Image

```sh
docker build -t uptime_monitor .
```

## To run docker image with .env

```sh
docker run -p 8080:8081 --env-file=.env -it uptime_monitor
```
