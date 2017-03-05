package api

import "testing"

const (
	owner = "owner"
	repo  = "repo"
	url   = "url"
)

func TestProjectCreation(t *testing.T) {
	var releases []Release

	assets := [1]Asset{Asset{Count: 1}}

	for index := 0; index < 3; index++ {
		releases = append(releases, Release{"release", assets[:]})
	}
	p := NewProject(owner, repo, url, releases)

	if p == nil {
		t.Error("Project was not created")
	}
	if p.Name.Owner != owner || p.Name.Repo != repo {
		t.Errorf("Project name is not correct. expected: %s/%s, got: %s", owner, repo, p.Name)
	}

	if p.DownloadCount() != 3 {
		t.Errorf("Download count is not correct. expected: 3, got: %d", p.DownloadCount())
	}
}
