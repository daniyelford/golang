package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// تابع هش (sha256) فقط با داده باینری کار می‌کنه، نه با رشته مستقیم
// تعریف بلاک
type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

// نود شبکه
type Node struct {
	Name   string
	Blocks []Block
	// یعنی کانالی که می‌تونه مقادیر از نوع Block رو حمل کنه
	Inbox chan Block // کانالی برای دریافت بلاک جدید
}

// محاسبه هش بلاک
func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.Data, block.PrevHash)
	/*
		sha256.New(): تابع هش SHA256
		h.Write(...): داده رو به هش می‌دهیم
		h.Sum(nil): نتیجه هش (به صورت []byte)
	*/
	h := sha256.New()
	h.Write([]byte(record))
	// hex.EncodeToString(...): تبدیل هش به رشته‌ی هگز (قابل خواندن)
	return hex.EncodeToString(h.Sum(nil))
}

// ساخت بلاک جدید
func generateBlock(oldBlock Block, data string) Block {
	newBlock := Block{
		Index:     oldBlock.Index + 1,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Data:      data,
		PrevHash:  oldBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

// شبکه بلاکچین
type BlockchainNetwork struct {
	Nodes []*Node
}

// پخش بلاک جدید به همه نودها
func (network *BlockchainNetwork) Broadcast(block Block) {
	for _, node := range network.Nodes {
		node.Inbox <- block
	}
}

// شروع گوش دادن نود
func (n *Node) Start() {
	go func() {
		for {
			block := <-n.Inbox
			// بررسی اعتبار بلاک
			lastBlock := n.Blocks[len(n.Blocks)-1]
			if block.Index == lastBlock.Index+1 && block.PrevHash == lastBlock.Hash && block.Hash == calculateHash(block) {
				n.Blocks = append(n.Blocks, block)
				fmt.Printf("[%s] بلاک %d اضافه شد: %s\n", n.Name, block.Index, block.Data)
			} else {
				fmt.Printf("[%s] بلاک نامعتبر دریافت شد ❌\n", n.Name)
			}
		}
	}()
}

func main() {
	// بلاک اولیه (Genesis Block)
	genesis := Block{
		Index:     0,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Data:      "Genesis Block",
		PrevHash:  "",
	}
	genesis.Hash = calculateHash(genesis)

	// ایجاد نودها
	node1 := &Node{Name: "Node1", Blocks: []Block{genesis}, Inbox: make(chan Block, 10)}
	node2 := &Node{Name: "Node2", Blocks: []Block{genesis}, Inbox: make(chan Block, 10)}
	node3 := &Node{Name: "Node3", Blocks: []Block{genesis}, Inbox: make(chan Block, 10)}

	network := &BlockchainNetwork{Nodes: []*Node{node1, node2, node3}}

	// شروع نودها
	node1.Start()
	node2.Start()
	node3.Start()

	// نود ۱ بلاک جدید می‌سازه
	newBlock1 := generateBlock(node1.Blocks[len(node1.Blocks)-1], "First block from Node1")
	network.Broadcast(newBlock1)

	time.Sleep(1 * time.Second)

	// نود ۲ بلاک جدید می‌سازه
	newBlock2 := generateBlock(node2.Blocks[len(node2.Blocks)-1], "Block from Node2")
	network.Broadcast(newBlock2)

	// newBlock3 := generateBlock(node3.Blocks[len(node3.Blocks)-1], "Block from Node3")
	// network.Broadcast(newBlock3)

	time.Sleep(1 * time.Second)

	// نمایش تعداد بلاک‌های هر نود
	fmt.Println("\n📌 وضعیت نهایی بلاکچین در هر نود:")
	for _, node := range network.Nodes {
		fmt.Printf("%s: %d بلاک\n", node.Name, len(node.Blocks))
	}
}
