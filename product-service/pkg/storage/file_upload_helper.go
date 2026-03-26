package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2/log"
)

const (
	MaxImageSize = 2 * 1024 * 1024
	AllowedImageExtensions = ".jpg,.jpeg,.png,.webp"
)

type FileUploadHelper struct {
	storage SupabaseInterface
}

func NewFileUploadHelper(storage SupabaseInterface) *FileUploadHelper {
	return &FileUploadHelper{
		storage: storage,
	}
}

func (h *FileUploadHelper) UploadPhoto(ctx context.Context, file *multipart.FileHeader, folder string) (*UploadResult, error) {

	if err := h.validateImageFile(file); err != nil {
		log.Errorf("[FileUploadHelper] validation failed: %v", err)
		return nil, err
	}

	return h.storage.UploadFile(ctx, file, folder)
}

func (h *FileUploadHelper) validateImageFile(file *multipart.FileHeader) error {

	if file.Size > MaxImageSize {
		return fmt.Errorf("file size max 2MB")
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !validateFileExtension(ext) {
		return fmt.Errorf("invalid extension: %s", ext)
	}

	return nil
}

func validateFileExtension(extension string) bool {
	allowed := strings.Split(AllowedImageExtensions, ",")
	for _, ext := range allowed {
		if strings.TrimSpace(ext) == extension {
			return true
		}
	}
	return false
}