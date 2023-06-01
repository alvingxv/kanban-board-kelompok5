package helpers

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrettyPrint(entity interface{}, messages ...string) {

	if len(messages) > 0 {
		for _, message := range messages {
			fmt.Println(message)
		}
	}
	marshaled, err := json.MarshalIndent(entity, "", "   ")
	if err != nil {
		log.Fatalf("marshaling error: %s", err)
	}
	fmt.Println(string(marshaled) + "\n")
}
