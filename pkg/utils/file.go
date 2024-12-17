package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"os"

	"github.com/99designs/gqlgen/graphql"
)

func CreateFileHeaderFromUpload(upload graphql.Upload) (*multipart.FileHeader, error) {
	// Create a temporary file to store the uploaded data
	tmpfile, err := os.CreateTemp("", "upload-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	// Ensure the file is closed and removed after use
	defer func() {
		tmpfile.Close()
		if err != nil {
			os.Remove(tmpfile.Name())
		}
	}()

	// Write the uploaded content to the temporary file
	size, err := io.Copy(tmpfile, upload.File)
	if err != nil {
		return nil, fmt.Errorf("failed to write upload content to temp file: %w", err)
	}

	// Rewind the file pointer for any future reads
	_, err = tmpfile.Seek(0, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("failed to reset file pointer: %w", err)
	}

	/*
		type Upload struct {
		    File        io.ReadSeeker
		    Filename    string
		    Size        int64
		    ContentType string
		}

		type FileHeader struct {
		    Filename string
		    Header   textproto.MIMEHeader
		    Size     int64
		    content   []byte
		    tmpfile   string
		    tmpoff    int64
		    tmpshared bool
		}
	*/

	// Create a new FileHeader
	// no like ALL the rest of the FileHeader including content, tmpfile, tmpoff tmpshared
	fileHeader := &multipart.FileHeader{
		Filename: upload.Filename,
		Header:   make(textproto.MIMEHeader),
		Size:     size, // Set the size of the file
	}

	// Add a Content-Disposition header
	fileHeader.Header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, upload.Filename))
	fileHeader.Header.Set("Content-Type", upload.ContentType)

	return fileHeader, nil
}
