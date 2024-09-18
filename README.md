# go-stream

**go-stream** is a Go (Golang) library designed for real-time stream data processing. It provides a flexible framework that allows developers to build data processing pipelines using operations like map, filter, and windowing. With its modular architecture, go-stream can be integrated with messaging systems like Google Pub/Sub, making it suitable for various real-time data processing use cases.

## Features

- **Flexible Stream Processing**: Create custom pipelines with map, filter, and window operations.
- **Extensible Operators**: Define processing functions tailored to specific needs.
- **Windowing Support**: Implement tumbling and sliding windows for data aggregation based on time intervals.

## Getting Started

### Installation

```bash
go get github.com/renatosaksanni/go-stream
```

### Example Usage

```go
package main

import (
    "fmt"
    "go-stream/internal/adapters/input"
    "go-stream/internal/adapters/output"
    "go-stream/internal/adapters/processing"
    "go-stream/internal/application"
    "go-stream/internal/infrastructure/pubsub"
    "go-stream/internal/ports"
    "time"
)

func main() {
    projectID := "your-gcp-project-id"
    subscriptionID := "your-pubsub-subscription-id"
    topicID := "your-pubsub-topic-id"

    // Initialize Pub/Sub client
    pubSubClient, err := pubsub.NewPubSubClient(projectID, subscriptionID, topicID)
    if err != nil {
        fmt.Printf("Error creating PubSub client: %v\n", err)
        return
    }

    // Create input and output ports
    consumer := input.NewPubSubConsumer(pubSubClient)
    producer := output.NewPubSubProducer(pubSubClient)

    // Define custom processing functions
    mapFunc := func(data interface{}) (interface{}, error) {
        // User-defined processing logic
        return data, nil
    }

    filterFunc := func(data interface{}) bool {
        // User-defined filter logic
        return true
    }

    aggregateFunc := func(data []interface{}) (interface{}, error) {
        // User-defined aggregation logic
        return data, nil
    }

    // Create operators
    mapOperator := processing.NewMapOperator(mapFunc)
    filterOperator := processing.NewFilterOperator(filterFunc)
    windowOperator := processing.NewTumblingWindowOperator(1*time.Minute, aggregateFunc)

    // Create service
    operators := []ports.Operator{mapOperator, filterOperator}
    service := application.NewStreamService(consumer, producer, operators, windowOperator)

    // Run service
    err = service.Start()
    if err != nil {
        fmt.Printf("Error running service: %v\n", err)
    }
}
```

## Documentation

For more information on how to use go-stream, please refer to the [documentation](#).

## Contributions

Contributions are welcome. Please read the [contribution guidelines](#) before submitting a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.