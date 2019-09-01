package main
import "fmt"
import "net/http"
import "io/ioutil"
import "os"
import "time"

type header struct {
	prevBlock string
	merkleRoot string
	merkleRootTest string
	timeStamp string
	numTransactions string
	targetDiff string
	nonce string
}

type transaction struct {
	transVersion string
	numInputs string
	prevTrans [200]string
	transIndex string
	scriptLength string
	numOutputs string
	amountBTC string
	pkScript_length string
	pkScript string
	lockTime string
}

var byteC int
var head header

func main() {

	arg := os.Args[1]

	fmt.Println()
	fmt.Println("-------------------------------")
	fmt.Println("Welcome to Bitcoin Block Viewer")
	fmt.Println("-------------------------------")
	fmt.Println()
	
	// https://blockchain.info/block/000000000000000001f942eb4bfa0aeccb6a14c268f4c72d5fff17270da771b9?format=hex
	
	URL := "https://blockchain.info/block/" + arg + "?format=hex"
	resp, _ := http.Get(URL)
	testBlock, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()	
	
	buildHeader(testBlock)
	displayHeader()

	////////////////start transaction loop////////////////////////////////

	var transactions [25]transaction

	for i:=1; i<=5; i++ {

		transactions[i].transVersion = convertEndian(string(testBlock[byteC:byteC+8]))
		byteC += 8

		//////////////////////start number of inputs loop////////////////////
	 
		transactions[i].numInputs = convertEndian(string(testBlock[byteC:byteC+2]))
		byteC += 2

		for k:=1;k<=fromHex(transactions[i].numInputs);k++ {
			//fmt.Println("***Input ", i, "***")
			transactions[i].prevTrans[k] = convertEndian(string(testBlock[byteC:byteC+64]))
			byteC += 64
			transactions[i].transIndex = convertEndian(string(testBlock[byteC:byteC+8]))
			byteC += 8
			//account for variable langth
			scriptLengthVar := convertEndian(string(testBlock[byteC:byteC+2]))
			//fmt.Println(scriptLengthVar)
    		if scriptLengthVar!="fd" && scriptLengthVar!="fe" && scriptLengthVar!="ff" {
				transactions[i].scriptLength = convertEndian(string(testBlock[byteC:byteC+2]))
				byteC += 2
			} else {
				jump := varLength(scriptLengthVar)
				transactions[i].scriptLength = convertEndian(string(testBlock[byteC+2:byteC+jump]))
				byteC += jump
			}
			byteC += fromHex(transactions[i].scriptLength)*2
			byteC += 8
		}

		//////////////////////end number of inputs loop////////////////////

		transactions[i].numOutputs = convertEndian(string(testBlock[byteC:byteC+2]))
		byteC += 2

		fmt.Println("-------------------------------")
		fmt.Println("Transaction", i, "Inputs")
		fmt.Println("-------------------------------")

		displayTransactionInputs(transactions[i])

		buildOutput(i, transactions[i], testBlock)
	}
	////////////////end transaction loop////////////////////////////////	
}

func buildOutput(i int, transactions transaction, testBlock []byte) {
	for j:=1;j<=fromHex(transactions.numOutputs);j++ {
		transactions.amountBTC = convertEndian(string(testBlock[byteC:byteC+16]))
		byteC += 16
		transactions.pkScript_length = convertEndian(string(testBlock[byteC:byteC+2]))
		byteC += 2
		//fmt.Println(string(testBlock[byteC:byteC+6]))
		byteC += 6 //???????????????
		jump := fromHex(transactions.pkScript_length) *2 -6 -4
		transactions.pkScript = string(testBlock[byteC:byteC+jump])
		byteC += jump
		//fmt.Println(convertEndian(string(testBlock[byteC:byteC+4]))) //lock time + 88ac
		byteC += 4
		numberOutputs := fromHex(transactions.numOutputs)
		if j==numberOutputs {
			byteC += 8 //lock time - just at end of outputs
		}
		fmt.Println("-------------------------------")
		fmt.Println("Transaction", i, "Output", j, "/", numberOutputs )
		fmt.Println("-------------------------------")

		//100 inputs: 000000000000000001643f7706f3dcbc3a386e4c1bfba852ff628d8024f875b6
		displayTransactionOutputs(transactions)			
	}
}

