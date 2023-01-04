//go:build !integration_test

package exif_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/obalunenko/exiftool-go/exif"
)

func TestMetadata_GetFilePermissions(t *testing.T) {
	tests := []struct {
		name string
		m    exif.Metadata
		want string
	}{
		{
			name: "",
			m: exif.Metadata{
				{
					FileSize:          "1533 KiB",
					FileType:          "MP4",
					FileTypeExtension: "mp4",
					FilePermissions:   "rw-r--r--",
					MIMEType:          "video/mp4",
					ImageWidth:        480,
					ImageHeight:       270,
					ImageSize:         "480x270",
					Megapixels:        0.13,
				},
			},
			want: "rw-r--r--",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.GetFilePermissions()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMetadata_GetFileSize(t *testing.T) {
	tests := []struct {
		name string
		m    exif.Metadata
		want string
	}{
		{
			name: "",
			m: exif.Metadata{
				{
					FileSize:          "1533 KiB",
					FileType:          "MP4",
					FileTypeExtension: "mp4",
					FilePermissions:   "rw-r--r--",
					MIMEType:          "video/mp4",
					ImageWidth:        480,
					ImageHeight:       270,
					ImageSize:         "480x270",
					Megapixels:        0.13,
				},
			},
			want: "1533 KiB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.GetFileSize()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMetadata_GetFileType(t *testing.T) {
	tests := []struct {
		name string
		m    exif.Metadata
		want string
	}{
		{
			name: "",
			m: exif.Metadata{
				{
					FileSize:          "1533 KiB",
					FileType:          "MP4",
					FileTypeExtension: "mp4",
					FilePermissions:   "rw-r--r--",
					MIMEType:          "video/mp4",
					ImageWidth:        480,
					ImageHeight:       270,
					ImageSize:         "480x270",
					Megapixels:        0.13,
				},
			},
			want: "MP4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.GetFileType()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMetadata_GetFileTypeExtension(t *testing.T) {
	tests := []struct {
		name string
		m    exif.Metadata
		want string
	}{
		{
			name: "",
			m: exif.Metadata{
				{
					FileSize:          "1533 KiB",
					FileType:          "MP4",
					FileTypeExtension: "mp4",
					FilePermissions:   "rw-r--r--",
					MIMEType:          "video/mp4",
					ImageWidth:        480,
					ImageHeight:       270,
					ImageSize:         "480x270",
					Megapixels:        0.13,
				},
			},
			want: "mp4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.GetFileTypeExtension()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMetadata_GetImageHeight(t *testing.T) {
	tests := []struct {
		name string
		m    exif.Metadata
		want int
	}{
		{
			name: "",
			m: exif.Metadata{
				{
					FileSize:          "1533 KiB",
					FileType:          "MP4",
					FileTypeExtension: "mp4",
					FilePermissions:   "rw-r--r--",
					MIMEType:          "video/mp4",
					ImageWidth:        480,
					ImageHeight:       270,
					ImageSize:         "480x270",
					Megapixels:        0.13,
				},
			},
			want: 270,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.GetImageHeight()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMetadata_GetImageSize(t *testing.T) {
	tests := []struct {
		name string
		m    exif.Metadata
		want string
	}{
		{
			name: "",
			m: exif.Metadata{
				{
					FileSize:          "1533 KiB",
					FileType:          "MP4",
					FileTypeExtension: "mp4",
					FilePermissions:   "rw-r--r--",
					MIMEType:          "video/mp4",
					ImageWidth:        480,
					ImageHeight:       270,
					ImageSize:         "480x270",
					Megapixels:        0.13,
				},
			},
			want: "480x270",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.GetImageSize()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMetadata_GetImageWidth(t *testing.T) {
	tests := []struct {
		name string
		m    exif.Metadata
		want int
	}{
		{
			name: "",
			m: exif.Metadata{
				{
					FileSize:          "1533 KiB",
					FileType:          "MP4",
					FileTypeExtension: "mp4",
					FilePermissions:   "rw-r--r--",
					MIMEType:          "video/mp4",
					ImageWidth:        480,
					ImageHeight:       270,
					ImageSize:         "480x270",
					Megapixels:        0.13,
				},
			},
			want: 480,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.GetImageWidth()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMetadata_GetMIMEType(t *testing.T) {
	tests := []struct {
		name string
		m    exif.Metadata
		want string
	}{
		{
			name: "",
			m: exif.Metadata{
				{
					FileSize:          "1533 KiB",
					FileType:          "MP4",
					FileTypeExtension: "mp4",
					FilePermissions:   "rw-r--r--",
					MIMEType:          "video/mp4",
					ImageWidth:        480,
					ImageHeight:       270,
					ImageSize:         "480x270",
					Megapixels:        0.13,
				},
			},
			want: "video/mp4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.GetMIMEType()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMetadata_GetMegapixels(t *testing.T) {
	tests := []struct {
		name string
		m    exif.Metadata
		want float64
	}{
		{
			name: "",
			m: exif.Metadata{
				{
					FileSize:          "1533 KiB",
					FileType:          "MP4",
					FileTypeExtension: "mp4",
					FilePermissions:   "rw-r--r--",
					MIMEType:          "video/mp4",
					ImageWidth:        480,
					ImageHeight:       270,
					ImageSize:         "480x270",
					Megapixels:        0.13,
				},
			},
			want: 0.13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.GetMegapixels()

			assert.Equal(t, tt.want, got)
		})
	}
}
