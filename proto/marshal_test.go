package proto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	r := Reader{
		B: []byte{0x0, 0x0, 0x1, 0x7a, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x1, 0x0, 0x4, 0x74, 0x65, 0x73, 0x74, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x0, 0x0, 0x1, 0x56, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2d, 0x9, 0xf8, 0xf8, 0x91, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x32, 0x30, 0x31, 0x35, 0x2d, 0x30, 0x39, 0x2d, 0x31, 0x35, 0x54, 0x32, 0x32, 0x3a, 0x31, 0x38, 0x3a, 0x30, 0x31, 0x2b, 0x30, 0x38, 0x3a, 0x30, 0x30, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x2d, 0x5b, 0xc0, 0xd7, 0x36, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x32, 0x30, 0x31, 0x35, 0x2d, 0x30, 0x39, 0x2d, 0x31, 0x35, 0x54, 0x32, 0x32, 0x3a, 0x31, 0x38, 0x3a, 0x30, 0x34, 0x2b, 0x30, 0x38, 0x3a, 0x30, 0x30, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x2d, 0xfd, 0xb7, 0xdc, 0x82, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x32, 0x30, 0x31, 0x35, 0x2d, 0x30, 0x39, 0x2d, 0x31, 0x35, 0x54, 0x32, 0x32, 0x3a, 0x31, 0x38, 0x3a, 0x30, 0x35, 0x2b, 0x30, 0x38, 0x3a, 0x30, 0x30, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x2d, 0xe7, 0x6f, 0xfd, 0x41, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x32, 0x30, 0x31, 0x35, 0x2d, 0x30, 0x39, 0x2d, 0x31, 0x35, 0x54, 0x32, 0x32, 0x3a, 0x31, 0x38, 0x3a, 0x35, 0x30, 0x2b, 0x30, 0x38, 0x3a, 0x30, 0x30, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x2d, 0x41, 0x18, 0xf6, 0xf5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x32, 0x30, 0x31, 0x35, 0x2d, 0x30, 0x39, 0x2d, 0x31, 0x35, 0x54, 0x32, 0x32, 0x3a, 0x31, 0x38, 0x3a, 0x35, 0x31, 0x2b, 0x30, 0x38, 0x3a, 0x30, 0x30, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0x0, 0x0, 0x0, 0x2d, 0x13, 0x20, 0xd9, 0x52, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x32, 0x30, 0x31, 0x35, 0x2d, 0x30, 0x39, 0x2d, 0x31, 0x35, 0x54, 0x32, 0x32, 0x3a, 0x31, 0x38, 0x3a, 0x35, 0x34, 0x2b, 0x30, 0x38, 0x3a, 0x30, 0x30},
	}
	fmt.Println(len(r.B))
	(&RequestOrResponse{M: &Response{ResponseMessage: &FetchResponse{}}}).Unmarshal(&r)
}

func TestUnmarshalTrunc(t *testing.T) {
	f, err := os.Open("truncated_response")
	if err != nil {
		t.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	r := Reader{
		B: b,
	}
	resp := FetchResponse{}
	(&RequestOrResponse{M: &Response{ResponseMessage: &resp}}).Unmarshal(&r)
	if r.Err != nil {
		t.Fatal(r.Err)
	}
	if len(resp) != 1 ||
		len(resp[0].FetchMessageSetInPartitions) != 1 ||
		len(resp[0].FetchMessageSetInPartitions[0].MessageSet) != 2 {
		t.Fatal("fail to parse truncated message set")
	}
}

func toJSON(v interface{}) string {
	buf, _ := json.MarshalIndent(v, "", "    ")
	return string(buf)
}
