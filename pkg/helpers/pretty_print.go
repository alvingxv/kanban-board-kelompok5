package helpers

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrettyPrint(entity interface{}) {
	marshaled, err := json.MarshalIndent(entity, "", "   ")
	if err != nil {
		log.Fatalf("marshaling error: %s", err)
	}
	fmt.Println(string(marshaled))
}
