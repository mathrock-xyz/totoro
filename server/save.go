package server

import (
	"bytes"
	"context"

	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"github.com/trianglehasfoursides/totoro/storage"
)

func save(ctx echo.Context) (err error) {
	_, err = storage.Bucket.PutObject(
		context.Background(),
		"",
		"",
		bytes.NewReader(data),
		int64(len(data)),
		minio.PutObjectOptions{},
	)

	return
}
