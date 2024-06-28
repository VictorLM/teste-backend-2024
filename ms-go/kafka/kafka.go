package kafka

import (
	"context"
	"fmt"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

// var connection *kafka.Conn

func Test(topic string, message string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"127.0.0.1:9092"},
		Topic:   "your-topic",
		GroupID: "your-group",
	})

	ctx := context.Background()
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.Fatalf("failed to read message: %v", err)
		}
		fmt.Printf("message at offset %d: %s = %s\n", msg.Offset, string(msg.Key), string(msg.Value))
	}

	// w := &kafka.Writer{
	// 	Addr:     kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094"),
	// 	Balancer: &kafka.LeastBytes{},
	// }

	// err := w.WriteMessages(context.Background(),
	// 	kafka.Message{
	// 		Topic: topic,
	// 		Key:   []byte("Key-A"),
	// 		Value: []byte(message),
	// 	},
	// )
	// if err != nil {
	// 	log.Fatal("failed to write messages:", err)
	// }

	// if err := w.Close(); err != nil {
	// 	log.Fatal("failed to close writer:", err)
	// }
}

// func connect(topic string) {
// 	if connection == nil {
// 		partition := 0
// 		var errConn error

// 		connection, errConn = kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

// 		if errConn != nil {
// 			log.Fatal("failed to dial leader:", errConn)
// 		}
// 	}
// }

// func ProduceMessage(topic string, message string) {
// 	connect(topic)

// 	connection.SetWriteDeadline(time.Now().Add(10 * time.Second))

// 	_, err := connection.WriteMessages(
// 		kafka.Message{Value: []byte(message)},
// 	)

// 	if err != nil {
// 		log.Fatal("failed to write messages:", err)
// 	}
// }

// func ConsumeMessages(topic string) {
// 	connect(topic)

// 	connection.SetReadDeadline(time.Now().Add(10 * time.Second))
// 	batch := connection.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

// 	b := make([]byte, 10e3) // 10KB max per message
// 	for {
// 		n, err := batch.Read(b)
// 		if err != nil {
// 			break
// 		}
// 		fmt.Println(string(b[:n]))
// 	}

// 	if err := batch.Close(); err != nil {
// 		log.Fatal("failed to close batch:", err)
// 	}
// }

// func CloseConnection() {
// 	if connection != nil {
// 		if err := connection.Close(); err != nil {
// 			log.Fatal("failed to close writer:", err)
// 		}
// 	}
// }
