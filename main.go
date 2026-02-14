package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// problem --> https://neetcode.io/problems/string-encode-and-decode/

func main() {
	s := Solution{}
	arr := []string{"", ""}
	fmt.Println(arr)

	enc := s.Encode(arr)
	dec := s.Decode(enc)

	// fmt.Println(enc)
	fmt.Println(dec)
}

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var res string
	for _, v := range strs {
		enc := base64.StdEncoding.EncodeToString([]byte(v))
		res += enc + ":"
	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
	counter, all := 0, len(strings.Split(encoded, ":"))
	var res []string

	arr := strings.Split(encoded, ":")

	for _, v := range arr {
		counter++
		if all == counter {
			break
		}
		dec, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			fmt.Println("error: ", err, "dec: ", string(dec))
		}
		res = append(res, string(dec))

	}

	return res
}
