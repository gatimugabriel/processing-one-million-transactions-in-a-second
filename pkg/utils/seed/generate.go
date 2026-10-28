package seed

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gatimugabriel/1m-transactions/internal/transactions"
	"github.com/gatimugabriel/1m-transactions/pkg/utils"
	"github.com/gatimugabriel/1m-transactions/pkg/utils/files"
)

func GenerateData() {
	defer utils.FunctionTimer("Seed GenerateData")()

	dataFile, err := files.CreateFile("data.jsonl")
	if err != nil {
		fmt.Println("Error creating file: ", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file: ", err)
		}
	}(dataFile)

	productNames := []string{"laptop", "desktop", "mouse", "stylus", "monitor", "speakers"}
	productPrices := []int64{10000, 20000, 30000, 40000, 50000, 60000}

	encoder := json.NewEncoder(dataFile)

	for i := 1; i <= 1000; i++ {
		orderId := i
		pIdx := rand.IntN(len(productNames))
		price := productPrices[pIdx]
		productId := productNames[pIdx] + strconv.Itoa(i)
		userId := generateRandomUnsecureString(10)
		quantity := rand.Int32N(10-1) + 1
		timestamp := time.Now()

		t := transactions.Transaction{
			OrderId:   orderId,
			UserId:    userId,
			ProductId: productId,
			Price:     price,
			Quantity:  quantity,
			Timestamp: timestamp,
		}

		if err := encoder.Encode(t); err != nil {
			log.Fatalf("Failed to encode transaction: %v", err)
		}
	}

}

func generateRandomUnsecureString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	sb.Grow(length)

	for i := 0; i < length; i++ {
		randIndex := rand.IntN(len(charset))
		sb.WriteByte(charset[randIndex])
	}

	return sb.String()
}
