package utility

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/QOSGroup/qbase/types"
	"github.com/QOSGroup/qstars/wire"
	"github.com/prometheus/common/log"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/bech32"
)

const (
	// expected address length
	//AddrLen = 20

	// Bech32 prefixes
	Bech32PrefixAccAddr = types.PREF_ADD
	Bech32PrefixAccPub  = "cosmosaccpub"
)

func Encbase64(input []byte) string {
	return base64.StdEncoding.EncodeToString(input[:])
}

func Decbase64(input string) []byte {
	bz, _ := base64.StdEncoding.DecodeString(input)
	return bz

}

func PubAddrRetrievalFromAmino(caPriBase64 string, cdc *wire.Codec) (string, string, ed25519.PrivKeyEd25519) {
	caHex := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + caPriBase64 + "\"}"
	var key ed25519.PrivKeyEd25519
	err := cdc.UnmarshalJSON([]byte(caHex), &key)
	if err != nil {
		log.Error(err.Error())
		return "", "", key
	}

	pub := key.PubKey().Bytes()
	addr := key.PubKey().Address()
	bech32Pub, _ := bech32.ConvertAndEncode(Bech32PrefixAccPub, pub)
	bech32Addr, _ := bech32.ConvertAndEncode(Bech32PrefixAccAddr, addr.Bytes())

	return bech32Pub, bech32Addr, key
}

func PubAddrRetrievalFromHex1(caPriHex string, cdc *wire.Codec) (string, string, ed25519.PrivKeyEd25519) {
	caHex, _ := hex.DecodeString(caPriHex[2:])
	var key ed25519.PrivKeyEd25519
	cdc.MustUnmarshalBinaryBare(caHex, &key)

	//bz := Decbase64(s)
	//var key ed25519.PrivKeyEd25519
	//copy(key[:], caHex)
	pub := key.PubKey().Bytes()
	addr := key.PubKey().Address()
	bech32Pub, _ := bech32.ConvertAndEncode(Bech32PrefixAccPub, pub)
	bech32Addr, _ := bech32.ConvertAndEncode(Bech32PrefixAccAddr, addr.Bytes())
	//privJson,_ := cdc.MarshalJSONIndent(key,"","\t")
	//fmt.Println(string(privJson))
	//fmt.Println(bech32Pub)
	//fmt.Println(bech32Addr)
	return bech32Pub, bech32Addr, key
}

//func main() {
//	s := "9Rg9mNEXVh9aUsxJ74Ogqe8O6wrBw8EeMhyK/GgHcfUsGprPgC7YXH6YEwGM+eXmc7oV1ci7ivlxo7k6amd3Lg=="
//	PubAddrRetrieval(s)
//	bz := Decbase64(s)
//	s1 := Encbase64(bz)
//	fmt.Printf("%x\n",bz)
//	fmt.Println(s1)
//}
