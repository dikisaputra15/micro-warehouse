package storage

import (
	"context"
	"fmt"
	"micro-warehouse/merchant-service/configs"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	storage_go "github.com/supabase-community/storage-go"
)

type SupabaseInterface interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader, folder string) (*UploadResult, error)
}

type SupabaseStorage struct {
	client *storage_go.Client
	cfg    configs.Config
}

func NewSupabaseStorage(cfg configs.Config) SupabaseInterface {
	client := storage_go.NewClient(
		cfg.Supabase.URL+"/storage/v1", // ✅ FIX
		cfg.Supabase.Key,
		nil,
	)

	return &SupabaseStorage{
		client: client,
		cfg:    cfg,
	}
}

func (s *SupabaseStorage) UploadFile(ctx context.Context, file *multipart.FileHeader, folder string) (*UploadResult, error) {

	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	// generate filename
	ext := strings.ToLower(filepath.Ext(file.Filename))
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	filePath := fmt.Sprintf("%s/%s", folder, filename)

	// detect content type
	contentType := file.Header.Get("Content-Type")
	if contentType == "" {
		switch ext {
		case ".jpg", ".jpeg":
			contentType = "image/jpeg"
		case ".png":
			contentType = "image/png"
		case ".webp":
			contentType = "image/webp"
		case ".svg":
			contentType = "image/svg+xml"
		default:
			contentType = "application/octet-stream"
		}
	}

	// upload (PAKAI CLIENT YANG SUDAH ADA)
	_, err = s.client.UploadFile(
		s.cfg.Supabase.Bucket,
		filePath,
		src,
		storage_go.FileOptions{
			ContentType: &contentType,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to upload file to supabase: %w", err)
	}

	// generate public URL
	publicURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s",
		s.cfg.Supabase.URL,
		s.cfg.Supabase.Bucket,
		filePath,
	)

	return &UploadResult{
		URL:      publicURL,
		Path:     filePath,
		Filename: filename,
	}, nil
}

type UploadResult struct {
	URL      string `json:"url"`
	Path     string `json:"path"`
	Filename string `json:"filename"`
}