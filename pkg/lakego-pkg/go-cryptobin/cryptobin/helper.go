package cryptobin

// 构造函数
func New() Cryptobin {
    return Cryptobin{
        multiple: "Aes",
        mode:     "ECB",
        padding:  "",
        config:   make(map[string]any),
    }
}

// 字节
func FromBytes(data []byte) Cryptobin {
    return New().FromBytes(data)
}

// 字符
func FromString(data string) Cryptobin {
    return New().FromString(data)
}

// Base64
func FromBase64String(data string) Cryptobin {
    return New().FromBase64String(data)
}

// Hex
func FromHexString(data string) Cryptobin {
    return New().FromHexString(data)
}