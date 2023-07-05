# opentelemetry-example





## jaeger安装
docker run -d -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 16686:16686 -p 14268:14268  -p 14269:14269   -p 9411:9411 -p 6831:6831/udp jaegertracing/all-in-one:latest

## prometheus grafana安装
https://www.prometheus.wang/quickstart/install-prometheus-server.html