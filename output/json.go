package output

import (
	"encoding/json"
	"io"

	"github.com/moncho/gh-download-count/api"
)

//JSONWriter writes project details in JSON format
type JSONWriter struct {
}

func (w JSONWriter) Write(out io.Writer, p *api.Project) error {

	e := json.NewEncoder(out)
	return e.Encode(&struct {
		api.Project
		Count int `json:"download_count"`
	}{
		Project: *p,
		Count:   p.DownloadCount(),
	})
}
