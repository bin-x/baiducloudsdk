package baiducloudsdk

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strings"
	"time"
)

type HttpClient struct {
	ApiKey string
	SecretKey string
	accessToken AccessToken
}

type AccessToken struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
	ExpiresAt int64 `json:"expires_at"`
}

const accessTokenUrl = "https://aip.baidubce.com/oauth/2.0/token"
const authPath = "./baidu-accesstoken"

func NewHttpClient(ApiKey string, SecretKey string) *HttpClient{
	return &HttpClient{
		ApiKey: ApiKey,
		SecretKey: SecretKey,
	}
}

func buildUrl(url, query string) string {
	if query == ""{
		return url
	}
	if strings.Contains(url,"?") {
		return url+"&"+query
	}
	return url+"?"+query
}


func (client *HttpClient) createAccessToken() error {
	now:=time.Now().Unix()
	if client.accessToken.AccessToken != "" && client.accessToken.ExpiresAt > now  {
		return nil
	}

	data,err:=ioutil.ReadFile(authPath)
	token := AccessToken{}
	if err ==nil {
		json.Unmarshal(data,&token)
		if token.ExpiresAt > now {
			client.accessToken = token
			return nil
		}
	}

	urlValue := url2.Values{
		"grant_type":{"client_credentials"},
		"client_id":{client.ApiKey},
		"client_secret":{client.SecretKey},
	}
	url := buildUrl(accessTokenUrl,urlValue.Encode())
	r,_ := http.NewRequest("GET",url,nil)
	c:=&http.Client{}
	response,err:=c.Do(r)
	if err!=nil{
		return err
	}
	body,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		return err
	}

	json.Unmarshal(body,&token)
	if token.Error != "" || token.ErrorCode != "" {
		return errors.New(token.Error+token.ErrorCode)
	}

	token.ExpiresAt = now + token.ExpiresIn
	client.accessToken = token
	data ,_ = json.Marshal(token)

	ioutil.WriteFile(authPath,data,0644)
	return nil
}

func (client *HttpClient) Request(url string, data url2.Values, headers http.Header) ([]byte, error) {
	if err:=client.createAccessToken();err!=nil{
		return nil,err
	}

	query := url2.Values{
		"access_token":{client.accessToken.AccessToken},
	}
	url = buildUrl(url,query.Encode())

	reqBody := strings.NewReader(data.Encode())
	r,_ := http.NewRequest("POST",url,reqBody)
	r.Header = headers
	c:=&http.Client{}
	response,err:=c.Do(r)
	if err!=nil{
		return nil,err
	}
	respBody,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		return nil,err
	}
	return respBody,nil
}