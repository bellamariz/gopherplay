package mimetype

import (
	"mime"
)

// Indicate the nature and format of the HLS output
func Configure() {
	mime.AddExtensionType(".m3u8", "application/vnd.apple.mpegURL")
	mime.AddExtensionType(".ts", "video/MP2T")
}
