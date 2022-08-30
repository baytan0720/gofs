package api

import "testing"

func TestGet(t *testing.T) {
	count := 0
	for count < 999 {
		count++
		Put("/", "./test.txt", "127.0.0.1:1234")
	}
}
