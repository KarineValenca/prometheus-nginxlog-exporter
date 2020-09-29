# NGINX-monitor

A Prometheus sidecar to add basic but very useful metrics for your NGINX.

## Metrics

The only exposed metrics (for now) are the following:

```
request_seconds_bucket{type, status, isError, errorMessage, method, addr, le}
request_seconds_count{type, status, isError, errorMessage, method, addr}
request_seconds_sum{type, status, isError, errorMessage, method, addr}
response_size_bytes{type, status, isError, errorMessage, method, addr}
```

Where, for a specific request, `type` tells which request protocol was used (`grpc`, `http`, etc), `status` registers the response status, `method` registers the request method, `addr` registers the requested endpoint address, `version` tells which version of your app handled the request, `isError` lets us know if the status code reported is an error or not.

In detail:

1. `request_seconds_bucket` is a metric defines the histogram of how many requests are falling into the well defined buckets represented by the label `le`;

2. `request_seconds_count` is a counter that counts the overall number of requests with those exact label occurrences;

3. `request_seconds_sum` is a counter that counts the overall sum of how long the requests with those exact label occurrences are taking;

4. `response_size_bytes` is a counter that computes how much data is being sent back to the user for a given request type. It captures the response size from the `content-length` response header. If there is no such header, the value exposed as metric will be zero;

## How to

### Config Files

Change your nginx.conf file to add the request time in your logs. In your `log_format` section, add the `rt=$request_time`. It should looks like this:

```
log_format upstream_time '$remote_addr - [$time_local] '
                         '"$request" $status $body_bytes_sent rt=$request_time';
```

Check your NGINX's access.log to see if the request time is being recorded correctly.

After that, create a `config.yml` file to hold all the configuration needed for the NGINX monitor. There is one example of a simple configuration (adjust source files to your own needs):

```yml
enable_experimental: true
namespaces:
  - name: app1
    format: '$remote_addr - [$time_local] "$request" $status $body_bytes_sent rt=$request_time'
    histogram_buckets: [.1, .3, 1.5, 10.5]
    source:
      files:
        - /mnt/nginxlogs/access.log
```

### Docker

Run the exporter as follows (adjust paths like /path/to/logs and /path/to/config to your own needs):

```
docker run \
    --name nginx-exporter \
    -p 4040:4040 \
    -v /path/to/logs:/mnt/nginxlogs \
    -v /path/to/config.yml:/etc/prometheus-nginxlog-exporter.yml \
    quay.io/martinhelmich/prometheus-nginxlog-exporter \
    -config-file /etc/prometheus-nginxlog-exporter.yml
```

### Docker-Compose

There is an example of a docker-compose file with a nginx and a nginx-monitor:

```yml
version: '3.8'
services:
  web:
    image: nginx
    volumes:
      - ./templates:/var/log/nginx
      - ./nginx.conf:/etc/nginx/nginx.conf
    ports:
      - '8080:80'
  nginx-exporter:
    image: valencakarine/nginx-monitor
    depends_on:
      - web
    ports:
      - '4040:4040'
    volumes:
      - ./templates:/mnt/nginxlogs
      - ./config.yml:/etc/prometheus-nginxlog-exporter.yml
    command: -config-file /etc/prometheus-nginxlog-exporter.yml
```

### Kubernetes

To Do

## Big Brother

This is part of a more large application called Big Brother.

## Credits

- [NGINX-to-Prometheus log file exporter](https://github.com/martin-helmich/prometheus-nginxlog-exporter)
- [tail](https://github.com/hpcloud/tail), MIT license
- [gonx](https://github.com/satyrius/gonx), MIT license
- [Prometheus Go client library](https://github.com/prometheus/client_golang), Apache License
- [HashiCorp configuration language](https://github.com/hashicorp/hcl), Mozilla Public License
