package output

import (
	"fmt"
	"io"

	"github.com/moncho/gh-download-count/api"
)

//QuietWriter just writes the download count.
type QuietWriter struct {
}

//Write writes the total download count on the given writer.
func (w QuietWriter) Write(out io.Writer, p *api.Project) error {
	_, err := fmt.Fprintf(out, "%d\n", p.DownloadCount())
	return err
}
