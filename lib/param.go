package nx

import (
	"fmt"
	"net/url"
)

func Param(params map[string]any, baseURL ...string) string {
	// 设置默认 baseURL
	defaultBaseURL := ""
	finalBaseURL := defaultBaseURL
	if len(baseURL) > 0 && baseURL[0] != "" {
		finalBaseURL = baseURL[0]
	}

	// 初始化 URL 值对象
	urlParams := url.Values{}

	// 遍历 map，将键值对加入 urlParams
	for key, value := range params {
		// 将 value 转换为字符串，并加入 urlParams
		urlParams.Add(key, fmt.Sprintf("%v", value))
	}

	// 拼接 URL 和编码后的参数
	if finalBaseURL == "" {
		return urlParams.Encode()
	}

	return fmt.Sprintf("%s?%s", finalBaseURL, urlParams.Encode())
}
