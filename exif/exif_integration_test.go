// +build integration_test

package exif_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/obalunenko/exiftool-go/exif"
)

func TestExtractMetadata(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    exif.Metadata
		wantErr bool
	}{
		{
			name: "jpeg landscape",
			args: args{
				filepath: filepath.Join("testdata", "467-300x200.jpg"),
			},
			want: exif.Metadata{
				{
					FileSize:          "4.1 KiB",
					FileType:          "JPEG",
					FileTypeExtension: "jpg",
					FilePermissions:   "-rw-r--r--",
					MIMEType:          "image/jpeg",
					ImageWidth:        300,
					ImageHeight:       200,
					ImageSize:         "300x200",
					Megapixels:        0.06,
				},
			},
			wantErr: false,
		},
		{
			name: "jpeg WEBP screenshot",
			args: args{
				filepath: filepath.Join("testdata", "914-200x300.jpg"),
			},
			want: exif.Metadata{
				{
					FileSize:          "6.4 KiB",
					FileType:          "JPEG",
					FileTypeExtension: "jpg",
					FilePermissions:   "-rw-r--r--",
					MIMEType:          "image/jpeg",
					ImageWidth:        200,
					ImageHeight:       300,
					ImageSize:         "200x300",
					Megapixels:        0.06,
				},
			},
			wantErr: false,
		},
		{
			name: "avi video",
			args: args{
				filepath: filepath.Join("testdata", "file_example_AVI_480_750kB.avi"),
			},
			want: exif.Metadata{
				{
					FileSize:          "725 KiB",
					FileType:          "AVI",
					FileTypeExtension: "avi",
					FilePermissions:   "-rw-r--r--",
					MIMEType:          "video/x-msvideo",
					ImageWidth:        480,
					ImageHeight:       270,
					ImageSize:         "480x270",
					Megapixels:        0.13,
				},
			},
			wantErr: false,
		},
		{
			name: "avi video",
			args: args{
				filepath: filepath.Join("testdata", "file_example_OGG_480_1_7mg.ogg"),
			},
			want: exif.Metadata{
				{
					FileSize:          "1694 KiB",
					FileType:          "OGV",
					FileTypeExtension: "ogv",
					FilePermissions:   "-rw-r--r--",
					MIMEType:          "video/ogg",
					ImageWidth:        480,
					ImageHeight:       270,
					ImageSize:         "480x270",
					Megapixels:        0.13,
				},
			},
			wantErr: false,
		},
		{
			name: "avi video",
			args: args{
				filepath: filepath.Join("testdata", "file_example_WEBP_50kB.webp"),
			},
			want: exif.Metadata{
				{
					FileSize:          "49 KiB",
					FileType:          "WEBP",
					FileTypeExtension: "webp",
					FilePermissions:   "-rw-r--r--",
					MIMEType:          "image/webp",
					ImageWidth:        1050,
					ImageHeight:       700,
					ImageSize:         "1050x700",
					Megapixels:        0.735,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Open(tt.args.filepath)
			require.NoError(t, err)

			t.Cleanup(func() {
				require.NoError(t, file.Close())
			})

			got, err := exif.ExtractMetadata(file)
			if tt.wantErr {
				assert.Error(t, err)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
