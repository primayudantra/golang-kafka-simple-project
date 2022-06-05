## Kafka Cluster

- 

## Kafka Broker

- 

## Kafka Topics

- Its a category
- Create, List, Describe, Delete

## Kafka Producer

- Its an Application or piece of codes to write the data to kafka
- Send data to topic

## Kafka Consumer

- Its an Application or piece of codes to read a data from kafka
- Sometimes consumer can be a producer also
- Read data from topic

---
## Commands
```sh
# Running docker-compose
> docker-compose up -d

# Go to inside docker container
> docker exec -it kafka /bin/sh

# CD to kafka bin
> cd /opt/kafka_2.13-2.8.1/bin

# ------------------------------------------------------

# Get Kafka List of Topics
> kafka-topics.sh --list --zookeeper zookeeper:2181

# Create Kafka Topic
> kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic messages_topic

# Describe Kafka Topic
> kafka-topics.sh --describe --zookeper zookeeper:2181 --topic messages_topic

# Delete Kafka Topic
> kafka-topics.sh --delete --zookeeper zookeeper:2181 --topic messages_topic

# Producer Kafka Topic
> kafka-console-producer.sh --broker-list kafka:9092 --topic messages_topic

# Consumer Kafka Topic
> kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic messages
_topic

```