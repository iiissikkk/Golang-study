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
	order    []string                  // порядок добавления
)

func methodNotAllowed(w http.ResponseWriter, r *http.Request, allowed string) {
	w.Header().Set("Allow", allowed)
	w.WriteHeader(http.StatusMethodNotAllowed) // 405
	msg := fmt.Sprintf("method not allowed: %s. use %s\n", r.Method, allowed)
	_, _ = w.Write([]byte(msg))
}

func newID(n int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
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
	if r.Method != http.MethodPost {
		methodNotAllowed(w, r, http.MethodPost)
		return
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("SAVE: request body not read:", err)
		_, _ = w.Write([]byte("error: request body not read\n"))
		return
	}

	text := strings.TrimSpace(string(body))
	if text == "" {
		w.WriteHeader(http.StatusBadRequest)
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

	// вернём id, чтобы клиент мог удалить
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte(id + "\n"))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		methodNotAllowed(w, r, http.MethodDelete)
		return
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("DELETE: request body not read:", err)
		_, _ = w.Write([]byte("error: request body not read\n"))
		return
	}

	id := strings.TrimSpace(string(body))
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
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

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte("deleted\n"))
		return
	}

	fmt.Printf("DELETE FAILED: id=%s (not found)\n", id)
	printStateLocked()
	mu.Unlock()

	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte("error: id not found\n"))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w, r, http.MethodGet)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if len(order) == 0 {
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
