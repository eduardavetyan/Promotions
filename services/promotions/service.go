package promotions

import (
	"bufio"
	"context"
	"log"
	"os"
	"promotions-app/db"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct{}

func (s Service) Find(id string) (Promotion, bool) {
	db := db.GetDB(PromotionCollection)

	var promotion Promotion

	filter := bson.D{{Key: "_id", Value: id}}
	err := db.FindOne(context.TODO(), filter).Decode(&promotion)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return promotion, false
		}
		log.Panic(err)
	}

	return promotion, true
}

func (s Service) Import(filePath string) bool {
	rowsChannel := make(chan [][]string)

	go readCsvByChunks(filePath, rowsChannel)
	go importCsvChunks(rowsChannel)

	return true
}

func readCsvByChunks(filePath string, rowsChannel chan [][]string) {
	defer os.Remove(filePath)
	defer close(rowsChannel)

	scanner := getBufioScanner(filePath)

	var rowsChunk [][]string

	for scanner.Scan() {
		row := scanner.Text()
		rowArr := strings.Split(row, ",")
		rowsChunk = append(rowsChunk, rowArr)

		if isChunkSizeReached(rowsChunk) {
			rowsChannel <- rowsChunk
			rowsChunk = [][]string{}
		}
	}

	if len(rowsChunk) > 0 {
		rowsChannel <- rowsChunk
	}
}

func getBufioScanner(filePath string) *bufio.Scanner {
	f, err := os.Open(filePath)
	if err != nil {
		log.Panic(err)
	}
	return bufio.NewScanner(f)
}

func isChunkSizeReached(rowsChunk [][]string) bool {
	const CHUNK_MAX_SIZE = 10000
	return len(rowsChunk) == CHUNK_MAX_SIZE
}

func importCsvChunks(rowsChannel chan [][]string) {
	for {
		rowsChunk, ok := <-rowsChannel
		if !ok {
			break
		}
		insertRowsChunkToDB(rowsChunk)
	}
}

func insertRowsChunkToDB(rowsChunk [][]string) {
	promotions := csvRowsToChunks(rowsChunk)

	if len(promotions) > 0 {
		bulkInsert(promotions)
	}
}

func csvRowsToChunks(rowsChunk [][]string) []interface{} {
	var promotions []interface{}

	for _, row := range rowsChunk {
		promotion, ok := csvRowToModel(row)
		if !ok {
			continue
		}
		promotions = append(promotions, promotion)
	}

	return promotions
}

func csvRowToModel(row []string) (Promotion, bool) {
	if len(row) < 3 {
		return Promotion{}, false
	}

	id := row[0]
	date := row[2]
	price, err := strconv.ParseFloat(row[1], 64)
	if err != nil {
		log.Print(err)
		return Promotion{}, false
	}

	return Promotion{id, price, date}, true
}

func bulkInsert(promotions []interface{}) {
	db := db.GetDB(PromotionCollection)

	_, err := db.InsertMany(context.TODO(), promotions)

	if err != nil {
		log.Print(err)
	}
}
