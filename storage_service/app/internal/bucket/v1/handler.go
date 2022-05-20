package v1

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
)

type BucketHandler struct {
	BucketService *BucketService
}

func (bh *BucketHandler) Register(router fiber.Router) {
	router.Post("/create", bh.createBucket)
	router.Delete("/delete/:bucket", bh.deleteBucket)
}

func (bh *BucketHandler) createBucket(ctx *fiber.Ctx) error {
	createBucketDto := &CreateBucketDTO{}
	err := ctx.BodyParser(createBucketDto)

	if err != nil && createBucketDto.BucketName == "" {
		return errors.New("bucket name parameter is required")
	}
	err = bh.BucketService.CreateBucket(ctx.Context(), createBucketDto.BucketName, minio.MakeBucketOptions{})

	if err != nil {
		return err
	}

	return ctx.Status(201).JSON(fiber.Map{
		"code":    201,
		"message": fmt.Sprintf("Success create %s bucket", createBucketDto.BucketName),
	})
}

func (bh *BucketHandler) deleteBucket(ctx *fiber.Ctx) error {
	bucketName := ctx.Params("bucket")

	if bucketName == "" {
		return errors.New("bucket name parameter is required")
	}

	err := bh.BucketService.DeleteBucket(ctx.Context(), bucketName)

	return err
}
