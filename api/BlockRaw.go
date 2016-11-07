package Blockrio

/*
{
  "status":"success",
  "data":{
    "hash":"0000000000000000038032f1019409ebc05a875bca076cc47fc511bd72589c71",
    "confirmations":1,
    "size":871829,
    "height":437571,
    "version":536870912,
    "merkleroot":"c10627ca191c0d86ad034a32092f848d25c7aa769c3d32131b0e08a03a3129b3",
    "tx":[
      "fa5a60a9a6f07615b3ec722a354082ab80cca2af3aeb7af487324912b1e51c6e",
      "c1c7fa3e43d5ce0232f92bf13779c29e01e5223b349d763df3e7e20e3dcb0c9c",
    ],
    "time":1478419344,
    "mediantime":1478417245
    ,"nonce":1710500624,
    "bits":"18045174",
    "difficulty":254620187304.06,
    "chainwork":"0000000000000000000000000000000000000000002d652723402cd446aa588c",
    "previousblockhash":"000000000000000002e43e97dbd7c80138dda9ddded0fd976926e47852784354"
  },
    "code":200,
    "message":""
}
*/

type ResponseBlockRaw struct {
	Status  string     `json:"status"`
	Data    []_DataRow `json:"data"`
	Code    float64    `json:"code"`
	Message string     `json:"message"`
}

type _DataRow struct {
	Time              float64  `json:"time"`
	MedianTime        float64  `json:"mediantime"`
	Nonce             float64  `json:"nonce"`
	Bits              string   `json:"bits"`
	Difficulty        float64  `json:"difficulty"`
	ChainWork         string   `json:"chainwork"`
	PreviousBlockHach string   `json:"previousblockhash"`
	Hash              string   `json:"hash"`
	Confirmations     float64  `json:"confirmations"`
	Size              float64  `json:"size"`
	Height            float64  `json:"height"`
	Version           float64  `json:"Version"`
	MerkleRoot        string   `json:"merkleroot"`
	Tx                []string `json:"tx"`
}
