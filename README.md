# NATS Simple Pub-Sub Mechanism
This is a basic pub-sub project using NATS. There is one publisher who publishes a "Hello World" message to a subject called "foo". There are two subscribers who subscribes to "foo" and receives the message as soon as it is published.

## Docker Installation of NATS
[Installing via Docker](https://docs.nats.io/running-a-nats-service/introduction/installation#installing-via-docker)

## Steps to Run
### Step 1: Run subscriber1
```go run subscriber1/main.go```

### Step 2: Run subscriber2
```go run subscriber2/main.go```

### Step 3: Run publisher
```go run publisher/main.go```

## Resources
[NATS - Go Client](https://github.com/nats-io/nats.go)