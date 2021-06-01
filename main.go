/*
 * forum
 *
 * Тестовое задание для реализации проекта \"Форумы\" на курсе по базам данных в Технопарке Mail.ru (https://park.mail.ru).
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"log"
	"net/http"

	sw "github.com/Snikimonkd/dataBases/internal"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":5000", router))
}
