package vivo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

type VivoClient struct {
	accessKey    string
	accessSecret string
	client       *http.Client
}

func NewVivoClient(accessKey, accessSecret string) *VivoClient {
	return &VivoClient{
		accessKey:    accessKey,
		accessSecret: accessSecret,
		client:       &http.Client{},
	}
}

func ParamsToSortQuery(params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var resultList []string
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		value := params[key]
		resultList = append(resultList, fmt.Sprintf("%s=%s", key, value))
	}

	result := strings.Join(resultList, "&")
	return result
}

// Signature 签名
func Signature(params map[string]string, accessSecret string) string {
	query := ParamsToSortQuery(params)
	sign := HmacSha256(query, accessSecret)
	return sign
}

// HmacSha256 加密
func HmacSha256(stringToSign string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(stringToSign))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

// Post POST请求基础模版
func (c *VivoClient) Post(reqBody string) ([]byte, error) {
	host := "https://developer-api.vivo.com.cn/router/rest"
	req, err := http.NewRequest("POST", host, strings.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Charset", "UTF-8")
	body, err := c.FormHttp(req)
	return body, err
}

func (c *VivoClient) FormHttp(req *http.Request) ([]byte, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *VivoClient) NewBaseParams() map[string]string {
	p := map[string]string{
		"access_key": c.accessKey,
		"timestamp":  strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
	}

	return p

}
func HandleParams(params interface{}, p map[string]string, accessSecret string) (map[string]string, error) {
	paramsByte, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	var mi map[string]interface{}
	err = json.Unmarshal(paramsByte, &mi)
	if err != nil {
		return nil, err
	}

	for k, v := range mi {
		vt := reflect.TypeOf(v)
		switch vt.Kind() {
		case reflect.Map, reflect.Array, reflect.Slice:
			value, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			p[k] = string(value)
		default:
			p[k] = v.(string)
		}
	}

	sign := Signature(p, accessSecret)
	p["sign"] = sign

	return p, err
}
