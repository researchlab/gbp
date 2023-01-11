package utils

import (
	"bytes"
	"crypto/des"
	"encoding/base64"

)

/*
一、简介
DES(Data Encryption Standard)由IBM公司在1971年提出。使用同一个密钥来加密和解密数据, 简称对称密钥算法。

二、密钥
DES使用一个56位的初始密钥，但是这里提供的是一个64位的值，这是因为在硬件实现中每8位可以用于奇偶校验。可以通过设定8位字符串，由crypto/des库的des.NewCipher(key)函数生成密钥

三、填充算法
DES分组的大小是64位，如果加密的数据长度不是64位的倍数，可以按照某种具体的规则来填充位。常用的填充算法有pkcs5，zero等
*/

// pkcs5 补码算法
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// pkcs5 减码算法
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// zero 补码算法
func zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func zeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

var (
	KEY = []byte("12345678") // des 最大的key长度为8 , 超过8 会报错
	IV  = []byte("abcdefghijk")
)

type Result struct {
	data []byte
}

func (p Result) Bytes() []byte{
	return p.data
}

func (p Result) Base64String()string {
	return	base64.StdEncoding.EncodeToString(p.data)
}

/*
四、加密模式
密码分组链接模式（Cipher Block Chaining，简称CBC）：是一种循环模式，前一个分组的密文和当前分组的明文异或操作后再加密，这样做的目的是增强破解难度。
电码本模式（Electronic Codebook Book，简称ECB）：是一种基础的加密方式，密文被分割成分组长度相等的块（不足补齐），然后单独一个个加密，一个个输出组成密文。
计算器模式（Counter，简称CTR）：计算器模式不常见，在CTR模式中， 有一个自增的算子，这个算子用密钥加密之后的输出和明文异或的结果得到密文，相当于一次一密。这种加密方式简单快速，安全可靠，而且可以并行加密，但是在计算器不能维持很长的情况下，密钥只能使用一次。
输出反馈模式（Output FeedBack，简称OFB）：实际上是一种反馈模式，目的也是增强破解的难度。
密码反馈模式（Cipher FeedBack，简称CFB）：实际上是一种反馈模式，目的也是增强破解的难度。
*/

