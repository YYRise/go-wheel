package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	rsa2 "github.com/wumansgy/goEncrypt/rsa"
	"testing"

	"github.com/golang-module/dongle"
)

const (
	publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1O3p0JN0/RrP7eY3f81i
zPf16FS0WMNGCJkd+y5c6yBzUvN0IEeoxiIWIBhoMKH0pzlzBg0rfttojSodOgNo
m/UCAzAYEgdIsNee5LSN/7e0T2/QvsIAHINuA8gI8fGoGiSA2TEzpUo6aVXwhZT3
4GGRdrSJ+m4iVk/Kt95tavBNk+NDVSeb5xAjxBchT5BjAMMlE0ffGZb0MMjjO5+e
9Tn8f99M2VMqpzXHXZzv1ABmqufzS20iWcSvnjhWcJ9hiKwO8Z30GgJyACmml+HM
xLYEFN9h2MWYgxLm9Z0rLMrWwMM+E2rCs8tsxAD5sO9RZMJPl1C0FIsMR53ngqbz
owIDAQAB
-----END PUBLIC KEY-----`

	privateKey1 = `-----BEGIN RSA PRIVATE KEY-----
MIIEpgIBAAKCAQEA1O3p0JN0/RrP7eY3f81izPf16FS0WMNGCJkd+y5c6yBzUvN0
IEeoxiIWIBhoMKH0pzlzBg0rfttojSodOgNom/UCAzAYEgdIsNee5LSN/7e0T2/Q
vsIAHINuA8gI8fGoGiSA2TEzpUo6aVXwhZT34GGRdrSJ+m4iVk/Kt95tavBNk+ND
VSeb5xAjxBchT5BjAMMlE0ffGZb0MMjjO5+e9Tn8f99M2VMqpzXHXZzv1ABmqufz
S20iWcSvnjhWcJ9hiKwO8Z30GgJyACmml+HMxLYEFN9h2MWYgxLm9Z0rLMrWwMM+
E2rCs8tsxAD5sO9RZMJPl1C0FIsMR53ngqbzowIDAQABAoIBAQCO1RE1ItUlO6kj
Un0ENAgEqojAUqGvsT33Yo7kAZO+/cOeb0UEqk0iq5bf7L9ncBynWDg6ZPc6X3/g
wdFdKxAvHck9zjM3VL+EMP+bNyrR0K8ZYk5Kx+Q/PEK+Mp8dfRdgggAUsZaNWB+a
rVVspiMo1wo28KBl5x8NevTnJkOLqXAyB7UyLWqnOL1fb988lZvZPR7ZUYroVIZa
pyXtZcafIJeKyQ3bvWI5+eFqOe61Z4Bx1+TpfZ3fKfSDW0vhxzNqaimOa8jSXtMJ
jMeOctL4nZ0TPo/jS3I+XlaH4ZQlFLuUWGscpxwfEeBN23I8HRLkZXJsw66yvRN3
s4bUKPXRAoGBAP/3oSZAECvfsYYzs76tnrAmR/0GxCqgguxDlWn5DowQzdWFOdHC
ZbTo/hUVoMSQnO1EKCFlnBS+wg/3TuIzUO0ewC1aeT7qHbOMDl0zKbNpS2Z9/j+U
zro+qz7XmkWolMCfmDrCrw9CtCxcMSII+ajbI8SAgFVMz9XnDt+xW9E9AoGBANT0
4F6kCUJTEyqf2+v84tjQ2wGIF6XtZPU9JR806zeMyahQ9F6z3hY8BYb0tIy5b3uJ
VlJ9TG1qg/t59TWxIq43mYSUJHe0aJi3ilooObQtHlhPu8nwmmX47sX0PyG2hMoD
kBVxTpTDmBaDz7O9uBnlMXJN5qEygctaixpEbmZfAoGBAMBA9kEMjRjnAyeRXcgy
D6aumhNqKZz6wltCx864yjxZwsBFOJBcOpgPCAg+HmqFU9jCAIJVF05dmNT1I8Ky
WG5BUoa+FaMzpOtenstRylh/Far9pyGKW1t4BpdEyRLY9CFZvbUk1OfZagqHlD/E
DgDN16eX/MwUzWYUDg/l3tjhAoGBAKGip/ZNjVWRFpggs9z/mfK1O7WC5Wgksp9N
ZLK2CN6l9p3RrFmBLk00C4HulGfHi+15RVLhFbRqx3iFje/N3iPbwaMWikNtZIKd
tN5Pb9To9gJTqpZRD+/cLOeFRrHBBjMK1z7fPKS/fN2B+JFVq7nD827t3+J0In4F
4FT0odMDAoGBAJk3ELB/FHY8xzZ4jF1wG/a1CK681Xm6SuU5KIELDSAUNoou6OPG
mS8gU20MMPAeV2z7khyDcSxlHsUyL73eLeaakbQov9NMW7cc99XX4wnP4W7FRpmr
QbHmKuHIRFHCFv+XX8c0aK2mDZMUlzJdy4FgD/YCEZ7kZMZKyvZW/ZuV
-----END RSA PRIVATE KEY-----`
)

func TestEncrypt(t *testing.T) {
	pubKey := RsaPri.GetPublicKey()

	str, err := NewPublic(pubKey).Encrypt("123456")
	if err != nil {
		t.Error("rsa public encrypt error", err)
		return
	}

	t.Log(str)
}

func TestEncrypt2(t *testing.T) {
	pubKey := "-----BEGIN RSA PUBLIC KEY-----\nMIIBCgKCAQEAuQh5m82SBwsM30udH5Ov1rMnm/9MTqiN/oBBIPgr4I2/QRuFHGWw\nmwJKs3+P3V/QyFTEi7gIANfF9edyMDJnrRyHiYz+5zrpak+3vK1TF59UnLxeAy6o\naRFW8DX1gRq9Y/CvatOehCSrBILB0iFL8X+0GMWOIRQYbKA1Bg+Jfy5K7BZSuYOw\n8TbbvEdwBOV2TlaNhEfveC6crd2fxyhSyEOVyrQxJ+z13fJOqlBD0QbBmA8oCE5c\nXJZb9vadY64HCaNOnhn6r8mXLxN7tNw5h4b7ygVMVyOtZDGO8GbkSFGMU62DJY3Z\ndPgockNdl6onPa3kqRgRpJR6bsmakTVsBwIDAQAB\n-----END RSA PUBLIC KEY-----\n"
	str, err := NewPublic(pubKey).Encrypt("IgkibX71IEf382PT")
	if err != nil {
		t.Error("rsa public encrypt error", err)
		return
	}

	t.Log(str)
}

func TestDecrypt(t *testing.T) {
	//decryptStr := "KTKXckjkCLI6Vk_y_XROnY-a6nJpllruL-CX-v_2AFxfghA2tZ2nkQyS6d1-IIYMlgwm4ivwlzy-phLtaN9BB03htA5D9rwjA_JwYtqAG4iwuvgaDl2SiZ_H2ACv-aV1kNRgnyjh14hs0JiSt5VHEiJ3D2xYzOCKwtEzoo0WczJ-MYb3u_-bfcnm9YtvgtG5-y3Jy7WYr-IwXdBKqPO0E-jzrtY7m3Q1yC4znHdzjNpxCj0I6YRx4CZ362b706qNX7sl3E5KTJeSmYrsurB-SxQT1CaqGzVt7mshx1v2qGnv5NBNXpj7ZPKWGJbgaCUxcuxd1Mg0o81HnfbsGuSlFQ=="
	pubKey := RsaPri.GetPublicKey()

	fmt.Println("pubKey:", pubKey)
	decryptStr, err := NewPublic(pubKey).Encrypt("123456")
	if err != nil {
		t.Error("rsa public encrypt error", err)
		return
	}
	fmt.Println("decryptStr:", decryptStr)
	str, err := RsaPri.Decrypt(decryptStr)
	if err != nil {
		t.Error("rsa private decrypt error", err)
		return
	}
	fmt.Println("str:", str)
	t.Log(str)
}

func BenchmarkEncryptAndDecrypt(b *testing.B) {
	pubKey := RsaPri.GetPublicKey()
	b.ResetTimer()
	rsaPublic := NewPublic(pubKey)
	rsaPrivate := RsaPri
	for i := 0; i < b.N; i++ {
		encryptString, _ := rsaPublic.Encrypt("123456")
		rsaPrivate.Decrypt(encryptString)
	}
}

func TestNewPublic(t *testing.T) {
	// 生成一个2048位的RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}

	// 创建公钥和私钥的x509编码
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	})

	fmt.Println("Private Key:", string(privateKeyPEM))
	fmt.Println("Public Key:", string(publicKeyPEM))
	// Base64编码密钥
	privateKeyBase64 := base64.StdEncoding.EncodeToString(privateKeyPEM)
	publicKeyBase64 := base64.StdEncoding.EncodeToString(publicKeyPEM)

	// 输出公私钥字符串
	fmt.Println("Private Key (Base64 Encoded):", privateKeyBase64)
	fmt.Println("Public Key (Base64 Encoded):", publicKeyBase64)
}

func TestBase64(t *testing.T) {
	strBase64 := base64.StdEncoding.EncodeToString([]byte(PriKeyStr))
	fmt.Println(strBase64)
}

func TestRsa(t *testing.T) {

}

func TestDongle(t *testing.T) {
	s := "eHRqbFdUc3RkYm9aVlFGamxYN2hUQkJvTnIxTVRGMkl3SDhVQ1BNb1BuTldIdTBNMVRFTmdVanFCOWZoVHhubXlGZjJRVzNrTUFRRjRDU0NONGpQWjBaQVFjemVPOXVrOE1Da0dKQmtjTFpHRUFCRlQ0aWlYZGZ4bXBSMWRIRkpSR0ZKUURXN1B1cVpkbm9QelBJVmJTNmpSWTZyRWU0MURudFlpdHhiZkh4VHl5dTBNUDEvN0I3bDJ3RHllWGxSYXIwUTBjbUpTeDc4TjFKeTdPSW9iaU0zRGFGdFFQUnhzcDRrTU1ybURRR09aZTVOOTgvU3ZNbmhFRmNsd3pDRktaRjYwWWVyVVRSenlSdTIrbURmelJrZEQ3amowbUwrU2lUZHZ5dml6MTR0ZHJhS3FxVEtNczJSY3F2Ulo3MTAwWDRYakFzNDlUVS8xRnVZSlFkY2NRPT0="
	dec := dongle.Decrypt.FromBase64String(s).ByRsa(PriKeyStr)
	fmt.Println("----err:", dec.Error)
	fmt.Println("----string:", dec.String())
}

func TestGoEncrypt(t *testing.T) {
	s := "eHRqbFdUc3RkYm9aVlFGamxYN2hUQkJvTnIxTVRGMkl3SDhVQ1BNb1BuTldIdTBNMVRFTmdVanFCOWZoVHhubXlGZjJRVzNrTUFRRjRDU0NONGpQWjBaQVFjemVPOXVrOE1Da0dKQmtjTFpHRUFCRlQ0aWlYZGZ4bXBSMWRIRkpSR0ZKUURXN1B1cVpkbm9QelBJVmJTNmpSWTZyRWU0MURudFlpdHhiZkh4VHl5dTBNUDEvN0I3bDJ3RHllWGxSYXIwUTBjbUpTeDc4TjFKeTdPSW9iaU0zRGFGdFFQUnhzcDRrTU1ybURRR09aZTVOOTgvU3ZNbmhFRmNsd3pDRktaRjYwWWVyVVRSenlSdTIrbURmelJrZEQ3amowbUwrU2lUZHZ5dml6MTR0ZHJhS3FxVEtNczJSY3F2Ulo3MTAwWDRYakFzNDlUVS8xRnVZSlFkY2NRPT0="

	//ciphertext, err := base64.StdEncoding.DecodeString(s)
	//if err != nil {
	//	return
	//}
	txt, err := RsaPri.Decrypt(s)
	fmt.Println("--明文--：", txt, err)

	priBase64 := base64.StdEncoding.EncodeToString([]byte(PriKeyStr))
	fmt.Println("priBase64:", priBase64)
	plaintext, err := rsa2.RsaDecryptByBase64(s, priBase64)
	fmt.Println("plaintext:", string(plaintext))
	fmt.Println("err:", err)

	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("明文：", string(plaintext)) // test
}

func TestGoEncrypt2(t *testing.T) {
	var privateKey = []byte(`-----BEGIN  WUMAN RSA PRIVATE KEY -----
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
`)
	//s := `mFWBdT4Y70ZNEQ7PVIFKwbkefufu52WGXYLrW0Vk1XuajrrEE54dqj4VK2yuGIeMq5bHKAdkDnACB2ABzHLQuobTDpkS0Nj5AlJvwbRDV3pOCB1x0q3aqEooTppeMs8P/WG3YCRDTQPWgZISPsFBQVT1tk77BiImcY4SZM9IL0B4TFUKS9sShnjAebxmJkj8jfYYh7gNzUY0YMvOV6HuiT5C0RsbTe1jwMyN87QEwvpvuPelkeQ8LX1AG+qsn2q4TvOYEKCNfNnePjMIQ/5MlesledwiqUpc/YtY3qj4Qx+8b5luaQ6kyu+zyOXV/A0XjjxIxqLWKU8eAl7eA3o72Q==`

	s := `aTIsL4hKtpQm8ECDwZAazZfql5G/7ThTR+II3mtQxJxfvuCpb1EhSh5uTciMABOQmejirZ9GyTjlShNlgI5YAHK20K2gEK2qigH0VfSTVINSq0dB8QZBvAO/OVR3ZSIBz+zlIIw61I+6ZMpnsdf3fZctnYq7sc9Xvhs2vNZMPkrH16KzXdUiCMnsioIdX5yPbkJ4H4QWyL4tWcNN4dGLdG8TT7AqR1pFWdgL6Gs/V0VQ9z5sPSi4ltQwlLlGI6iTSSK/ZQR1lLFpfog1s9HgEURHG3xdLg+FBVKGXE5AIKY5eTDXEyoBF4v4fP8cn8AbS2P+xEB+iaDoLk6SWYwKbQ==`
	// pem 解码
	block, _ := pem.Decode(privateKey)

	// X509 解码
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		panic(fmt.Sprintf("生成密钥对失败:%s", err))
	}
	r := NewPrivate(priKey)
	txt, err := r.Decrypt(s)
	fmt.Println("--明文--：", txt, err)

	priBase64 := base64.StdEncoding.EncodeToString(privateKey)
	//fmt.Println("priBase64:", priBase64)
	plaintext, err := rsa2.RsaDecryptByBase64(s, priBase64)
	fmt.Println("plaintext:", string(plaintext))
	fmt.Println("err:", err)
	if err != nil {
		return
	}
	fmt.Println("明文：", string(plaintext)) // test
}
