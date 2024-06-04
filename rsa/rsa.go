package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

const (
	PubKeyStr = `-----BEGIN  WUMAN  RSA PUBLIC KEY -----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA29fttcEDUvhGJEhQXEIH
8blptZRF5itec4GtEGtkSr4Wjmsf2o2XKOr6YEbTOeDA/DdnDSbVzK2ZUscqyBxb
KGwI/Bpv9l5K/sh9+Oj2Y8YH53+XkqRSGvmhHqolhb+gcfH+FKG5IflGuiOREs4h
02TVmPAFPTmZjYBeVexJgmPodGPOe36QVnMeOG8tHOFxItkMvJUpilzs85xdHqTT
jWCtk/SjHrp5NGSkHSmionOtrFiksS/gTX0EzrptmAGHTjZV0NX7Nu8Ma45rVdMR
wXrDPbk0yR0iFdBEZ1ceGsNg2VjrZ3LCZi3zO+ieA7sBjHARHai5MuFlh9KJ8+Yk
wwIDAQAB
-----END  WUMAN  RSA PUBLIC KEY -----`

	PriKeyStr = `-----BEGIN  WUMAN RSA PRIVATE KEY -----
MIIEpgIBAAKCAQEA29fttcEDUvhGJEhQXEIH8blptZRF5itec4GtEGtkSr4Wjmsf
2o2XKOr6YEbTOeDA/DdnDSbVzK2ZUscqyBxbKGwI/Bpv9l5K/sh9+Oj2Y8YH53+X
kqRSGvmhHqolhb+gcfH+FKG5IflGuiOREs4h02TVmPAFPTmZjYBeVexJgmPodGPO
e36QVnMeOG8tHOFxItkMvJUpilzs85xdHqTTjWCtk/SjHrp5NGSkHSmionOtrFik
sS/gTX0EzrptmAGHTjZV0NX7Nu8Ma45rVdMRwXrDPbk0yR0iFdBEZ1ceGsNg2Vjr
Z3LCZi3zO+ieA7sBjHARHai5MuFlh9KJ8+YkwwIDAQABAoIBAQCyBCtMXbqfWMMT
ZisMSbu9FPJwQlxHgR6+UWceQJe5nisNr9jfVH/udje/9hncaA5dLU+Y6rV9Q6U/
zl7qI2v9U14DJjU7PidkIF1BTQMWz6he4IaQC9cgWLsK5aP0pbL6EYY4lqwewodu
+pXisF/bmW8MpG7ZoOaiGixJT0hG97aS3YD506RqdnK4a9yG5ycoVUZTzFTIM+aq
MCTOWJLbHJBnY52v7Rqaor1jZ0o+C/Cykbts25VHWZ9ygxBfI/S75jTg/zibpqbo
TnICGBqzzfAsThHYiji3ZgEF2bSadxQZ956Rvm5Dlhk0A6ylwKl1gJTAxyZNRjx2
zogLHNFhAoGBAP3TmQSpyYtjr8xdD+AsinMi3q3p/6+FPBh5wrV9Dvd5aSrKCYnU
j/9LOYmCP6JfKc90K3L0PKTfS1LDZXHoeYpzxB0uD2iTIkYQKraUp7rzqLBMkdTB
nTuOByqnTB45WlWlfy6m/8CO+0r5DkrO1fU8gWkg3+bUQoUTPOhuFkqdAoGBAN25
1oVmFyDvxw62WZ95SPktqzmErFH/X+7wpGInUb0vxnGJ0SP2bVgMLyL9Ecz8p5hC
ZEZVbo/WFs17hqz8Z6Xx4Gq7aoredpsZswpApdODs5sQf72LXj4Cuf1iwMGIyQxa
8hElV9uRIeKfEG2E/HBLLPO8qhGBCsWIT1Jms97fAoGBAPQFIf2usTkVXCPvb9zH
VU79PfEanhoCz9SD8mGCWgomqaleVK8yMEFx8120XzLdpBdyCndYQJkMpqBpgzRw
F7C4PNkEuAGEOhX7YuTmox4DM7BR3H0aqetgTpl9/pqr7qGaGlwiZoubqhDYwRnA
IUfDpHIKDdcfRtgit5KIi1utAoGBAKqoAK8IFsEpDGMMgwq1hS8UsXdB4If0MNht
q3hInycoAGsfEjPF1f8w0Y7yjaLiy/PrFdb0pnZa544ch1nZo8Ub2AkOW0CrXUqf
iyhW/ctA0RqGpmszO8QqwRB/07CiIWw7C5maznaWzCfrGe/RraKYme63xYZXdfz3
n2Xi2oqtAoGBAPB2SL1UafzZiC70e+NeDeaCLTCorCvIrN73RngYXU1OKUynrcjL
xOs8yVFs5cK0sNwEkiozcldfOfU2j70tyrGz+txi5+Db6ex5VXmEKSqdZXRDtqqc
KKRFfpuebqdaR50SDa4Lq6JbqYtwg0CLZeju4Mq41i9F2p4myqVDUsZs
-----END  WUMAN RSA PRIVATE KEY -----
`
)

