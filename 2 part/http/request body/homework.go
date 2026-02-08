package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

var (
	mu       sync.Mutex
	messages = make(map[string]string) // id -> text
	order    []string                  // чтобы красиво выводить в порядке добавления
)

func newID(n int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		// очень редкий случай: если rand сломался, сделаем "запасной" id
		return "fallbackid"
	}
	for i := 0; i < n; i++ {
		b[i] = alphabet[int(b[i])%len(alphabet)]
	}
	return string(b)
}

func printStateLocked() {
	fmt.Println("Current messages:")
	if len(order) == 0 {
		fmt.Println("  (empty)")
		return
	}
	for _, id := range order {
		text, ok := messages[id]
		if !ok {
			continue
		}
		fmt.Printf("  %s: %s\n", id, text)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("SAVE: request body not read:", err)
		_, _ = w.Write([]byte("error: request body not read\n"))
		return
	}

	text := strings.TrimSpace(string(body))
	if text == "" {
		_, _ = w.Write([]byte("error: empty message\n"))
		return
	}

	id := newID(8)

	mu.Lock()
	messages[id] = text
	order = append(order, id)

	fmt.Printf("ADDED: %s: %s\n", id, text)
	printStateLocked()
	mu.Unlock()

	// клиенту полезно вернуть id, чтобы потом удалить
	_, _ = w.Write([]byte(id))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("DELETE: request body not read:", err)
		_, _ = w.Write([]byte("error: request body not read\n"))
		return
	}

	id := strings.TrimSpace(string(body))
	if id == "" {
		_, _ = w.Write([]byte("error: empty id\n"))
		return
	}

	mu.Lock()
	text, ok := messages[id]
	if ok {
		delete(messages, id)

		// убираем id из order
		newOrder := make([]string, 0, len(order))
		for _, x := range order {
			if x != id {
				newOrder = append(newOrder, x)
			}
		}
		order = newOrder

		fmt.Printf("DELETED: %s: %s\n", id, text)
		printStateLocked()
		mu.Unlock()

		_, _ = w.Write([]byte("deleted\n"))
		return
	}

	fmt.Printf("DELETE FAILED: id=%s (not found)\n", id)
	printStateLocked()
	mu.Unlock()

	_, _ = w.Write([]byte("not found\n"))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	if len(order) == 0 {
		mu.Unlock()
		_, _ = w.Write([]byte("(empty)\n"))
		return
	}

	var sb strings.Builder
	for _, id := range order {
		text, ok := messages[id]
		if !ok {
			continue
		}
		sb.WriteString(id)
		sb.WriteString(": ")
		sb.WriteString(text)
		sb.WriteString("\n")
	}
	mu.Unlock()

	_, _ = w.Write([]byte(sb.String()))
}

func main() {
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/get", getHandler)

	fmt.Println("Listening on :9091")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Server is not started:", err)
	}
}
