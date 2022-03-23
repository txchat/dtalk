package minio

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"github.com/txchat/dtalk/pkg/oss"
	"math"
	"os"
	"sort"
	"testing"
)

/*
[[Oss]]
AppId = "dtalk"
OssType = "minio"
RegionId = ""
AccessKeyId = "XYI4T3QIT8YQQLRHA0YV"
AccessKeySecret = "FqgvB3CzsBK5xwEphEC6i4Y6dTkWAjyfQ9TS1kLZ"
Role = ""
Policy = ""
DurationSeconds = 3600
Bucket = "dtalk-test"
EndPoint = "127.0.0.1:9000"
*/

var (
	minioConfig = &oss.Config{
		RegionId:        "",
		AccessKeyId:     "XYI4T3QIT8YQQLRHA0YV",
		AccessKeySecret: "FqgvB3CzsBK5xwEphEC6i4Y6dTkWAjyfQ9TS1kLZ",
		Role:            "",
		Policy:          "",
		DurationSeconds: 3600,
		Bucket:          "dtalk-test",
		EndPoint:        "127.0.0.1:9000",
	}
	smFilePath = "/Users/cccccccccchy/Downloads/oa.sql"
	bgFilePath = "/Users/cccccccccchy/Downloads/movie_bg.mp4"
	OssMinio   = &Minio{}
)

func TestMain(m *testing.M) {
	OssMinio = New(minioConfig)
	os.Exit(m.Run())
}

func TestMinio_Upload_sm(t *testing.T) {
	file, err := os.OpenFile(smFilePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fileStat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	url, _, err := OssMinio.Upload("test_smFile.gif", file, fileStat.Size())
	require.NoError(t, err)
	t.Log(url)
}

func TestMinio_Upload_bg(t *testing.T) {
	file, err := os.OpenFile(bgFilePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fileStat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	url, _, err := OssMinio.Upload("test_movie_bg.mp4", file, fileStat.Size())
	require.NoError(t, err)
	t.Log(url)
}

func TestMinio_Multipart_bg(t *testing.T) {
	file, err := os.OpenFile(smFilePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fileStat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var chunkSize int64 = 1024 * 1024 * 5
	var num int64
	var key string = "oa.sql"
	var uploadId string
	var parts []oss.Part

	uploadId, err = OssMinio.InitiateMultipartUpload(key)
	require.NoError(t, err)

	num = int64(math.Ceil(float64(fileStat.Size()) / float64(chunkSize)))

	// 执行并发上传段
	partChan := make(chan oss.Part, num)
	var i int64 = 0
	for ; i < num; i++ {
		go func(i int64) {
			b := make([]byte, chunkSize)
			_, _ = file.Seek(i*(chunkSize), 0)
			if len(b) > int(fileStat.Size()-i*chunkSize) {
				b = make([]byte, fileStat.Size()-i*chunkSize)
			}

			file.Read(b)
			r := bytes.NewReader(b)
			etag, err := OssMinio.UploadPart(key, uploadId, r, int32(i+1), 0, int64(len(b)))
			require.NoError(t, err)
			partChan <- oss.Part{
				ETag:       etag,
				PartNumber: int32(i + 1),
			}

			t.Log("partNumber", i+1, "len", len(b))

		}(i)
	}

	parts = make([]oss.Part, 0, num)
	// 等待上传完成
	for {
		part, ok := <-partChan
		if !ok {
			break
		}
		parts = append(parts, part)

		if len(parts) == int(num) {
			close(partChan)
		}
	}

	sort.Sort(oss.Parts(parts))

	url, _, err := OssMinio.CompleteMultipartUpload(key, uploadId, parts)
	require.NoError(t, err)
	t.Log(url)
}
