how we can implement a kafka consumer-producer theme 
in a more kind of production like way

in order to illustrate this we should write a program that :
- simply gets a message from topic_one
- put it to the topic_two

requirements of this project :
- basic understanding on how apache kafka works
    - what are producers, consumers and so forth
- basic concurrency in golang
    - how go-routines works, channels, select statements

1. install dependency 
    1.1 sigmento.kafka-go package
    go get github.com/segmentio/kafka-go

2. create reader
    2.1 implement struct "Reader"
    2.2 create cunstructor "NewKafkaReader" that returns *Reader which contains *kafkago.reader
        2.2.1 register your brocker lokalhost
        2.2.2 topic
        2.2.3 groupid
    
3. define the method "FetchMessage" 
    3.1 make for loop -- reader has fetchmessage method
    , and whenever reader get a message :
        3.1.1 check wheter our context has expired or not "case <- ctx.Done(): >
            3.1.1.1 if it is (expired), then return error messages
            3.1.1.2 if its not, send the message to the "messages" channel

4. create writer (to write message to second topic)
    2.1 implement struct "Writer"
    2.2 create cunstructor "NewKafkaWriter" that returns *Reader which contains *kafkago.reader
        2.2.1 register your brocker lokalhost
        2.2.2 topic
        2.2.3 groupid

5. define the method "WriteMessages" for the struct that takes our message  
    5.1   make for loop with select statements
        5.1.1 check wheter our context has expired or not "case <- ctx.Done(): >
            5.1.1.1 if it is (expired), then return error messages
            5.1.1.2 if its not, get the message from the "messages" channel
                - after getting it, use the writer method WriteMessages and put the value
    5.2 after write a message, you have to put it into CommitChannel
        - why should we use commit? and how it works?
            - after you commit a messages, it wont be deleted after it was sents
            - despite fact that we send our messages, this massages will be stil present
            - so how does kafka understand whether we should send/read the messages again or not? by using commited messages
                - if you restart the program, those commited messages will be automatically sent again? why?
                    - because we haven't commited those messages, so you have to commit the messages if you dont want them to resended

6. in main.go
    6.1 we have to write three go routines
        - FetchMessage
        - WriteMessage
        - CommitMessage
 
    6.2 we have to write two channels
        - messages
        - messagesCommitChan
            - also we have to have a buffer channels, since sometimes we can put a message to the channel or the other part of go routine wont be able to read it immediately, if we do something like that, there will be deadlock.
           
7. how to run the project
    7.1 run main.go
    7.2 create topic1 and topic2
        ~/kafka/bin/kafka-topics.sh --create --topic topic1 --bootstrap-server localhost:9092
        ~/kafka/bin/kafka-topics.sh --create --topic topic2 --bootstrap-server localhost:9092
    7.3 write the message/event on topic1 as a producer 
        ~/kafka/bin/kafka-console-producer.sh --topic topic1 --bootstrap-server localhost:9092
            - those messages will be automatically sent to topic2 (as a consumer) via reader-writer you just wrote
    7.4 read the event/consume the message on topic2 as consumer
        ~/kafka/bin/kafka-console-consumer.sh --topic topic2 --from-beginning --bootstrap-server localhost:9092
            