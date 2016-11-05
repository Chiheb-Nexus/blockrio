package main
import (
	"fmt"
	"Blockr.io/api"
)

func main() {
	//info := Blockrio.GetCoinInfo("dgc")
	//fmt.Printf("%T\n%v\n",info, info)
	//fmt.Printf("%T\n%0.8f\n", info.Data.Markets.Btce.Value, info.Data.Markets.Btce.Value)
	// For dgc
	//fmt.Printf("%T\n%0.8f\n", info.Data.Markets.Cryptsy.Value, info.Data.Markets.Cryptsy.Value)
	//fmt.Println("-----------------------------------------")
	
	//current := Blockrio.GetExchangeCurrent("btc")
	//fmt.Println(current)
	
	
	fmt.Println("-----------------------------------------")
	a := []string{"12345", "12346"}
	
	//block := Blockrio.GetBlockInfo("btc", a)
	//fmt.Println(block)
	
	fmt.Println("-----------------------------------------")
	
	block_tx := Blockrio.GetBlockTxs("btc", a)
	fmt.Println(block_tx)
	
}
	
	
	
	
		




	
