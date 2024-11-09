package nx

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
)

// 定义 Options 结构体，包含所有可配置的参数
type MultipartOptions struct {
	FieldName     string            // 文件字段名，定义在请求体中的名称
	FileFieldName string            // 请求中的文件字段名
	FileReader    io.Reader         // 文件内容，来自任何 io.Reader
	ExtraFields   map[string]string // 其他额外的字段
}

// 定义 MultipartRequest 结构体，用于封装返回的请求体和内容类型
type MultipartRequest struct {
	Body        *bytes.Buffer
	ContentType string
}

// 使用 Options 结构体创建 multipart 请求体
func CreateMultipartRequestBody(options MultipartOptions) (*MultipartRequest, error) {
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 创建文件字段
	part, err := writer.CreateFormFile(options.FieldName, options.FileFieldName)
	if err != nil {
		return nil, fmt.Errorf("error creating form file: %w", err)
	}

	// 将文件内容写入请求体
	_, err = io.Copy(part, options.FileReader)
	if err != nil {
		return nil, fmt.Errorf("error copying file content: %w", err)
	}

	// 添加其他字段
	for key, value := range options.ExtraFields {
		err = writer.WriteField(key, value)
		if err != nil {
			return nil, fmt.Errorf("error adding field %s: %w", key, err)
		}
	}

	// 关闭 writer
	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("error closing writer: %w", err)
	}

	// 返回封装的结构体
	return &MultipartRequest{
		Body:        &requestBody,
		ContentType: writer.FormDataContentType(),
	}, nil
}
