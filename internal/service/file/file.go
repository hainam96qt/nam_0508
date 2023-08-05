package file

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"os"

	configs "nam_0508/config"
	"nam_0508/internal/model"
	db "nam_0508/internal/repo/dbmodel"
	convert_type "nam_0508/pkg/util/convert-type"
)

// Size constants
const (
	// 8MB
	maxSizeUpload = 8 * 1 << 20
)

type (
	Service struct {
		conf *configs.Config

		DatabaseConn *sql.DB
		Query        *db.Queries
	}
)

func NewUploadFileService(conf *configs.Config, DatabaseConn *sql.DB) *Service {
	query := db.New(DatabaseConn)
	return &Service{
		conf:         conf,
		DatabaseConn: DatabaseConn,
		Query:        query,
	}
}

func (s *Service) UploadFile(ctx context.Context, identity *model.Identity, file multipart.File, fileHeader *multipart.FileHeader) error {
	if fileHeader.Size > maxSizeUpload {
		return errors.New("images larger than 8 mb")
	}

	// save to /temp
	tempFile, err := os.CreateTemp("./temp", "upload-*.png")
	if err != nil {
		return err
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	tempFile.Write(fileBytes)

	var metadata = db.ImageMetadata{
		FileName: fileHeader.Filename,
		Size:     fileHeader.Size,
		Note:     "Note something for upload file",
	}

	b, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	// create database record
	newImage := db.CreateImageParams{
		UserID:   convert_type.NewNullInt32(int32(identity.USerID)),
		Path:     convert_type.NewNullString(tempFile.Name()),
		Metadata: b,
	}
	err = s.Query.CreateImage(ctx, newImage)
	if err != nil {
		return err
	}
	return nil
}
