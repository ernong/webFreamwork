package util

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	OK                  = 0
	Success             = "SUCCESS"
	MaxIdleConns        = 10
	MaxIdleConnsPerHost = 10
	IdleConnTimeout     = 10
)

// SendRequestWithHeader 发送请求
func SendRequestWithHeader(urlStr, method, queryParam, sign string, postData io.Reader) (buffer *bytes.Buffer, err error) {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   3 * time.Second,
				KeepAlive: 10 * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
		},
		Timeout: 10 * time.Second,
	}
	//reqQueryParam := ""
	URL := urlStr
	Method := method
	if len(Method) == 0 {
		Method = "GET"
	}
	if len(queryParam) != 0 {
		//reqQueryParam = queryParam
		queryParam = strings.Replace(queryParam, " ", "%20", -1) //转义请求参数中所有的空格
		queryParam = strings.Replace(queryParam, "/", "%2F", -1) //转义请求参数中所有的斜线
		queryParam = strings.Replace(queryParam, "(", "%28", -1) //转义请求参数中所有的左括号
		queryParam = strings.Replace(queryParam, ")", "%29", -1) //转义请求参数中所有的右括号
		queryParam = strings.Replace(queryParam, ",", "%2C", -1) //转义请求参数中所有的逗号
		queryParam = strings.Replace(queryParam, ";", "%3B", -1) //转义请求参数中所有的分号
		l, err := url.Parse("?" + queryParam)
		if err != nil {
			//log.Errorf(err, "parse http url error")
			return nil, err
		}
		param := l.Query().Encode()
		URL = fmt.Sprintf("%s?%s", URL, param)
		//log.Debugf("req to nubia url:[%v]", URL)
	}
	reqBuffer := &bytes.Buffer{}
	if postData != nil {
		reqBuffer.ReadFrom(postData)
	}
	//reqData := reqBuffer.Bytes()
	req, err := http.NewRequest(Method, URL, reqBuffer)
	if nil != err {
		return nil, err
	}
	if sign != "" {
		token := fmt.Sprintf("Backend_%v", sign)
		req.Header.Add("Authorization", token)
	}

	/*
		reqMsg := ""

		//log.Debugf("%v request method:[%v], URL:%v, queryParameter:[%v]", b.name, method, urlStr, queryParam)
		if strings.Compare(Method, "GET") == 0 && len(queryParam) != 0 { //记录get请求参数
			reqMsg = fmt.Sprintf("[%v][%s]", time.Now().Format("20060102150405.000000"), reqQueryParam)
		} else {
			reqMsg = fmt.Sprintf("[%v][%s]", time.Now().Format("20060102150405.000000"), reqData)
		}
		log.Debugf(" %v", reqMsg)
	*/

	resp, err := client.Do(req)
	if nil != err {
		//log.Errorf(err, "request failed:%v")
		return nil, err
	}

	// DEBUG("req:%v", req)

	var data io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		data, err = gzip.NewReader(resp.Body)
		if nil != err {
			return nil, err
		}
	default:
		data = resp.Body
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	buffer = &bytes.Buffer{}
	_, err = buffer.ReadFrom(data)
	if nil != err {
		return nil, err
	}

	// log.Debugf("Response status:[%v], body read size:[%v]", resp.StatusCode, n)
	// log.Debugf("Request:[%s],Response body:[%s]", URL, buffer.Bytes())

	return buffer, nil
}

