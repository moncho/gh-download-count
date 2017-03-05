package api

//Project representation of Github project
type Project struct {
	Name     ProjectName `json:"project"`
	URL      string
	releases []Release
}

//ProjectName the identifer of a project
type ProjectName struct {
	Owner string
	Repo  string
}

//NewProject creates a new project
func NewProject(owner, repo, url string, releases []Release) *Project {
	return &Project{ProjectName{owner, repo}, url, releases}

}

//DownloadCount returns the number of time this project has been
//downloaded.
func (p *Project) DownloadCount() int {
	count := 0
	for _, r := range p.releases {
		count += r.DownloadCount()
	}
	return count
}

//Releases returns the release list of this project
func (p *Project) Releases() []Release {
	return p.releases
}

func (p ProjectName) String() string {
	return p.Owner + "/" + p.Repo
}
