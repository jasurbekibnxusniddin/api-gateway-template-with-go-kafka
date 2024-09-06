# api-gateway-template-with-go-kafka
go kafka api

# running kafka broker
    docker-compose up -d

# creating topic
- step 1
```
    docker exec -it <kafka-container-id> /bin/bash
```
- to view container id:
        docker ps 

- step 2
```
    /opt/bitnami/kafka/bin/kafka-topics.sh --create --topic your-topic-name --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
```

- view topics list:
```
    /opt/bitnami/kafka/bin/kafka-topics.sh --list --bootstrap-server localhost:9092
```

 /opt/bitnami/kafka/bin/kafka-topics.sh --create --topic update_todo --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1