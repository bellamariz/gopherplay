package mimetype

import (
	"mime"
)

// Indicate the nature and format of the HLS output
func Configure() {
	mime.AddExtensionType(".m3u8", "application/vnd.apple.mpegURL") //nolint:errcheck //no reason to check mime errors
	mime.AddExtensionType(".ts", "video/MP2T")                      //nolint:errcheck //no reason to check mime errors
}
