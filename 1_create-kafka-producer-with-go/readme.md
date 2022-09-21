This is my experiment to create kafka producer with Go. 
https://kafka.apache.org/quickstart
- quickstart notes summary from official apache-kafka website -

this is for ubuntu linux 20.04 user via command line :
1. create topic
~/kafka/bin/kafka-topics.sh --create --topic <TopicName> --bootstrap-server localhost:9092

2. view details of created topic
~/kafka/bin/kafka-topics.sh --describe --topic <TopicName> --bootstrap-server localhost:9092

3. write some events into the topic 
~/kafka/bin/kafka-console-producer.sh --topic <TopicName> --bootstrap-server localhost:9092

    3.1. write some events / create producer into the topic with Golang
    read the producer.go source code for details

4. read the events
~/kafka/bin/kafka-console-consumer.sh --topic <TopicName> --from-beginning --bootstrap-server localhost:9092

5. import/export data as streams of event with kafka connect
6. process your events with kafka streams with golang

1
2 
3 ~/kafka/bin/kafka-console-producer.sh --topic ContohTopic --bootstrap-server localhost:9092
4 ~/kafka/bin/kafka-console-consumer.sh --topic ContohTopic --from-beginning --bootstrap-server localhost:9092 --partition 0
