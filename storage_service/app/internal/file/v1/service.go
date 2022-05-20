package v1

import (
	"context"
	"errors"
	"github.com/minio/minio-go/v7"
	"io"
	"socia/storage/pkg/minio/client"
)

type FileService struct {
	minioClient *client.MinioClient
}

func NewFileService(minioClient *client.MinioClient) *FileService {
	return &FileService{
		minioClient: minioClient,
	}
}

func (fs *FileService) UploadFile(ctx context.Context, bucketName, fileName string, reader io.Reader, size int64, options minio.PutObjectOptions) (*minio.UploadInfo, error) {
	if bucketName == "" {
		return nil, errors.New("bucket name is required")
	}

	if fileName == "" {
		return nil, errors.New("file name is required")
	}

	info, err := fs.minioClient.UploadFile(ctx, bucketName, fileName, reader, size, options)

	if err != nil {
		return nil, errors.New("cannot upload file")
	}

	return &info, nil
}

func (fs *FileService) GetFile(ctx context.Context, bucketName, fileName string, options minio.GetObjectOptions) (*File, error) {
	if bucketName == "" {
		return nil, errors.New("bucket name is required")
	}

	if fileName == "" {
		return nil, errors.New("file name is required")
	}

	obj, err := fs.minioClient.GetFile(ctx, bucketName, fileName, options)

	if obj == nil {
		return nil, errors.New("cannot get file")
	}

	stat, err := obj.Stat()

	if err != nil {
		return nil, errors.New("cannot get file")
	}

	buffer := make([]byte, stat.Size)

	_, err = obj.Read(buffer)

	if err != nil && err != io.EOF {
		return nil, errors.New("cannot get file")
	}

	file := File{Name: stat.Key, ContentType: stat.Metadata.Get("Content-Type"), Bytes: buffer}

	defer obj.Close()

	return &file, nil
}

func (fs *FileService) RemoveFile(ctx context.Context, bucketName, fileName string, options minio.RemoveObjectOptions) error {
	if bucketName == "" {
		return errors.New("bucket name is required")
	}

	if fileName == "" {
		return errors.New("file name is required")
	}

	err := fs.minioClient.DeleteFile(ctx, bucketName, fileName, options)

	if err != nil {
		return errors.New("cannot delete file")
	}

	return nil
}
