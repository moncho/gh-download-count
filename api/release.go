package api

import "fmt"

//Asset is an artifact associated with a release
type Asset struct {
	Count int `json:"download_count"`
}

//Release a release of a project in Github
type Release struct {
	Name   string
	Assets []Asset
}

//DownloadCount returns the number of time this release has been
//downloaded.
func (r *Release) DownloadCount() int {
	count := 0
	for _, a := range r.Assets {
		count += a.Count
	}
	return count
}
func (r *Release) String() string {
	return fmt.Sprintf("Release %s has been downloaded %d times. ", r.Name, r.DownloadCount())
}
