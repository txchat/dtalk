package minio

import (
	"context"
	"fmt"
	"io"
	"sort"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
	"github.com/txchat/dtalk/pkg/oss"
)

var _ oss.Oss = (*Minio)(nil)

type Minio struct {
	cfg    *oss.Config
	client *minio.Core
}

func New(cfg *oss.Config) *Minio {
	m := &Minio{
		cfg:    cfg,
		client: nil,
	}
	err := m.createClient()
	if err != nil {
		panic(err)
	}

	return m

}

func (m *Minio) Config() *oss.Config {
	return m.cfg
}

func (m *Minio) AssumeRole() (*oss.AssumeRoleResp, error) {
	return nil, errors.New("minio not support assume role")
}

func (m *Minio) Upload(key string, body io.Reader, size int64) (url, uri string, err error) {
	ctx := context.Background()
	_, err = m.client.PutObject(ctx, m.cfg.Bucket, key, body, size, "", "", minio.PutObjectOptions{})
	if err != nil {
		return "", "", errors.WithMessage(err, "minio.client.PutObject")
	}
	return m.parseUrl(key), m.parseUri(key), nil
}

func (m *Minio) InitiateMultipartUpload(key string) (uploadId string, err error) {
	ctx := context.Background()
	uploadId, err = m.client.NewMultipartUpload(ctx, m.cfg.Bucket, key, minio.PutObjectOptions{})
	if err != nil {
		return "", errors.WithMessage(err, "minio.client.NewMultipartUpload")
	}
	return uploadId, nil
}

func (m *Minio) UploadPart(key, uploadId string, body io.Reader, partNumber int32, offset,
	partSize int64) (ETag string, err error) {
	ctx := context.Background()
	output, err := m.client.PutObjectPart(ctx, m.cfg.Bucket, key, uploadId, int(partNumber), body, partSize, "", "", nil)
	if err != nil {
		return "", errors.WithMessage(err, "minio.client.PutObjectPart")
	}
	return output.ETag, nil
}

func (m *Minio) CompleteMultipartUpload(key, uploadId string, parts []oss.Part) (url, uri string, err error) {
	ctx := context.Background()
	minioParts := make([]minio.CompletePart, len(parts), len(parts))
	sort.Sort(oss.Parts(parts))
	for i := range parts {
		minioParts[i] = minio.CompletePart{
			ETag:       parts[i].ETag,
			PartNumber: int(parts[i].PartNumber),
		}
	}

	_, err = m.client.CompleteMultipartUpload(ctx, m.cfg.Bucket, key, uploadId, minioParts, minio.PutObjectOptions{})
	if err != nil {
		return "", "", errors.WithMessage(err, "minio.client.CompleteMultipartUpload")
	}
	return m.parseUrl(key), m.parseUri(key), nil
}

func (m *Minio) AbortMultipartUpload(key, uploadId string) error {
	ctx := context.Background()
	err := m.client.AbortMultipartUpload(ctx, m.cfg.Bucket, key, uploadId)
	if err != nil {
		return errors.WithMessage(err, "minio.client.AbortMultipartUpload")
	}
	return nil
}

func (m *Minio) createClient() (err error) {
	m.client, err = minio.NewCore(
		m.cfg.EndPoint,
		&minio.Options{
			Creds: credentials.NewStaticV4(m.cfg.AccessKeyId, m.cfg.AccessKeySecret, ""),
			// todo
			Secure: false,
		})
	if err != nil {
		return errors.WithMessage(err, "minio minio.NewCore")
	}
	return nil
}

// parseUrl 获取对象在 minio 上的完整访问URL
func (m *Minio) parseUrl(key string) string {
	return fmt.Sprintf("%s/%s/%s", m.cfg.PublicUrl, m.cfg.Bucket, key)
}

// parseUri ...
func (m *Minio) parseUri(key string) string {
	return fmt.Sprintf("/%s/%s", m.cfg.Bucket, key)
}

// GetHost ...
func (m *Minio) GetHost() string {
	return fmt.Sprintf("http://%s", m.cfg.EndPoint)
}
