package output

import (
	"fmt"
	"io"

	"github.com/moncho/gh-download-count/api"
)

//DefaultProjectWriter is the default project writer.
type DefaultProjectWriter struct {
}

//Write writes the project name and the total download count on the given writer.
func (w DefaultProjectWriter) Write(out io.Writer, p *api.Project) error {
	_, err := fmt.Fprintf(out, "Project %s has been downloaded %d times.\n", p.Name, p.DownloadCount())
	return err
}
