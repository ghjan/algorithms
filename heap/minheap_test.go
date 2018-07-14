package heap

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	H := Create(5)
	assert.Equal(t, 0, H.Size)
	assert.Equal(t, MINH, H.Keys[0])
}

func TestHeap_Insert(t *testing.T) {
	keysString := "46 23 26 24 10"
	keys := strings.Split(keysString, " ")
	H := Create(len(keys))
	for _, key := range keys {
		if k, err := strconv.Atoi(key); err != nil {
			fmt.Println(err)
		} else {
			if err := H.Insert(k); err != nil {
				fmt.Println(err)
			}
		}
	}
	assert.Equal(t, len(keys), H.Size)
	for _, value := range H.Keys {
		fmt.Printf(" %d", value)
	}
}

func TestHeap_Path(t *testing.T) {
	keysString := "46 23 26 24 10"
	positionsString := "5 4 3"
	keys := strings.Split(keysString, " ")
	H := Create(len(keys))
	for _, key := range keys {
		k, _ := strconv.Atoi(key)
		H.Insert(k)
	}
	assert.Equal(t, len(keys), H.Size)

	positions := strings.Split(positionsString, " ")
	for _, search := range positions {
		X, _ := strconv.Atoi(search)
		fmt.Println(H.Path(X))
	}

}
