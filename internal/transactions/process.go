package transactions

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gatimugabriel/1m-transactions/pkg/utils"
)

func ProcessTransactions() {
	defer utils.FunctionTimer("ProcessTransactions")()

	file, err := os.Open("data.jsonl")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var tx Transaction
		if err := json.Unmarshal([]byte(scanner.Text()), &tx); err != nil {
			continue
		}
		err := updateDatabase(tx)
		if err != nil {
			fmt.Println("DB insert error:", err)
			return
		}
	}
}
