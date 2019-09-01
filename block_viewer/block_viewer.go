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
	targetDiff1 string
	targetDiff2 string
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

var byteCounter int
var head header

//https://blockchain.info/block/000000000000000001f942eb4bfa0aeccb6a14c268f4c72d5fff17270da771b9?format=hex
//100 inputs: 000000000000000001643f7706f3dcbc3a386e4c1bfba852ff628d8024f875b6

func main() {

	arg := os.Args[1]

	fmt.Println()
	fmt.Println("-------------------------------")
	fmt.Println("Welcome to Bitcoin Block Viewer")
	fmt.Println("-------------------------------")
	fmt.Println()
	
	/* get raw block data from blockchain.com  */
	URL := "https://blockchain.info/block/" + arg + "?format=hex"
	resp, _ := http.Get(URL)
	testBlock, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()	
	
	/* build and display header */
	buildHeader(testBlock)
	displayHeader()

	var transactions [25]transaction

	/////////////////////////start transaction loop////////////////////////////////
	for i:=1; i<=5; i++ {

		transactions[i].transVersion = convertEndian(string(testBlock[byteCounter:byteCounter+8]))
		byteCounter += 8

		/* account for variable length and find the number of inputs */
		varLengthNumTrans := convertEndian(string(testBlock[byteCounter:byteCounter+2]))
		if varLengthNumTrans!="fd" && varLengthNumTrans!="fe" && varLengthNumTrans!="ff" {
			transactions[i].numInputs = convertEndian(string(testBlock[byteCounter:byteCounter+2]))
			byteCounter += 2
		} else {
			jump := varLength(varLengthNumTrans)
			transactions[i].numInputs = convertEndian(string(testBlock[byteCounter+2:byteCounter+jump]))
			byteCounter += jump
		}

		//////////////////////start number of inputs loop////////////////////
		/* loop throught the number of inputs to display previous transactions */
		for k:=1;k<=fromHex(transactions[i].numInputs);k++ {
			transactions[i].prevTrans[k] = convertEndian(string(testBlock[byteCounter:byteCounter+64]))
			byteCounter += 64
			transactions[i].transIndex = convertEndian(string(testBlock[byteCounter:byteCounter+8]))
			byteCounter += 8

			// account for variable langth for script length
			scriptLengthVar := convertEndian(string(testBlock[byteCounter:byteCounter+2]))
    		if scriptLengthVar!="fd" && scriptLengthVar!="fe" && scriptLengthVar!="ff" {
				transactions[i].scriptLength = convertEndian(string(testBlock[byteCounter:byteCounter+2]))
				byteCounter += 2
			} else {
				jump := varLength(scriptLengthVar)
				transactions[i].scriptLength = convertEndian(string(testBlock[byteCounter+2:byteCounter+jump]))
				byteCounter += jump
			}

			byteCounter += fromHex(transactions[i].scriptLength)*2
			byteCounter += 8
		}
		//////////////////////end number of inputs loop////////////////////

		/* account for variable length and find the number of outputs */
		varLengthNumTrans2 := convertEndian(string(testBlock[byteCounter:byteCounter+2]))
		if varLengthNumTrans2!="fd" && varLengthNumTrans2!="fe" && varLengthNumTrans2!="ff" {
			transactions[i].numOutputs = convertEndian(string(testBlock[byteCounter:byteCounter+2]))
			byteCounter += 2
		} else {
			jump2 := varLength(varLengthNumTrans2)
			transactions[i].numOutputs = convertEndian(string(testBlock[byteCounter+2:byteCounter+jump2]))
			byteCounter += jump2
		}

		fmt.Println("-------------------------------")
		fmt.Println("Transaction", i, "Inputs")
		fmt.Println("-------------------------------")

		/* display transaction inputs */
		displayTransactionInputs(transactions[i])

		/* build and display transaction outputs. one block for every output */
		buildOutput(i, transactions[i], testBlock)
	}
	////////////////end transaction loop////////////////////////////////	
}

func buildVarLength() {

}

/* iterate and build the transaction outputs */
func buildOutput(i int, transactions transaction, testBlock []byte) {
	for j:=1;j<=fromHex(transactions.numOutputs);j++ {
		transactions.amountBTC = convertEndian(string(testBlock[byteCounter:byteCounter+16]))
		byteCounter += 16
		transactions.pkScript_length = convertEndian(string(testBlock[byteCounter:byteCounter+2]))
		byteCounter += 2
		byteCounter += 6
		jump := fromHex(transactions.pkScript_length) *2 -6 -4
		transactions.pkScript = string(testBlock[byteCounter:byteCounter+jump])
		byteCounter += jump
		byteCounter += 4
		numberOutputs := fromHex(transactions.numOutputs)
		if j==numberOutputs {
			byteCounter += 8 //lock time - just at end of outputs
		}
		fmt.Println("-------------------------------")
		fmt.Println("Transaction", i, "Output", j, "/", numberOutputs )
		fmt.Println("-------------------------------")
		displayTransactionOutputs(transactions)			
	}
}

