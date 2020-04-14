package dnsd

import (
	"time"

	"bytes"
	"strings"

	"crypto/sha256"

	"encoding/binary"
	"encoding/hex"

	"github.com/miekg/dns"
)

// hashShrink will define how much of the SHA256 sum of a query will actually be used
const hashShrink int = 8

// Query defines a DNS query object
type Query struct {
	Hash string `json:"hash"`
	Time int64 `json:"time"`
	Address string `json:"address"`
	Name string `json:"name"`
	Type string `json:"type"`
	Class string `json:"class"`
}

// CreateQuery will create a new query object
func CreateQuery(responseWriter dns.ResponseWriter, question dns.Question) Query {
	query := Query{}
	query.Time = int64(time.Now().Unix())
	query.Address = strings.Split(responseWriter.RemoteAddr().String(), ":")[0]
	query.Name = strings.TrimSuffix(question.Name, ".")
	query.Type = ConvertType(question.Qtype)
	query.Class = ConvertClass(question.Qclass)
	query.Hash = hex.EncodeToString(toHash(query))
	return query
}

// toBytes will serialize the query into a byte array
func toBytes(query Query) []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, query.Time)
	buffer.WriteString(query.Address)
	buffer.WriteString(query.Name)
	binary.Write(buffer, binary.BigEndian, query.Type)
	binary.Write(buffer, binary.BigEndian, query.Class)
	return buffer.Bytes()
}

// toHash will return the SHA256 sum of the query
func toHash(query Query) []byte {
	hash := sha256.Sum256(toBytes(query))
	return hash[:hashShrink]
}
