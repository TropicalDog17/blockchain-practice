package merkle

type Transaction struct {
	data  string
	label string
}

type Block struct {
	merkleRoot string
	txList     []Transaction
}
