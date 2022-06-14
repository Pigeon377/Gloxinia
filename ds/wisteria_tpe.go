package ds

// data struct was divided into two part
// type ident + data
// type ident => (1 byte,from 70 to 255) (ascii code 2 can't be used,because it is separator byte)
// data       => data,be care for exclude ascii code 2
// because I use skip list and sstable as storage data struct
// key always sort

const (
	intIdent        byte = 70
	floatIdent      byte = 71
	stringIdent     byte = 72
	listIntIdent    byte = 73
	listFloatIdent  byte = 74
	listStringIdent byte = 75
	setIdent        byte = 76
	mapIdent        byte = 77
	jsonIdent       byte = 78
)

// Wisteria
// bytes exclude ident
// it should be recognized in this file
type Wisteria interface {
	ToBytes() []byte
}