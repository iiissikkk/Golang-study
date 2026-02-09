package http_server

import (
	"errors"
	"fmt"
	"net/http"
)

func StartHttpServer() error {
	http.HandleFunc("/ping", func(w http.ResponseWriter, e *http.Request) {
		fmt.Println("Send smth /ping")
		w.Write([]byte("Hello from Docker!\n"))
	})

	err := http.ListenAndServe(":5050", nil)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	} else {
		return err
	}
}
