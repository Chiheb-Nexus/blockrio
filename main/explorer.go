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

/*	current := Blockrio.GetExchangeCurrent("btc")
	fmt.Println(current)


	fmt.Println("-----------------------------------------")
	aa := []string{"12345", "12346"}

	block := Blockrio.GetBlockInfo("btc", aa)
	fmt.Println(block)

	fmt.Println("-----------------------------------------")

	// FIXME: Bug with dgc

	block_tx := Blockrio.GetBlockTxs("ltc", aa)
	fmt.Println(block_tx)

	fmt.Println("-----------------------------------------")
	a := []string{"60c1f1a3160042152114e2bba45600a5045711c3a8a458016248acec59653471", "4addbc5ec75e087a44a34e8c1c3bd05fd495771072879a89a8c9aaa356255cb2"}
	tx_info := Blockrio.GetTxsInfo("btc", a)
	fmt.Println(tx_info)
	fmt.Println("-----------------------------------------")

	block_raw := Blockrio.GetBlockRaw("btc", []string{"12345","12346"})
	fmt.Println(block_raw[0].Status)

	fmt.Println("-----------------------------------------")

	add := []string{"198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi","1L8meqhMTRpxasdGt8DHSJfscxgHHzvPgk"}
	address_info := Blockrio.GetAddressInfo("btc", add)
	fmt.Println(address_info)


	fmt.Println("-----------------------------------------")
	/*add := []string{"198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi"}
	address_balance := Blockrio.GetAddressBalance("btc", add)
	fmt.Println(address_balance)
	*/
	fmt.Println("-----------------------------------------")
	add := []string{"198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi"}
	// FIXME: Error handling JSON
	address_txs := Blockrio.GetAddressTransactions("btc", add)
	fmt.Println(address_txs)

	fmt.Println("-----------------------------------------")
	// FIXME: Error handling JSON
	//addd := []string{"198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi"}
	//unspent_add := Blockrio.GetAddressUnspent("btc", addd )
	//fmt.Println(unspent_add)

	fmt.Println("-----------------------------------------")
	//adds := []string{"1NZ3a62JrXmvFNnL8Pqa9owSX13H8Tm4wA"}
	//unconfirmed_balance := Blockrio.GetUnconfirmedBalance("btc", adds)
	//fmt.Println(unconfirmed_balance)




}
