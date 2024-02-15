package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

func getNonce(length int) string {
	var text = ""
	const possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	for i := 0; i < length; i++ {
		text += string(possible[rand.Intn(62-1)+1])
	}
	return text
}

type OauthData struct {
	Account_type           string `form:"account_type" json:"account_type"`
	Accountid              string `form:"accountid" json:"accountid"`
	Baseurl                string `form:"baseurl" json:"baseurl"`
	Consumer_secret        string `form:"consumer_secret" json:"consumer_secret"`
	Deploy                 string `form:"deploy" json:"deploy"`
	Httpmethod             string `form:"httpmethod" json:"httpmethod"`
	Is_inactive            string `form:"is_inactive" json:"is_inactive"`
	Is_pettycash           string `form:"is_pettycash" json:"is_pettycash"`
	Keyword_subsidiary     string `form:"keyword_subsidiary" json:"keyword_subsidiary"`
	Oauth_consumer_key     string `form:"oauth_consumer_key" json:"oauth_consumer_key"`
	Oauth_nonce            string `form:"oauth_nonce" json:"oauth_nonce"`
	Oauth_signature_method string `form:"oauth_signature_method" json:"oauth_signature_method"`
	Oauth_timestamp        int64  `form:"oauth_timestamp" json:"oauth_timestamp"`
	Oauth_token            string `form:"oauth_token" json:"oauth_token"`
	Oauth_version          string `form:"oauth_version" json:"oauth_version"`
	Script                 string `form:"script" json:"script"`
	Token_secret           string `form:"token_secret" json:"token_secret"`
}

func GetNetsuite(method string) (interface{}, error) {
	client := &http.Client{}

	data := OauthData{
		Account_type:           "Bank",
		Accountid:              "5318962_SB1",
		Baseurl:                "https://5318962-sb1.restlets.api.netsuite.com/app/site/hosting/restlet.nl",
		Consumer_secret:        "ed59bbc069fea27a03c6cc74d15416b3d09bde79aaea3f0b8c575863f544a5eb",
		Deploy:                 "1",
		Httpmethod:             method,
		Is_inactive:            "F",
		Is_pettycash:           "T",
		Keyword_subsidiary:     "1",
		Oauth_consumer_key:     "2197e5f8ff53011b9665e78349418f44bfbadffabe34f534c1ae9468ce48b83c",
		Oauth_nonce:            getNonce(32),
		Oauth_signature_method: "HMAC-SHA256",
		Oauth_timestamp:        time.Now().Unix(),
		Oauth_token:            "33ad33fff1164f91c4171bf4b3ac99c73f3a0f759ad73bf53c55222849907600",
		Oauth_version:          "1.0",
		Script:                 "326",
		Token_secret:           "4c9d928736069aac43074b17c50c0651cc7af11bed0f2e47064035ef60e65466",
	}

	values := reflect.ValueOf(data)
	typeOfS := values.Type()

	var dataOauth = ""
	var keys []string
	for index := 0; index < values.NumField(); index++ {
		key := typeOfS.Field(index).Tag.Get("json")
		keys = append(keys, key)
		if key != "token_secret" && key != "consumer_key" && key != "httpmethod" && key != "baseurl" && key != "consumer_secret" && key != "accountid" {
			if key == "oauth_timestamp" {
				var keys = key + "="
				var value = values.Field(index)
				dataOauth += fmt.Sprint(keys, value, "&")
			} else if key == "script" {
				var value = values.Field(index)
				dataOauth += key + "=" + value.String()
			} else {
				var value = values.Field(index)
				dataOauth += key + "=" + value.String() + "&"
			}
		}
	}

	encodedData := url.QueryEscape(dataOauth)
	encodeBaseUrl := url.QueryEscape(data.Baseurl)
	completeData := data.Httpmethod + "&" + encodeBaseUrl + "&" + encodedData
	hmacsha256Data := hmac.New(sha256.New, []byte(data.Consumer_secret+"&"+data.Token_secret))
	hmacsha256Data.Write([]byte(completeData))
	macSum := hmacsha256Data.Sum(nil)
	base64EncodedData := base64.StdEncoding.EncodeToString(macSum)
	oauth_signature := url.QueryEscape(base64EncodedData)

	timeStamp := fmt.Sprint(data.Oauth_timestamp, `",`, `realm="`+data.Accountid, `"`)
	OAuth := `OAuth oauth_signature="` + oauth_signature + `",` + `oauth_version="` + data.Oauth_version + `",` + `oauth_nonce="` + data.Oauth_nonce + `",` + `oauth_signature_method="` + data.Oauth_signature_method + `",` + `oauth_consumer_key="` + data.Oauth_consumer_key + `",` + `oauth_token="` + data.Oauth_token + `",` + `oauth_timestamp="` + timeStamp

	req, err := http.NewRequest(http.MethodGet, "https://5318962-sb1.restlets.api.netsuite.com/app/site/hosting/restlet.nl", nil)

	if err != nil {
		return "request failed", err
	}

	req.Header.Set("Authorization", OAuth)
	req.Header.Set("Content-Type", "application/json")

	query := req.URL.Query()
	query.Add("script", "326")
	query.Add("deploy", "1")
	query.Add("is_inactive", "F")
	query.Add("is_pettycash", "T")
	query.Add("account_type", "Bank")
	query.Add("keyword_subsidiary", "1")
	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)

	if err != nil {
		return "request failed", err
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return responseData, err
	}
	// Always close the response body
	defer res.Body.Close()

	return string(responseData), err
}
