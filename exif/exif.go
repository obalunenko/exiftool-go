// Package exif provides bindings for [exiftool](https://exiftool.org/) to extract file metadata.
package exif

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os/exec"
)

// Metadata represents file metadata info.
type Metadata []info

// GetFileSize returns file size.
func (m Metadata) GetFileSize() string {
	return m[0].FileSize
}

// GetFileType returns file type.
func (m Metadata) GetFileType() string {
	return m[0].FileType
}

// GetFileTypeExtension returns file type extension.
func (m Metadata) GetFileTypeExtension() string {
	return m[0].FileTypeExtension
}

// GetFilePermissions returns file permissions.
func (m Metadata) GetFilePermissions() string {
	return m[0].FilePermissions
}

// GetMIMEType returns file mime type.
func (m Metadata) GetMIMEType() string {
	return m[0].MIMEType
}

// GetImageWidth returns media width.
func (m Metadata) GetImageWidth() int {
	return m[0].ImageWidth
}

// GetImageHeight returns media height.
func (m Metadata) GetImageHeight() int {
	return m[0].ImageHeight
}

// GetImageSize returns media size.
func (m Metadata) GetImageSize() string {
	return m[0].ImageSize
}

// GetMegapixels returns megapixels.
func (m Metadata) GetMegapixels() float64 {
	return m[0].Megapixels
}

type info struct {
	FileSize          string  `json:"FileSize"`
	FileType          string  `json:"FileType"`
	FileTypeExtension string  `json:"FileTypeExtension"`
	FilePermissions   string  `json:"FilePermissions"`
	MIMEType          string  `json:"MIMEType"`
	ImageWidth        int     `json:"ImageWidth"`
	ImageHeight       int     `json:"ImageHeight"`
	ImageSize         string  `json:"ImageSize"`
	Megapixels        float64 `json:"Megapixels"`
}

const (
	exiftool = "exiftool"
)

// ErrCommandNotFound returned when exif command not found in PATH.
var ErrCommandNotFound = errors.New("exif command not found")

// ExtractMetadata extracts file Metadata.
func ExtractMetadata(source io.Reader) (Metadata, error) {
	exiftoolpath, err := exec.LookPath(exiftool)
	if err != nil {
		return nil, ErrCommandNotFound
	}

	cmd := exec.Command(exiftoolpath, "-j", "-") // #nosec

	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Stdin = source

	// exif will exit and print valid output to stdout
	// if it exits with an unrecognized filetype, don't process
	// that situation here
	if err := cmd.Run(); err != nil && stdout.Len() == 0 {
		return nil, fmt.Errorf("execute failed: %s :%w", stderr.String(), err)
	}

	// no exit error but also no output
	if stdout.Len() == 0 {
		return nil, fmt.Errorf("no output")
	}

	var m Metadata

	if err := json.Unmarshal(stdout.Bytes(), &m); err != nil {
		return nil, fmt.Errorf("unmarhal: %w", err)
	}

	return m, nil
}