var _ Public = (*rsaPub)(nil)
var _ Private = (*rsaPri)(nil)

var (
	RsaPri Private
)

/*
	func init() {
		// 生成RSA密钥对
		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			panic(fmt.Sprintf("生成密钥对失败:%s", err))
		}
		RsaPri = NewPrivate(privateKey)
		//RsaPub = NewPublic(string(pem.EncodeToMemory(&pem.Block{
	}
*/

func init() {
	// pem 解码
	block, _ := pem.Decode([]byte(PriKeyStr))

	// X509 解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(fmt.Sprintf("生成密钥对失败:%s", err))
	}
	RsaPri = NewPrivate(privateKey)
}

type Public interface {
	// Encrypt 加密
	Encrypt(encryptStr string) (string, error)
}

type Private interface {
	// Decrypt 解密
	Decrypt(decryptStr string) (string, error)

	GetPublicKey() string
}

type rsaPub struct {
	PublicKey string
}

func NewPublic(publicKey string) Public {
	return &rsaPub{
		PublicKey: publicKey,
	}
}

func (pub *rsaPub) Encrypt(encryptStr string) (string, error) {
	// pem 解码
	block, _ := pem.Decode([]byte(pub.PublicKey))

	// x509 解码
	//publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	//publicKeyInterface, err := x509.ParsePKIXPublicKey([]byte(encryptStr))
	if err != nil {
		return "", err
	}

	// 类型断言
	//publicKey := publicKeyInterface.(*rsa.PublicKey)

	//对明文进行加密
	encryptedStr, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(encryptStr))
	if err != nil {
		return "", err
	}
	fmt.Println("加密后：", string(encryptedStr))
	//返回密文
	return base64.StdEncoding.EncodeToString(encryptedStr), nil
}

type rsaPri struct {
	//PrivateKey string
	PrivateKey   *rsa.PrivateKey
	PublicKeyPEM string
}

func NewPrivate(privateKey *rsa.PrivateKey) Private {
	pri := &rsaPri{
		PrivateKey: privateKey,
	}
	pri.initPublicKey()
	return pri
}

func (pri *rsaPri) initPublicKey() {
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&(pri.PrivateKey.PublicKey)),
	})
	pri.PublicKeyPEM = string(publicKeyPEM)
}

func (pri *rsaPri) Decrypt(decryptStr string) (string, error) {
	//// pem 解码
	//block, _ := pem.Decode([]byte(pri.PrivateKey))
	//
	//// X509 解码
	//privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	//if err != nil {
	//	return "", err
	//}
	decryptBytes, err := base64.StdEncoding.DecodeString(decryptStr)
	if err != nil {
		return "", err
	}
	//对密文进行解密
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, pri.PrivateKey, decryptBytes)
	if err != nil {
		return "", err
	}
	//返回明文
	return string(decrypted), nil
}

func (pri *rsaPri) GetPublicKey() string {
	return pri.PublicKeyPEM
}
