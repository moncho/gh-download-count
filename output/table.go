package output

import (
	"io"
	"strconv"

	"github.com/moncho/gh-download-count/api"
	"github.com/olekukonko/tablewriter"
)

//ASCIITableWriter writer project details on an ASCII table
type ASCIITableWriter struct{}

//Write writes the details of the given project on an ASCII table format
//to the given writer
func (a ASCIITableWriter) Write(out io.Writer, p *api.Project) error {
	table := tablewriter.NewWriter(out)
	table.SetHeader([]string{"RELEASE", "Downloads"})

	for _, r := range p.Releases() {
		table.Append([]string{r.Name, strconv.Itoa(r.DownloadCount())})
	}
	table.Render()
	return nil
}
