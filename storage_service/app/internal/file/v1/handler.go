package v1

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"socia/storage/pkg/utils"
)

type FileHandler struct {
	FileService *FileService
}

func (fh *FileHandler) Register(router fiber.Router) {
	router.Post("/upload/:bucket", fh.uploadFile)
	router.Get("/get/:bucket/:file", fh.getFile)
	router.Delete("/delete/:bucket/:file", fh.deleteFile)
}

func (fh *FileHandler) uploadFile(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	bucketName := ctx.Params("bucket")

	if err != nil {
		return errors.New("cannot upload file")
	}

	if file != nil {
		buffer, err := file.Open()

		if buffer != nil {
			defer buffer.Close()
		}

		fileName := utils.GenerateFileObjectName(file.Filename)

		info, err := fh.FileService.UploadFile(ctx.Context(), bucketName, fileName, buffer, file.Size, minio.PutObjectOptions{
			ContentType: file.Header["Content-Type"][0],
		})

		if err != nil {
			return err
		}

		return ctx.Status(201).JSON(&info)
	}

	return nil
}

func (fh *FileHandler) getFile(ctx *fiber.Ctx) error {
	bucketName := ctx.Params("bucket")
	fileName := ctx.Params("file")
	obj, err := fh.FileService.GetFile(ctx.Context(), bucketName, fileName, minio.GetObjectOptions{})

	if err != nil {
		return err
	}

	ctx.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", obj.Name))
	ctx.Set("Content-Type", obj.ContentType)
	return ctx.Send(obj.Bytes)
}

func (fh *FileHandler) deleteFile(ctx *fiber.Ctx) error {
	bucketName := ctx.Params("bucket")
	fileName := ctx.Params("file")

	err := fh.FileService.RemoveFile(ctx.Context(), bucketName, fileName, minio.RemoveObjectOptions{})

	if err != nil {
		return err
	}

	return nil
}