/* iterate and extract the header information */
func buildHeader(testBlock []byte) {
	byteCounter += 8
	head.prevBlock = convertEndian(string(testBlock[byteCounter:byteCounter+64]))
	byteCounter += 64
	head.merkleRoot = convertEndian(string(testBlock[byteCounter:byteCounter+64]))
	head.merkleRootTest = string(testBlock[byteCounter:byteCounter+64])
	byteCounter += 64
	head.timeStamp = convertEndian(string(testBlock[byteCounter:byteCounter+8]))
	byteCounter += 8
	head.targetDiff1 = convertEndian(string(testBlock[byteCounter:byteCounter+8]))
	byteCounter += 2
	head.targetDiff2 = convertEndian(string(testBlock[byteCounter:byteCounter+6]))
	byteCounter += 6
	head.nonce = convertEndian(string(testBlock[byteCounter:byteCounter+8]))
	byteCounter += 8

	//account for variable langth
	varLengthNumTrans := convertEndian(string(testBlock[byteCounter:byteCounter+2]))
	if varLengthNumTrans!="fd" && varLengthNumTrans!="fe" && varLengthNumTrans!="ff" {
		head.numTransactions = convertEndian(string(testBlock[byteCounter:byteCounter+2]))
		byteCounter += 2
	} else {
		jump := varLength(varLengthNumTrans)
		head.numTransactions = convertEndian(string(testBlock[byteCounter+2:byteCounter+jump]))
		byteCounter += jump
	}
}

/* display transaction input block and iterate through the number of inputs
showing each previous transaction */
func displayTransactionInputs(transactions transaction) {
	fmt.Println("version number: ", fromHex(transactions.transVersion))
	numberInputs := fromHex(transactions.numInputs)
	fmt.Println("number of inputs: ", numberInputs)
	for p:=1; p<=numberInputs; p++ {
		fmt.Println("previous transaction",p, ":", transactions.prevTrans[p])
	}
}

/* display a separate block for each transaction output */
func displayTransactionOutputs(transactions transaction) {
	fmt.Println("script length: ", fromHex(transactions.scriptLength), "bytes")
	fmt.Println("number of outputs: ", fromHex(transactions.numOutputs))
	fmt.Println("amount: ", float64(fromHex(transactions.amountBTC))/100000000, "BTC")
	fmt.Println("pk_script length: ", fromHex(transactions.pkScript_length), "bytes")
	fmt.Println("receiver address: ", transactions.pkScript, "(hash 160)")
}

/* display header - self explanatory */
func displayHeader() {
	fmt.Println("-------------------------------")
	fmt.Println("Header")
	fmt.Println("-------------------------------")
	fmt.Println("previous block: ", head.prevBlock)
	fmt.Println("merkle root: ", head.merkleRoot)
	timeStamp := fromHex(head.timeStamp)
	fmt.Println("timestamp: ", timeStamp, "(unix time)")
    timeNotUnix := time.Unix(int64(timeStamp), 0)
	fmt.Println("timestamp: ", timeNotUnix, "(converted)")
	//difficulty
	difficult := float64(65535 / float64(fromHex(head.targetDiff2)) * float64(findPower(2,40)))
	fmt.Println("target difficulty: ", head.targetDiff1)
	fmt.Println("target difficulty: ", difficult, "(converted)")

	fmt.Println("nonce: ", fromHex(head.nonce))
	//fmt.Println("variable length: ", varLengthNumTrans)
	fmt.Println("number of transactions: ", fromHex(head.numTransactions))
}

/* convertEndian takes a string and converts it from big endian 
to little or vice versa. ex. a1bd => bda1 */
func convertEndian (conversion string) string{
	var convertEndian string
	for i:=0; i < len(conversion); i=i+2 {
		convertEndian = convertEndian + string(conversion[len(conversion) -i -2])
		convertEndian = convertEndian + string(conversion[len(conversion) -i -1])	
	}
	return convertEndian
}

/* varLength deals with variable length in a Bitcoin block. 
for example if the varLength byte is fd then there are 2 bytes
after that expressing the size of the number. If the varLength 
byte is not fd, fe, or ff then the number is represented by one byte */
func varLength (conversion string) int {
	var byteCounter int
	if conversion=="fd"{
		byteCounter = 6
	} else if conversion=="fe"{
		byteCounter = 10
	} else if conversion=="ff"{
		byteCounter = 18
	} 
	return byteCounter
}

/* fromHex function just converts a a hex string into a regular number.
ex. 10a1 => 4257 */
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
	return convertEndianed
}

/* the findPower function simply returns a^b - 
this is the only code that I didn't write myself */
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
