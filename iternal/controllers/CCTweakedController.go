package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/kaptinlin/jsonrepair"
)

type Message struct {
	Content  string `json:"content"`
	Msg_type int    `json:"msg_type"`
}

var message_queue = []Message{}

func preprocess_message(r *http.Request) []byte {
	message_bytes, _ := io.ReadAll(r.Body)
	message_body := string(message_bytes)

	repaired, err := jsonrepair.JSONRepair(message_body)
	if err != nil {
		log.Printf("Failed to fix JSON: %v", err)
	}
	repaired_bytes := []byte(repaired)

	return repaired_bytes
}

func decodeMessage(r *http.Request) (Message, error) {
	var messageRC Message

	processed_message := preprocess_message(r)

	err := json.Unmarshal(processed_message, &messageRC)
	if err != nil {
		log.Println(err)
		return Message{}, err
	}
	return messageRC, nil
}

func clearQueue() {
	if len(message_queue) <= 1 {
		return
	}
	message_queue = message_queue[1:]
}

func PostSendMessage(w http.ResponseWriter, r *http.Request) {

	decoded_message, err := decodeMessage(r)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Failed to decode message %v", err)
		log.Printf("Failed to decode message %v", err)
		return
	}
	message_queue = append(message_queue, decoded_message)

	clearQueue()

}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	if len(message_queue) == 0 {
		message_queue = append(message_queue, Message{"no message accepted yet", 404})
	}
	fmt.Fprintln(w, message_queue[0])
}
