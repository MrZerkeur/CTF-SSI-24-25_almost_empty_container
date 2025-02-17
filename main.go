package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

// randomString returns a random string of length n.
func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	// Seed the random generator.
	rand.Seed(time.Now().UnixNano())

	// Since scratch has an empty filesystem, we'll create a directory in "/".
	randomDir := filepath.Join("/", randomString(10))

	// Create the directory.
	if err := os.Mkdir(randomDir, 0755); err != nil {
		panic(err)
	}

	// Define the full path for "FLAG.txt" within the random directory.
	flagPath := filepath.Join(randomDir, "FLAG.txt")
	flagContent := "RETRO{N0_B1NARY}"

	// Write the file.
	if err := ioutil.WriteFile(flagPath, []byte(flagContent), 0644); err != nil {
		panic(err)
	}

	// Sleep to keep the container running.
	time.Sleep(1000000000 * time.Second)
}
