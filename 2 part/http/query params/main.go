// https://smth:443/header?v=123&t=10s
// квери параметр : ключ - значение
// ? - говорит о том что пойдут дальше квери параметры
// & - разделяет один квери параметр от другого
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fooParam := r.URL.Query().Get("foo")
	booParam := r.URL.Query().Get("boo")

	fmt.Println("foo параметр:", fooParam)
	fmt.Println("boo параметр:", booParam)
}

func main() {
	http.HandleFunc("/default", handler)

	fmt.Println("Запускаю HTTP сервер!")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка во время работы HTTP сервера:", err)
	}
}
