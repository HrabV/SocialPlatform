package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	bucketV1 "socia/storage/internal/bucket/v1"
	"socia/storage/internal/config"
	fileV1 "socia/storage/internal/file/v1"
	"socia/storage/internal/health/handler"
	"socia/storage/pkg/minio/client"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("[CONFIG] %s\n", err.Error())
	}

	mc, err := client.NewMinioClient(cfg.MinioAccessKey, cfg.MinioSecretKey, cfg.MinioEndpoint, cfg.MinioUseSSL)

	if err != nil {
		log.Fatalf("[MINIO] %s\n", err.Error())
	}

	app := fiber.New(fiber.Config{
		ServerHeader:                 cfg.ServerHeader,
		EnablePrintRoutes:            true,
		DisablePreParseMultipartForm: true,
		BodyLimit:                    cfg.ServerBodyLimit,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(400).JSON(fiber.Map{
				"code":    400,
				"message": err.Error(),
			})
		},
	})

	bsV1 := bucketV1.NewBucketService(mc)
	bhV1 := bucketV1.BucketHandler{BucketService: bsV1}
	bhV1.Register(app.Group("/v1/buckets/"))

	fsV1 := fileV1.NewFileService(mc)
	fhV1 := fileV1.FileHandler{FileService: fsV1}
	fhV1.Register(app.Group("/v1/files/"))

	healthHandler := handler.HealthHandler{}
	healthHandler.Register(app.Group("/healthcheck"))

	log.Fatal(app.Listen(cfg.ServerAddr))
}
