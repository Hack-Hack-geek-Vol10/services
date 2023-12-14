package usecase

import (
	"context"
	"log"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
	image "github.com/schema-creator/services/image-service/api/v1"
	"github.com/schema-creator/services/image-service/internal/domain"
	"github.com/schema-creator/services/image-service/internal/infra"
	"google.golang.org/grpc/status"
)

const (
	ContentTypeJpeg = "image/jpeg"
	ContentTypePng  = "image/png"
	ContentTypeGif  = "image/gif"
	ContentTypeWebp = "image/webp"
	ContentTypeBmp  = "image/bmp"
	ContentTypeTiff = "image/tiff"
	ContentTypeSvg  = "image/svg+xml"
)

type imageService struct {
	image.UnimplementedImageServer
	imageRepo infra.ImageRepo
}

func NewImageService(imageRepo infra.ImageRepo) image.ImageServer {
	return &imageService{
		imageRepo: imageRepo,
	}
}

func (i *imageService) UploadImage(ctx context.Context, arg *image.UploadImageRequest) (*image.UploadImageResponse, error) {
	defer newrelic.FromContext(ctx).StartSegment("grpc-UploadImage").End()
	switch arg.ContentType {
	case ContentTypePng:
		arg.Key = arg.Key + ".png"
	case ContentTypeJpeg:
		arg.Key = arg.Key + ".jpeg"
	case ContentTypeGif:
		arg.Key = arg.Key + ".gif"
	case ContentTypeWebp:
		arg.Key = arg.Key + ".webp"
	case ContentTypeBmp:
		arg.Key = arg.Key + ".bmp"
	case ContentTypeTiff:
		arg.Key = arg.Key + ".tiff"
	case ContentTypeSvg:
		arg.Key = arg.Key + ".svg"
	default:
		return nil, status.Error(http.StatusBadRequest, "invalid content type")
	}

	path, key, err := i.imageRepo.UploadImage(ctx, &domain.UploadImageParam{
		Key:         arg.Key,
		Body:        arg.Data,
		ContentType: arg.ContentType,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &image.UploadImageResponse{
		Path: path,
		Key:  key,
	}, nil
}

func (i *imageService) DeleteImage(ctx context.Context, arg *image.DeleteImageRequest) (*image.DeleteImageResponse, error) {
	defer newrelic.FromContext(ctx).StartSegment("grpc-DeleteImage").End()
	if arg.Key == "" {
		return nil, status.Error(http.StatusBadRequest, "invalid key")
	}

	err := i.imageRepo.DeleteImage(ctx, arg.Key)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &image.DeleteImageResponse{
		Success: true,
	}, nil
}

func (i *imageService) mustEmbedUnimplementedImageServiceServer() {}
