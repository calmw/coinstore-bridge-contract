package signature

import (
	"crypto/tls"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type SigDataPost struct {
	FromAddress string `json:"fromAddress"`
	TxData      string `json:"txData"`
	TaskID      int    `json:"taskId"`
	ChainID     int    `json:"chainId"`
	Fingerprint string `json:"fingerprint"`
}

type MachineResp struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func GetSignedRlpData(url string, data SigDataPost) ([]byte, error) {
	b, err := os.ReadFile("/tmp/sig1.pem")
	if err != nil {
		log.Fatal(err)
	}
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
	c, _ := tls.X509KeyPair(bytes, pkey)
	cfg := &tls.Config{
		Certificates:       []tls.Certificate{c},
		InsecureSkipVerify: true,
	}
	tr := &http.Transport{
		TLSClientConfig: cfg,
	}
	client := &http.Client{Transport: tr}
	postData := fmt.Sprintf(`{"fromAddress":"%s","txData":"%s","taskId":%d,"chainId":%d,"fingerprint":"%s"}`,
		data.FromAddress, data.TxData, data.TaskID, data.ChainID, data.Fingerprint,
	)
	fmt.Println("post data:")
	fmt.Println(postData)
	//strings.NewReader(postData)
	d := strings.NewReader(postData)
	request, _ := http.NewRequest("POST", url, d)
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	} else {
		resData, _ := io.ReadAll(resp.Body)
		//fmt.Println("post response:")
		//fmt.Println(string(resData))
		return resData, nil
	}
}

func RandInt(min, max int) int {
	rand.NewSource(time.Now().UnixNano())
	randomNum := rand.Intn(max-min+1) + min
	return randomNum
}
