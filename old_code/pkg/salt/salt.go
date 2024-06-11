package salt

import (
	"encoding/hex"
	"hash"
	"math/rand"
)

type Service interface {
	Hashed(str, salt []byte) string
	RandBytesSlice(lens int) []byte
}

type saltService struct {
	hash.Hash
}

func NewSaltService(h hash.Hash) Service {
	return &saltService{h}
}

func (s *saltService) Hashed(password, salt []byte) string {
	defer s.Hash.Reset()
	s.Hash.Write(salt)
	s.Hash.Write(password)
	return hex.EncodeToString(s.Hash.Sum(nil))
}

func (s *saltService) RandBytesSlice(lens int) []byte {
	Bytes := make([]byte, lens)

	for i := 0; i < lens; i++ {
		num := 48 + rand.Intn(74)
		Bytes[i] = byte(num)
	}
	return Bytes
}
