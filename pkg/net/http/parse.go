package http

import (
	"mime/multipart"
	"net/http"
)

const (
	defaultMultipartMemory = 32 << 20 // 32 MB
)

// FromFile 请求表单文件获取
func FromFile(r *http.Request, name string) (*multipart.FileHeader, error) {
	if r.MultipartForm == nil {
		if err := r.ParseMultipartForm(defaultMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := r.FormFile(name)
	if err != nil {
		return nil, err
	}
	f.Close()
	return fh, err
}
