package validators

import (
	"math/big"

	"github.com/sunvim/yaoguang/crypto"
)

type Writer interface {
	SetPower(id *crypto.PublicKey, power *big.Int) (flow *big.Int, err error)
}

type Reader interface {
	Power(id crypto.Address) (*big.Int, error)
}

type Iterable interface {
	IterateValidators(func(id crypto.Addressable, power *big.Int) error) error
}

type IterableReader interface {
	Reader
	Iterable
}

type ReaderWriter interface {
	Reader
	Writer
}

type IterableReaderWriter interface {
	ReaderWriter
	Iterable
}

type History interface {
	ValidatorChanges(blocksAgo int) IterableReader
	Validators(blocksAgo int) IterableReader
}
