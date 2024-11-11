package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	inputFile, err := os.Open("structs.txt")
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)

	outputFile, err := os.OpenFile("wallets.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer outputFile.Close()
	writer := bufio.NewWriter(outputFile)

	matches := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		matches = append(matches, line)
	}

	for true {
		w := generateEthWallet()
		if checkNiceWallet(w.Address, matches) {
			wallet := fmt.Sprintf("%s %s\n", w.Address, w.Secret)
			_, err := writer.WriteString(wallet)
			if err != nil {
				log.Println("Error writing to file:", err)
			} else {
				writer.Flush()
				log.Printf("Generate wallet: %s\n", w.Address)
			}
		}
		time.Sleep(time.Millisecond)
	}
}

func generateEthWallet() Wallet {
	key, _ := crypto.GenerateKey()
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	privateKey := hex.EncodeToString(key.D.Bytes())
	return Wallet{address, privateKey}
}

func checkNiceWallet(addr string, patterns []string) bool {
	address := strings.ToLower(addr[2:])
	for _, p := range patterns {
		if matchesPattern(address, p) {
			return true
		}
	}
	return false
}

func matchesPattern(input, pattern string) bool {
	pattern = "^" + pattern + "$"
	pattern = strings.ReplaceAll(pattern, "*", ".*")
	pattern = strings.ReplaceAll(pattern, "x", ".")
	matched, _ := regexp.MatchString(pattern, input)
	return matched
}

type Wallet struct {
	Address string `json:"a"`
	Secret  string `json:"s"`
}
