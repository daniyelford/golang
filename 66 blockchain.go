package main

import (
	"crypto/sha256" //sha256 hash algorithm
	"encoding/hex"  //converting data to hex
	"fmt"
	"math/big"  //big numbers
	"math/rand" //random numbers
	"strconv"   //converting string to various data types and vice versa
	"time"      //time ,date
)

type Transaction struct {
	Sender    string
	Receiver  string
	Amount    float64
	Timestamp string
	Signature string
}
type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	PreviousHash string
	Hash         string
	Nonce        int
	Difficulty   int
}

type Blockchain struct {
	Blocks []Block
}

//previoushash+transactions+index+timestamp+nonce

func calculateHash(block Block) string {
	data := strconv.Itoa(block.Index) + block.Timestamp + fmt.Sprintf("%v", block.Transactions) + block.PreviousHash + strconv.Itoa(block.Nonce)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func validateTransaction(tx Transaction, balance map[string]float64) bool {
	if balance[tx.Sender] >= tx.Amount {
		balance[tx.Sender] -= tx.Amount
		balance[tx.Receiver] += tx.Amount
		return true
	}
	return false
}
func mineBlock(previousBlock Block, transactions []Transaction, difficulty int, balance map[string]float64) Block {
	var newBlock Block
	newBlock.Index = previousBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Transactions = transactions
	newBlock.PreviousHash = previousBlock.Hash
	newBlock.Difficulty = difficulty
	for _, tx := range transactions {
		if !validateTransaction(tx, balance) {
			fmt.Println("Transaction is not valid")
			return Block{}
		}
	}
	fmt.Println("mining started")

	if newBlock.Index%5 == 0 {
		newBlock.Difficulty++
	}
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))

	for {
		newBlock.Nonce = rand.Intn(1000000)
		newBlock.Hash = calculateHash(newBlock)
		hashint := new(big.Int)
		hashint.SetString(newBlock.Hash, 16)
		if hashint.Cmp(target) == -1 {
			fmt.Println("block minded", newBlock.Hash)
			break
		}
	}
	return newBlock
}
func generateGenesisBlock() Block {
	return Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Transactions: []Transaction{{Sender: "abc", Receiver: "def", Amount: 0, Timestamp: time.Now().String()}},
		PreviousHash: "0",
		Nonce:        0,
		Difficulty:   4,
		Hash:         " ",
	}
}

type BlockchainNetwork struct {
	Nodes []Blockchain
}

func (network *BlockchainNetwork) CreateNode() Blockchain {
	genesisBlock := generateGenesisBlock()
	genesisBlock.Hash = calculateHash(genesisBlock)
	newNode := Blockchain{}
	newNode.Blocks = append(newNode.Blocks, genesisBlock)
	network.Nodes = append(network.Nodes, newNode)
	fmt.Println("new node created in blockchain network")
	return newNode
}
func (network *BlockchainNetwork) SyncBlock() {
	if len(network.Nodes) > 1 {
		latestBlocks := make(map[int]Block)
		for i := range network.Nodes {
			latestBlocks[i] = network.Nodes[i].Blocks[len(network.Nodes[i].Blocks)-1]
		}
		for i := range network.Nodes {
			network.Nodes[i].Blocks = append(network.Nodes[i].Blocks, latestBlocks[i])
		}
	}
}
func (network *BlockchainNetwork) BroadcastBlock(newBlock Block) {
	for i := 0; i < len(network.Nodes); i++ {
		network.Nodes[i].Blocks = append(network.Nodes[i].Blocks, newBlock)
	}
}
func main() {

	balance := map[string]float64{
		"client1": 1000,
		"client2": 500,
	}

	network := BlockchainNetwork{}
	network.CreateNode()
	network.CreateNode()
	network.CreateNode()
	network.CreateNode()
	network.CreateNode()

	for i := 1; i <= 7; i++ {
		transactions := []Transaction{
			{Sender: "client1", Receiver: "client2", Amount: float64(rand.Intn(100)), Timestamp: time.Now().String()},
		}

		if len(network.Nodes[0].Blocks) > 0 {
			newBlock := mineBlock(network.Nodes[0].Blocks[len(network.Nodes[0].Blocks)-1], transactions, 4, balance)
			if newBlock.Hash != "" {
				network.BroadcastBlock(newBlock)
				fmt.Printf("Block #%d added to the blockchain\n", newBlock.Index)
			}
		}
	}
	network.SyncBlock()
	for i, node := range network.Nodes {
		fmt.Printf("\nNode %d Blockchain:\n", i+1)
		for _, block := range node.Blocks {
			fmt.Printf("Block %d: Hash: %s\n", block.Index, block.Hash)
		}
	}
}
