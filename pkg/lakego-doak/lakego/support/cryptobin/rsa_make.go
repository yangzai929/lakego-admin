package cryptobin

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
)

// 私钥
// bits = 512 | 1024 | 2048 | 4096
func (this *Rsa) MakePrvKey(bits int) ([]byte, error) {
    privateKey, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        return nil, err
    }

    X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)

    privateBlock := pem.Block{
        Type: "RSA PRIVATE KEY",
        Bytes: X509PrivateKey,
    }

    rs := pem.EncodeToMemory(&privateBlock)
    return rs, nil
}

// 公钥
func (this *Rsa) MakePubKeyFromPrvKey(prvKey []byte) ([]byte, error) {
    privateKey, err := this.ParseRSAPrivateKeyFromPEM(prvKey)
    if err != nil {
        return nil, err
    }

    publicKey := privateKey.PublicKey

    X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
    if err != nil {
        return nil, err
    }

    publicBlock := pem.Block{
        Type: "PUBLIC KEY",
        Bytes: X509PublicKey,
    }

    rs := pem.EncodeToMemory(&publicBlock)
    return rs, nil
}

// 带密码私钥
func (this *Rsa) MakePassPrvKey(bits int, password string, PEMCipher ...string) ([]byte, error) {
    privateKey, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        return nil, err
    }

    PEMCiphers := map[string]x509.PEMCipher{
        "DES":    x509.PEMCipherDES,
        "3DES":   x509.PEMCipher3DES,
        "AES128": x509.PEMCipherAES128,
        "AES192": x509.PEMCipherAES192,
        "AES256": x509.PEMCipherAES256,
    }

    usePEMCipher := x509.PEMCipherAES256
    if len(PEMCipher) > 0 {
        userPEMCipher, ok := PEMCiphers[PEMCipher[0]]
        if ok {
            usePEMCipher = userPEMCipher
        }
    }

    x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)

    privateBlock, err := x509.EncryptPEMBlock(
        rand.Reader,
        "RSA PRIVATE KEY",
        x509PrivateKey,
        []byte(password),
        usePEMCipher,
    )
    if err != nil {
        return nil, err
    }

    rs := pem.EncodeToMemory(privateBlock)
    return rs, nil
}

// 公钥
func (this *Rsa) MakePubKeyFromPassPrvKey(prvKey []byte, password string) ([]byte, error) {
    privateKey, err := this.ParseRSAPrivateKeyFromPEMWithPassword(prvKey, password)
    if err != nil {
        return nil, err
    }

    publicKey := privateKey.PublicKey

    X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
    if err != nil {
        return nil, err
    }

    publicBlock := pem.Block{
        Type: "PUBLIC KEY",
        Bytes: X509PublicKey,
    }

    rs := pem.EncodeToMemory(&publicBlock)
    return rs, nil
}