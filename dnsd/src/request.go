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

// hashShrink will define how much of the SHA256 sum of a request will actually be used
const hashShrink int = 8

// Request defines a DNS request object
type Request struct {
	Hash string `json:"hash"`
	Time int64 `json:"time"`
	Address string `json:"address"`
	Name string `json:"name"`
	Type uint16 `json:"type"`
	Class uint16 `json:"class"`
}

// CreateRequest will create a new request object
func CreateRequest(responseWriter dns.ResponseWriter, question dns.Question) Request {
	request := Request{}
	request.Time = int64(time.Now().Unix())
	request.Address = strings.Split(responseWriter.RemoteAddr().String(), ":")[0]
	request.Name = question.Name
	request.Type = question.Qtype
	request.Class = question.Qclass
	request.Hash = hex.EncodeToString(toHash(request))
	return request
}

// toBytes will serialize the request into a byte array
func toBytes(request Request) []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, request.Time)
	buffer.WriteString(request.Address)
	buffer.WriteString(request.Name)
	binary.Write(buffer, binary.BigEndian, request.Type)
	binary.Write(buffer, binary.BigEndian, request.Class)
	return buffer.Bytes()
}

// toHash will return the SHA256 sum of the request
func toHash(request Request) []byte {
	hash := sha256.Sum256(toBytes(request))
	return hash[:hashShrink]
}