// SendRequest 发送请求
func SendRequest(urlStr, method, queryParam string, postData io.Reader) (buffer *bytes.Buffer, err error) {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   3 * time.Second,
				KeepAlive: 10 * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
		},
		Timeout: 10 * time.Second,
	}
	//reqQueryParam := ""
	URL := urlStr
	Method := method
	if len(Method) == 0 {
		Method = "GET"
	}
	if len(queryParam) != 0 {
		//reqQueryParam = queryParam
		queryParam = strings.Replace(queryParam, " ", "%20", -1) //转义请求参数中所有的空格
		queryParam = strings.Replace(queryParam, "/", "%2F", -1) //转义请求参数中所有的斜线
		queryParam = strings.Replace(queryParam, "(", "%28", -1) //转义请求参数中所有的左括号
		queryParam = strings.Replace(queryParam, ")", "%29", -1) //转义请求参数中所有的右括号
		queryParam = strings.Replace(queryParam, ",", "%2C", -1) //转义请求参数中所有的逗号
		queryParam = strings.Replace(queryParam, ";", "%3B", -1) //转义请求参数中所有的分号
		l, err := url.Parse("?" + queryParam)
		if err != nil {
			//log.Errorf(err, "parse http url error")
			return nil, err
		}
		param := l.Query().Encode()
		URL = fmt.Sprintf("%s?%s", URL, param)
		//log.Debugf("req to nubia url:[%v]", URL)
	}
	reqBuffer := &bytes.Buffer{}
	if postData != nil {
		reqBuffer.ReadFrom(postData)
	}
	//reqData := reqBuffer.Bytes()
	req, err := http.NewRequest(Method, URL, reqBuffer)
	if nil != err {
		return nil, err
	}
	/*
		reqMsg := ""

		//log.Debugf("%v request method:[%v], URL:%v, queryParameter:[%v]", b.name, method, urlStr, queryParam)
		if strings.Compare(Method, "GET") == 0 && len(queryParam) != 0 { //记录get请求参数
			reqMsg = fmt.Sprintf("[%v][%s]", time.Now().Format("20060102150405.000000"), reqQueryParam)
		} else {
			reqMsg = fmt.Sprintf("[%v][%s]", time.Now().Format("20060102150405.000000"), reqData)
		}
		log.Debugf(" %v", reqMsg)
	*/

	resp, err := client.Do(req)
	if nil != err {
		//log.Errorf(err, "request failed:%v")
		return nil, err
	}

	// DEBUG("req:%v", req)

	var data io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		data, err = gzip.NewReader(resp.Body)
		if nil != err {
			return nil, err
		}
	default:
		data = resp.Body
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	buffer = &bytes.Buffer{}
	_, err = buffer.ReadFrom(data)
	if nil != err {
		return nil, err
	}

	// log.Debugf("Response status:[%v], body read size:[%v]", resp.StatusCode, n)
	// log.Debugf("Request:[%s],Response body:[%s]", URL, buffer.Bytes())

	return buffer, nil
}

// ParsingJSONFromString 解析json结构：eg：
// {"BitTorrent":0,"Bithumb":0,"HuobiToken":0,"IPFS":0,"James":0,"MacCoin":0,"NBACoin":0,"Skypeople":0,"TRXTestCoin":0,"binance":0,"ofoBike":0}
func ParsingJSONFromString(jstr string) map[string]int64 {
	if jstr == "" {
		return nil
	}
	jsonMap := make(map[string]int64, 0)
	jsonStr := jstr[1 : len(jstr)-1] //去除前后{}
	for _, param := range strings.Split(jsonStr, ",") {
		if param != "" {
			for key, value := range strings.Split(param, ":") {
				if key == 0 {
					value = strings.Replace(value, "\"", "", -1)
					value = strings.Replace(value, "'", "", -1)
					jsonMap[value] = 0
				} else {
					ValueInt, err := strconv.ParseInt(value, 10, 64)
					if err != nil {
						ValueInt = 0
					}
					if ValueInt > 0 {
						jsonMap[value] = ValueInt
					}
				}

			}
		}
	}
	return jsonMap
}

// ToJSONStr ...
func ToJSONStr(val interface{}) string {
	if nil == val {
		return ""
	}
	real := reflect.ValueOf(val)
	if real.IsNil() {
		return ""
	}
	if real.Kind() == reflect.Ptr && !real.Elem().IsValid() {
		return ""
	}
	if (real.Kind() == reflect.Slice || real.Kind() == reflect.Array || real.Kind() == reflect.Map) && real.IsNil() {
		fmt.Printf("list:%#v\n", real)
		return ""
	}
	data, err := json.Marshal(val)
	if nil != err {
		return fmt.Sprintf("%#v", val)
	}
	return string(data)
}
