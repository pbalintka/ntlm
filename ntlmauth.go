package main

import (
    "encoding/base64"
    "encoding/binary"
	"fmt"
)

func GetInt(s []byte) int {
	var b [8]byte

    copy(b[:8-len(s)], s)
	return int(binary.LittleEndian.Uint64(b[:]))
}

func main() {
    var negmsg string = "TlRMTVNTUAABAAAAA7IAAAoACgApAAAACQAJACAAAABMSUdIVENJVFlVUlNBLU1JTk9S"
    negmsg = "TlRMTVNTUAABAAAABzIAAAYABgArAAAACwALACAAAABXT1JLU1RBVElPTkRPTUFJTg=="

    var chmsg string = "TlRMTVNTUAACAAAAAAAAACgAAAABggAAU3J2Tm9uY2UAAAAAAAAAAA=="
    //chmsg = "TlRMTVNTUAACAAAADAAMADAAAAABAoEAASNFZ4mrze8AAAAAAAAAAGIAYgA8AAAARABPAE0AQQBJAE4AAgAMAEQATwBNAEEASQBOAAEADABTAEUAUgBWAEUAUgAEABQAZABvAG0AYQBpAG4ALgBjAG8AbQADACIAcwBlAHIAdgBlAHIALgBkAG8AbQBhAGkAbgAuAGMAbwBtAAAAAAA="

    fmt.Println(negmsg)

    s, _ := base64.StdEncoding.DecodeString(negmsg)
    fmt.Println(string(s[0:8])) // protocol
    fmt.Println(uint8(s[8])) // type
    fmt.Println(uint8(s[12])) // flags
    fmt.Println(uint8(s[13])) // flags
    fmt.Println(uint8(s[14])) // flags
    fmt.Println(uint8(s[15])) // flags
    fmt.Println(GetInt(s[16:18])) // dom length
    fmt.Println(GetInt(s[18:20])) // dom max length
    fmt.Println(GetInt(s[20:22])) // dom offset
    fmt.Println(GetInt(s[24:26])) // host length
    fmt.Println(GetInt(s[28:30])) // host offset
    fmt.Println(string(s[32:41])) // hostname
    fmt.Println(string(s[41:51])) // domain

    sc, _ := base64.StdEncoding.DecodeString(chmsg)
    fmt.Println(string(sc[0:8])) // protocol
    fmt.Println(uint8(sc[8])) // type
    fmt.Println(GetInt(sc[16:18])) // msg len
    fmt.Println(string(sc[24:32])) // payload

    ret := make([]byte, 40)
    copy(ret, []byte("NTLMSSP\x00"))
    binary.LittleEndian.PutUint16(ret[8:], 2)
    binary.LittleEndian.PutUint16(ret[16:], 40)
    binary.LittleEndian.PutUint16(ret[20:], 0x01)
    binary.LittleEndian.PutUint16(ret[21:], 0x82)
    copy(ret[24:], []byte("SrvNonce"))

    fmt.Println(sc)
    fmt.Println(ret)

    e := base64.StdEncoding.EncodeToString(ret)
    fmt.Println(e)
}

