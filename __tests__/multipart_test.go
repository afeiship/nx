/**
 * @Author: aric 1290657123@qq.com
 * @Date: 2024-10-15 21:06:33
 * @LastEditors: aric 1290657123@qq.com
 * @LastEditTime: 2024-10-15 21:58:34
 */
package nx

import (
	"fmt"
	"os"
	"testing"

	nx "github.com/afeiship/nx/lib"
)

// $ go test __tests__/multipart_test.go -v
// https://chatgpt.com/c/672eb949-2c80-8013-9375-cf6495cb88b6
func TestCreateMultipartRequestBody(t *testing.T) {
	filePath := "hello.txt"
	// get io.Reader from file
	file, err := os.Open(filePath)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	// MultipartOptions
	opts := nx.MultipartOptions{
		FieldName:     "file",
		FileFieldName: "file.txt",
		FileReader:    file,
		ExtraFields: map[string]string{
			"name": "hello.txt",
		},
	}

	// create multipart request body
	requestBody, err := nx.CreateMultipartRequestBody(opts)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(requestBody.Body)
	fmt.Println(requestBody.ContentType)

}
