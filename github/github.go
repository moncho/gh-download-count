package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/moncho/gh-download-count/api"
)

//NewProject creates a project from a github repo
func NewProject(owner, repo string) (*api.Project, error) {
	ghAPI, projectURL := buildProjectURLs(owner, repo)

	resp, err := http.Get(ghAPI)

	if !is2xx(resp) {
		return nil, fmt.Errorf("Github responded with a %d code.\n", resp.StatusCode)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var releases []api.Release

	if err := json.Unmarshal(data, &releases); err != nil {
		return nil, err
	}
	return api.NewProject(owner, repo, projectURL, releases), nil
}

func is2xx(resp *http.Response) bool {
	return resp.StatusCode > 199 && resp.StatusCode < 300
}

func buildProjectURLs(owner, repo string) (string, string) {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", owner, repo),
		fmt.Sprintf("https://github.com/%s/%s", owner, repo)
}
