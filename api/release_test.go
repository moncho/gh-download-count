package api

import "testing"

func TestReleases(t *testing.T) {
	assets := [2]Asset{Asset{Count: 1}, Asset{Count: 2}}
	r := Release{"release", assets[:]}

	if r.DownloadCount() != 3 {
		t.Errorf("Unexpected number of download for release %s. Got %d, expected 3", r.Name, r.DownloadCount())
	}
}
