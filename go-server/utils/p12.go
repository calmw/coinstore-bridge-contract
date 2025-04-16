package utils

import (
	"crypto/tls"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/// https://www.cnblogs.com/paulwhw/p/14015824.html

// SigDataPost
// {"fromAddress":"TTgY73yj5vzGM2HGHhVt7AR7avMW4jUx6n","txData":"c914f2fa2a214bf1c6bf80de09ecda76b9e7bc379c16b33e484f5ccf47e92b0a","taskId":2,"chainId":728126428,"fingerprint":"00b3a7749898d575ea92577f0f3ed3689355ea3dc276076fd2f716ead07d15b6"}
type SigDataPost struct {
	FromAddress string `json:"fromAddress"`
	TxData      string `json:"txData"`
	TaskID      int    `json:"taskId"`
	ChainID     int    `json:"chainId"`
	Fingerprint string `json:"fingerprint"`
}

func RequestWithPem(url string, data SigDataPost) ([]byte, error) {
	b, _ := os.ReadFile("./sig.pem")
	pem.Decode(b)
	var pemBlocks []*pem.Block
	var v *pem.Block
	var pkey []byte
	for {
		v, b = pem.Decode(b)
		if v == nil {
			break
		}
		if v.Type == "PRIVATE KEY" {
			pkey = pem.EncodeToMemory(v)
		} else {
			pemBlocks = append(pemBlocks, v)
		}
	}

	bytes := pem.EncodeToMemory(pemBlocks[0])
	keyString := string(pkey)
	CertString := string(bytes)
	fmt.Printf("Cert :\n %s \n Key:\n %s \n ", CertString, keyString)
	c, _ := tls.X509KeyPair(bytes, pkey)
	cfg := &tls.Config{
		Certificates: []tls.Certificate{c},
	}
	tr := &http.Transport{
		TLSClientConfig: cfg,
	}
	client := &http.Client{Transport: tr}
	strings.NewReader(fmt.Sprintf(`{"fromAddress":"%s","txData":"%s","taskId":%d,"chainId":%d,"fingerprint":"%s"}`,
		data.FromAddress, data.TxData, data.TaskID, data.ChainID, data.Fingerprint,
	))
	request, _ := http.NewRequest("POST", url, nil)
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	} else {
		data, _ := io.ReadAll(resp.Body)
		return data, nil
	}
}