func buildHeader(testBlock []byte) {
	byteC += 8
	head.prevBlock = convertEndian(string(testBlock[byteC:byteC+64]))
	byteC += 64
	head.merkleRoot = convertEndian(string(testBlock[byteC:byteC+64]))
	head.merkleRootTest = string(testBlock[byteC:byteC+64])
	byteC += 64
	head.timeStamp = convertEndian(string(testBlock[byteC:byteC+8]))
	byteC += 8
	head.targetDiff = convertEndian(string(testBlock[byteC:byteC+8]))
	byteC += 8
	head.nonce = convertEndian(string(testBlock[byteC:byteC+8]))
	byteC += 8

	//account for variable langth
	varLengthNumTrans := convertEndian(string(testBlock[byteC:byteC+2]))

	if varLengthNumTrans!="fd" && varLengthNumTrans!="fe" && varLengthNumTrans!="ff" {
		head.numTransactions = convertEndian(string(testBlock[byteC:byteC+2]))
		byteC += 2
	} else {
		jump := varLength(varLengthNumTrans)
		head.numTransactions = convertEndian(string(testBlock[byteC+2:byteC+jump]))
		byteC += jump
	}
}

func displayTransactionInputs(transactions transaction) {
	fmt.Println("version number: ", fromHex(transactions.transVersion))
	numberInputs := fromHex(transactions.numInputs)
	fmt.Println("number of inputs: ", numberInputs)
	///previous transactions loop start/////////
	for p:=1; p<=numberInputs; p++ {
		fmt.Println("previous transaction",p, ":", transactions.prevTrans[p])
	}
}

func displayTransactionOutputs(transactions transaction) {
	//fmt.Println("transaction index: ", transactions.transIndex)
	fmt.Println("script length: ", fromHex(transactions.scriptLength), "bytes")
	fmt.Println("number of outputs: ", fromHex(transactions.numOutputs))
	fmt.Println("amount: ", float64(fromHex(transactions.amountBTC))/100000000, "BTC")
	fmt.Println("pk_script length: ", fromHex(transactions.pkScript_length), "bytes")
	fmt.Println("receiver address: ", transactions.pkScript, "(hash 160)")
}

func displayHeader() {
	fmt.Println("-------------------------------")
	fmt.Println("Header")
	fmt.Println("-------------------------------")
	fmt.Println("previous block: ", head.prevBlock)
	//fmt.Println("merkle root test: ", head.merkleRootTest)
	fmt.Println("merkle root: ", head.merkleRoot)
	timeStamp := fromHex(head.timeStamp)
	fmt.Println("timestamp: ", timeStamp, "(unix time)")
    timeNotUnix := time.Unix(int64(timeStamp), 0)
    fmt.Println("timestamp: ", timeNotUnix, "(converted)")
	fmt.Println("target difficulty: ", head.targetDiff)
	fmt.Println("nonce: ", head.nonce)
	//fmt.Println("variable length: ", varLengthNumTrans)
	fmt.Println("number of transactions: ", fromHex(head.numTransactions))
}

func convertEndian (conversion string) string{
	var convertEndian string
	for i:=0; i < len(conversion); i=i+2 {
		convertEndian = convertEndian + string(conversion[len(conversion) -i -2])
		convertEndian = convertEndian + string(conversion[len(conversion) -i -1])	
	}
	return convertEndian
}

func varLength (conversion string) int {
	var byteC int
	if conversion=="fd"{
		byteC = 6
	} else if conversion=="fe"{
		byteC = 10
	} else if conversion=="ff"{
		byteC = 18
	} 
	return byteC
}

func fromHex (conversion string) int {
	var convertEndianed int
	var multiple int
	for i:=0; i < len(conversion); i++ {
		amount := 0
		if conversion[i] > 60 {
			multiple = int(conversion[i]) -87
		} else {
			multiple = int(conversion[i]) -48
		}
		
		amount = (multiple * findPower(16, len(conversion) -i -1))
		convertEndianed += amount
	}
	//fmt.Println(convertEndianed)
	return convertEndianed
}

func findPower(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
