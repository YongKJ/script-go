package ApiUtil

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	Global "script-go/src/application/config"
	"script-go/src/application/pojo/dto/Log"
	"script-go/src/application/util/DataUtil"
	"script-go/src/application/util/LogUtil"
	"strings"
	"time"
)

var resTemplate = getResTemplate()

func getResTemplate() *http.Client {
	proxyUrl, err := url.Parse(Global.SocksProxy)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApiUtil", "getResTemplate", "err", err))
	}

	if Global.ProxyEnable {
		return &http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
		}
	}

	return &http.Client{
		Timeout: 60 * time.Second,
	}
}

func RequestWithParamsByGet(api string, params map[string]any) string {
	urlStr := getUrl(api, params)
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApiUtil", "RequestWithParamsByGet", "req", err))
	}

	resp, err := resTemplate.Do(req)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApiUtil", "RequestWithParamsByGet", "resp", err))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApiUtil", "RequestWithParamsByGet", "body", err))
	}

	return string(body)
}

func RequestWithParamsByPost(api string, params map[string]any) string {
	urlStr := getUrl(api, params)
	req, err := http.NewRequest("POST", urlStr, nil)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApiUtil", "RequestWithParamsByPost", "req", err))
	}

	resp, err := resTemplate.Do(req)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApiUtil", "RequestWithParamsByPost", "resp", err))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApiUtil", "RequestWithParamsByPost", "body", err))
	}

	return string(body)
}

func RequestWithUrlByGetAndDownload(url string, taskFunc func(resp *http.Response)) {
	resp, err := http.Get(url)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApiUtil", "RequestWithUrlByGetAndDownload", "resp", err))
	}
	defer resp.Body.Close()
	taskFunc(resp)
}

func RequestWithParamsByGetToBuffer(api string, params map[string]any) []byte {
	resp, err := http.Get(getUrl(api, params))
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApiUtil", "RequestWithUrlByGetAndDownload", "resp", err))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApiUtil", "RequestWithParamsByGetToBuffer", "body", err))
	}

	return body
}

func RequestWithParamsAndJsonBodyByPost(api string, params map[string]any, mapData map[string]any) string {
	jsonData, err := json.Marshal(mapData)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApiUtil", "RequestWithParamsAndJsonBodyByPost", "jsonData", err))
	}

	urlStr := getUrl(api, params)
	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(jsonData))
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApiUtil", "RequestWithParamsAndJsonBodyByPost", "req", err))
	}

	resp, err := resTemplate.Do(req)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApiUtil", "RequestWithParamsAndJsonBodyByPost", "resp", err))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("ApiUtil", "RequestWithParamsAndJsonBodyByPost", "body", err))
	}

	return string(body)
}

func RequestWithParamsByGetToEntity(api string, params map[string]any, class any) any {
	return DataUtil.JsonToObject(RequestWithParamsByGet(api, params), class)
}

func getUrl(api string, params map[string]any) string {
	urlStr := api
	if !strings.HasPrefix(api, "http") {
		urlStr = "http://" + api
	}
	for key, value := range params {
		sep := "?"
		if strings.Contains(urlStr, "?") {
			sep = "&"
		}
		urlStr += sep + url.QueryEscape(key) + "=" + url.QueryEscape(value.(string))
	}
	return urlStr
}
