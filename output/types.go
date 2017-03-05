package output

import (
	"io"

	"github.com/moncho/gh-download-count/api"
)

//ProjectWriter interface to write project details into a writer
type ProjectWriter interface {
	Write(out io.Writer, p *api.Project) error
}
