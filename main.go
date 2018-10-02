package main

import (
	"log"      // пакет для логирования
	"net/http" // пакет для поддержки HTTP протокола
	"simple_HTTP_server/handlers"
)



func main() {
	http.HandleFunc("/date", handlers.Date) // установим роутер
	http.HandleFunc("/wait", handlers.Wait)
	err := http.ListenAndServe(":9000", nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
