package main

/*
strconv مخفف String Convert هست و برای تبدیل رشته (string) به نوع‌های عددی/بولین و برعکس استفاده میشه.
خیلی وقت‌ها داده‌ها به صورت string میاد (مثلاً از فایل، شبکه، JSON، ورودی کاربر) ولی ما باید اون‌ها رو به int, float, bool تبدیل کنیم تا بشه روی اون‌ها عملیات انجام داد. برعکسش هم لازمه → وقتی می‌خوایم داده‌ی عددی رو به کاربر نمایش بدیم یا داخل رشته ذخیره کنیم باید اون رو به string تبدیل کنیم
*/
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
type BlockchainNetwork struct {
	Nodes []Blockchain
}

// previoushash+transactions+index+timestamp+nonce
func calculateHash(block Block) string {
	// Itoa = Integer to ASCII
	// عدد صحیح int رو به رشته تبدیل می‌کنه
	data := strconv.Itoa(block.Index) + block.Timestamp + fmt.Sprintf("%v", block.Transactions) + block.PreviousHash + strconv.Itoa(block.Nonce)

	// Atoi = ASCII to Integer
	// رشته رو به int تبدیل می‌کنه
	// n, err := strconv.Atoi("456")

	// base: مبنا (۲، ۱۰، ۱۶ و ...)
	// bitSize: سایز (0, 8, 16, 32, 64)
	// strconv.ParseInt(s string, base int, bitSize int) (int64, error)
	// n, _ := strconv.ParseInt("ff", 16, 64)
	// fmt.Println(n)   // 255

	// به اعشاری
	// strconv.ParseFloat(s string, bitSize int) (float64, error)
	// f, _ := strconv.ParseFloat("3.14", 64)
	// fmt.Println(f)   // 3.14

	// strconv.ParseBool("true")  // true
	// strconv.ParseBool("1")     // true
	// strconv.ParseBool("false") // false
	// strconv.ParseBool("0")     // false

	/*
		برعکس: تبدیل انواع به رشته
		strconv.FormatInt(i int64, base int) → int64 به string با مبنا دلخواه.
		strconv.FormatFloat(f float64, fmt byte, prec, bitSize int) → float به string.
		strconv.FormatBool(b bool) → "true" یا "false".
		مثال:
		fmt.Println(strconv.FormatInt(255, 16))   // "ff"
		fmt.Println(strconv.FormatBool(true))     // "true"
	*/

	// الگوریتم SHA-256 روی داده اجرا میشه
	// خروجی یک [32]byte هست (۳۲ بایت ثابت → ۲۵۶ بیت).
	hash := sha256.Sum256([]byte(data))
	// hash[:] → تبدیل [32]byte به slice از بایت‌ها.
	// hex.EncodeToString → بایت‌ها رو به یک رشته‌ی hex (پایه ۱۶) تبدیل می‌کنه
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
	// یک متغیر از نوع Block ساخته میشه. این بلاک جدید هست که می‌خوایم استخراج (mine) کنیم
	var newBlock Block
	// شماره بلاک رو مشخص می‌کنیم.هر بلاک شماره (Index) داره و بلاک جدید یکی بیشتر از بلاک قبلیه
	newBlock.Index = previousBlock.Index + 1
	// زمان فعلی سیستم گرفته میشه و به رشته تبدیل میشه، برای ذخیره‌ی زمان ایجاد بلاک
	newBlock.Timestamp = time.Now().String()
	// تراکنش‌هایی که توی این بلاک باید ذخیره بشن به بلاک جدید اضافه میشن
	newBlock.Transactions = transactions
	// هش بلاک قبلی ذخیره میشه. این باعث میشه بلاک‌ها به هم زنجیر بشن (Blockchain).
	newBlock.PreviousHash = previousBlock.Hash
	// سختی استخراج بلاک رو تنظیم می‌کنیم. (این عدد نشون میده پیدا کردن بلاک جدید چقدر سخت باشه.)
	newBlock.Difficulty = difficulty
	// برای هر تراکنش (tx) بررسی می‌کنه که معتبر هست یا نه (با validateTransaction).
	// اگر حتی یکی معتبر نباشه، پیام چاپ میشه و یک بلاک خالی برمی‌گرده.
	// ⛔ این کار جلوی اضافه شدن تراکنش‌های تقلبی به بلاک رو می‌گیره.
	for _, tx := range transactions {
		if !validateTransaction(tx, balance) {
			fmt.Println("Transaction is not valid")
			return Block{}
		}
	}
	fmt.Println("mining started")
	// هر ۵ بلاک یکبار سختی ماینینگ رو افزایش میده (تا استخراج سخت‌تر بشه و شبکه پایدار بمونه)
	if newBlock.Index%5 == 0 {
		newBlock.Difficulty++
	}
	// 	اینجا هدف (Target) ساخته میشه.
	// یعنی عددی که هش بلاک باید از اون کوچیک‌تر باشه تا بلاک معتبر باشه.
	// Lsh یعنی Left Shift → عدد ۱ رو می‌بره به سمت چپ (یعنی ضرب در ۲^(256-difficulty)).
	// هرچی سختی بیشتر باشه، Target کوچیک‌تر میشه → پیدا کردن هش معتبر سخت‌تر میشه
	// یک عدد صحیح بزرگ (Big Integer) با مقدار 1 درست می‌کنه
	target := big.NewInt(1)
	// target.Lsh(x, n)
	// یعنی Left Shift → عدد x رو n بیت به چپ می‌بره.
	// در واقع x << n (اما برای big.Int).
	// یعنی عدد 1 ضربدر 2^n میشه.
	// 1 << 0 = 1 → یعنی همون 1.
	// 1 << 1 = 2 → یعنی 2 به توان 1.
	// 1 << 2 = 4 → یعنی 2 به توان 2.
	// 1 << 8 = 256 → یعنی 2 به توان 8.
	target.Lsh(target, uint(256-difficulty))

	for {
		newBlock.Nonce = rand.Intn(1000000)
		newBlock.Hash = calculateHash(newBlock)
		// هش به صورت یک رشته هگزادسیمال هست (base 16). اینجا اون رو به عدد (big.Int) تبدیل می‌کنیم تا بشه مقایسه کرد
		hashint := new(big.Int)
		hashint.SetString(newBlock.Hash, 16)
		// Cmp دو عدد رو مقایسه می‌کنه
		// -1 یعنی hashint < target
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
