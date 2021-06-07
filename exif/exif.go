package exif

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os/exec"
)

type Metadata []info

func (m Metadata) GetFileSize() string {
	return m[0].FileSize
}

func (m Metadata) GetFileType() string {
	return m[0].FileType
}

func (m Metadata) GetFileTypeExtension() string {
	return m[0].FileTypeExtension
}

func (m Metadata) GetFilePermissions() string {
	return m[0].FilePermissions
}

func (m Metadata) GetMIMEType() string {
	return m[0].MIMEType
}

func (m Metadata) GetImageWidth() int {
	return m[0].ImageWidth
}

func (m Metadata) GetImageHeight() int {
	return m[0].ImageHeight
}

func (m Metadata) GetImageSize() string {
	return m[0].ImageSize
}

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

var ErrCommandNotFound = errors.New("exif command not found")

func ExtractMetadata(source io.Reader) (Metadata, error) {
	exiftoolpath, err := exec.LookPath(exiftool)
	if err != nil {
		return nil, ErrCommandNotFound
	}

	cmd := exec.Command(exiftoolpath, "-j", "-")

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
