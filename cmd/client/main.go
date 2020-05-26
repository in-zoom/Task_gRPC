package main

import (
	"Tes_gRPC/pkg/api"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/grpc"
)

const (
	address         = "localhost:8080"
	defaultFilename = "json.json"
)

//Функция парсит переданный фаил
func parseFile(file string) (*api.AddRequest, error) {
	var consignment *api.AddRequest
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, nil
}

func main() {
	// Установим соединение с сервером
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не могу подключиться: %v", err)
	}
	defer conn.Close()
	client := api.NewAdderClient(conn)

	// Передадим в обработку consignment.json,
	// иначе возьмём путь к файлу из аргументов командной строки
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Не возможно распарсить фаил: %v", err)
	}

	r, err := client.Add(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Не удалось создать: %v", err)
	}
	log.Printf("Создан: %t", r.Result)

}