// 返回密文数据
// data 明文数据
func DesECBEncrypt(data, key []byte) Result {
	// NewCipher 创建一个新的加密块
	block, err := des.NewCipher(key)
	if err != nil {
		return Result{} 
	}
	bs := block.BlockSize()
	// pkcs5 填充
	data = pkcs5Padding(data, bs)
	if len(data)%bs != 0 {
		return Result{} 
	}
	r := Result{
		data: make([]byte, len(data)),
	}
	dst := r.data
	for len(data) > 0 {
		//Encrypt 加密第一个块，将其结果保存到dst
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return r
}

// 返回明文数据
// data 密文数据
func DesECBDecrypter(data, key []byte) Result {
	//NewCipher 创建一个新的加密块
	block, err := des.NewCipher(key)
	if err != nil {
		return Result{} 
	}

	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return Result{} 
	}

	r := Result{
		data: make([]byte, len(data)),
	}
	dst := r.data
	for len(data) > 0 {
		//Encrypt 加密第一个块, 将其结果保存到dst
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	// pkcs5 填充
	r.data = pkcs5UnPadding(r.data)
	return r
}

//---------------DES ECB解密--------------------
// data: 密文数据
// key: 密钥字符串
// 返回明文数据
func DesECBDecrypter(data, key []byte) Result {
    //NewCipher创建一个新的加密块
    block, err := des.NewCipher(key)
    if err != nil {
        return Result{} 
    }

    bs := block.BlockSize()
    if len(data)%bs != 0 {
        return Result{} 
    }

		r := Result{
		data: make([]byte, len(data)),
		}
    dst := r.data 
    for len(data) > 0 {
        //Encrypt加密第一个块，将其结果保存到dst
        block.Decrypt(dst, data[:bs])
        data = data[bs:]
        dst = dst[bs:]
    }

    // pkcs5填充
    r.data= pkcs5UnPadding(r.data)

    return r 
}

//---------------DES CBC加密--------------------
// data: 明文数据
// key: 密钥字符串
// iv:iv向量
// 返回密文数据
func DesCBCEncrypt(data, key, iv []byte) Result {
    block, err := des.NewCipher(key)
    if err != nil {
        return Result{} 
    }

    // pkcs5填充
    data = pkcs5Padding(data, block.BlockSize())
    cryptText := make([]byte, len(data))

    blockMode := cipher.NewCBCEncrypter(block, iv)
    blockMode.CryptBlocks(cryptText, data)
		return Result{data:cryptText}
}

//---------------DES CBC解密--------------------
// data: 密文数据
// key: 密钥字符串
// iv:iv向量
// 返回明文数据
func DesCBCDecrypter(data, key, iv []byte) Result {
    block, err := des.NewCipher(key)
    if err != nil {
        return Result{} 
    }

    blockMode := cipher.NewCBCDecrypter(block, iv)
    cryptText := make([]byte, len(data))
    blockMode.CryptBlocks(cryptText, data)
    // pkcs5填充
    cryptText = pkcs5UnPadding(cryptText)

		return Result{data:cryptText}
}

//---------------DES CTR加密--------------------
// data: 明文数据
// key: 密钥字符串
// iv:iv向量
// 返回密文数据
func DesCTREncrypt(data, key, iv []byte) Result {
    block, err := des.NewCipher(key)
    if err != nil {
        return Result{} 
    }

    // pkcs5填充
    data = pkcs5Padding(data, block.BlockSize())
    cryptText := make([]byte, len(data))

    blockMode := cipher.NewCTR(block, iv)
    blockMode.XORKeyStream(cryptText, data)
		return Result{data:cryptText}
}

//---------------DES CTR解密--------------------
// data: 密文数据
// key: 密钥字符串
// iv:iv向量
// 返回明文数据
func DesCTRDecrypter(data, key, iv []byte) Result {
    block, err := des.NewCipher(key)
    if err != nil {
        return Result{} 
    }

    blockMode := cipher.NewCTR(block, iv)
    cryptText := make([]byte, len(data))
    blockMode.XORKeyStream(cryptText, data)

    // pkcs5填充
    cryptText = pkcs5UnPadding(cryptText)

		return Result{data:cryptText}
}

//---------------DES OFB加密--------------------
// data: 明文数据
// key: 密钥字符串
// iv:iv向量
// 返回密文数据
func DesOFBEncrypt(data, key, iv []byte) Result {
    block, err := des.NewCipher(key)
    if err != nil {
        return Result{} 
    }

    // pkcs5填充
    data = pkcs5Padding(data, block.BlockSize())
    cryptText := make([]byte, len(data))

    blockMode := cipher.NewOFB(block, iv)
    blockMode.XORKeyStream(cryptText, data)
		return Result{data:cryptText}
}

//---------------DES OFB解密--------------------
// data: 密文数据
// key: 密钥字符串
// iv:iv向量
// 返回明文数据
func DesOFBDecrypter(data, key, iv []byte) Result {
    block, err := des.NewCipher(key)
    if err != nil {
        return Result{} 
    }

    blockMode := cipher.NewOFB(block, iv)
    cryptText := make([]byte, len(data))
    blockMode.XORKeyStream(cryptText, data)

    // pkcs5填充
    cryptText = pkcs5UnPadding(cryptText)

		return Result{data:cryptText}
}

//---------------DES CFB加密--------------------
// data: 明文数据
// key: 密钥字符串
// iv:iv向量
// 返回密文数据
func DesCFBEncrypt(data, key, iv []byte) Result {
    block, err := des.NewCipher(key)
    if err != nil {
        return Result{} 
    }

    // pkcs5填充
    data = pkcs5Padding(data, block.BlockSize())
    cryptText := make([]byte, len(data))

    blockMode := cipher.NewCFBDecrypter(block, iv)
    blockMode.XORKeyStream(cryptText, data)
		return Result{data:cryptText}
}

//---------------DES CFB解密--------------------
// data: 密文数据
// key: 密钥字符串
// iv:iv向量
// 返回明文数据
func DesCFBDecrypter(data, key, iv []byte) Result {
    block, err := des.NewCipher(key)
    if err != nil {
        return Result{} 
    }
    blockMode := cipher.NewCFBEncrypter(block, iv)
    cryptText := make([]byte, len(data))
    blockMode.XORKeyStream(cryptText, data)

    // pkcs5填充
    cryptText = pkcs5UnPadding(cryptText)

		return Result{data:cryptText}
}
