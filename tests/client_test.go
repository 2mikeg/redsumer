package gorgp

import (
	"fmt"
	"testing"

	"github.com/2mikeg/gorgp/pkg/client"
)

func TestClient(t *testing.T) {

	client := client.RedisClient("localhost", 6379, 0)
	clientTypeString := fmt.Sprint(client)

	if clientTypeString != "Redis<localhost:6379 db:0>" {
		t.Error(clientTypeString)
	}

}
