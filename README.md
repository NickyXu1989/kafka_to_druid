# kafka_to_druid
auto config druid supervisor and transfer prometheus data in kafka to druid
written in golang


env settings:
KAFKA_BROKER_LIST=127.0.0.1:9092
KAFKA_TOPIC=prometheus_metrics
KAFKA_LIMIT_PER_SEC=1000
REDIS_PASSWORD=123456
CLUSTER_MODE=false
DRUID_API_SERVER=http://127.0.0.1:8888


support limit control per second
