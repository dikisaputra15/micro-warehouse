package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"micro-warehouse/product-service/configs"
	"mime/multipart"
	"path/filepath"
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

// UploadFile implements SupabaseInterface.
func (s *SupabaseStorage) UploadFile(ctx context.Context, file *multipart.FileHeader, folder string) (*UploadResult, error) {
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	// generate filename
	ext := filepath.Ext(file.Filename)
	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("%d%s", timestamp, ext)

	filePath := fmt.Sprintf("%s/%s", folder, filename)

	// detect content type
	contentType := file.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	// 🔥 FIX PENTING: convert ke bytes
	fileBytes, err := io.ReadAll(src)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// 🔥 FIX PENTING: pakai FileOptions
	upsert := true
	_, err = s.client.UploadFile(
		s.cfg.Supabase.Bucket,
		filePath,
		bytes.NewReader(fileBytes),
		storage_go.FileOptions{
			ContentType: &contentType,
			Upsert:      &upsert,
		},
	)
	if err != nil {
		fmt.Printf("UPLOAD ERROR DETAIL: %+v\n", err)
		return nil, fmt.Errorf("failed to upload file to supabase: %w", err)
	}

	// 🔥 FIX URL
	publicUrl := s.client.GetPublicUrl(s.cfg.Supabase.Bucket, filePath)

	return &UploadResult{
		URL:      publicUrl.SignedURL,
		Path:     filePath,
		Filename: filename,
	}, nil
}

type UploadResult struct {
	URL      string `json:"url"`
	Path     string `json:"path"`
	Filename string `json:"filename"`
}

func NewSupabaseStorage(cfg configs.Config) SupabaseInterface {
	client := storage_go.NewClient(cfg.Supabase.URL, cfg.Supabase.Key, nil)
	return &SupabaseStorage{
		client: client,
		cfg:    cfg,
	}
}
