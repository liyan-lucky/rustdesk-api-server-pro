package github

import (
	"encoding/json"
	"fmt"
	"rustdesk-api-server-pro/internal/errcode"
	"rustdesk-api-server-pro/util"
)

type Release struct {
	ID              int     `json:"id"`
	NodeID          string  `json:"node_id"`
	TagName         string  `json:"tag_name"`
	TargetCommitish string  `json:"target_commitish"`
	Name            string  `json:"name"`
	Draft           bool    `json:"draft"`
	Prerelease      bool    `json:"prerelease"`
	CreatedAt       string  `json:"created_at"`
	PublishedAt     string  `json:"published_at"`
	Assets          []Asset `json:"assets"`
}

type Asset struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Label              string `json:"label"`
	ContentType        string `json:"content_type"`
	State              string `json:"state"`
	Size               int    `json:"size"`
	DownloadCount      int    `json:"download_count"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

func GetReleases(repo string) (*[]Release, error) {
	endpoint := fmt.Sprintf("https://api.github.com/repos/%s/releases", repo)
	resp, err := util.HttpGetString(endpoint)
	if err != nil {
		return nil, errcode.Errorf(errcode.ERRE001.Code, errcode.ERRE001.Message)
	}
	releases := &[]Release{}
	if err = json.Unmarshal([]byte(resp), releases); err != nil {
		return nil, errcode.Errorf(errcode.ERRE002.Code, errcode.ERRE002.Message)
	}
	return releases, nil
}

func GetLatestRelease(repo string) (*Release, error) {
	endpoint := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)
	resp, err := util.HttpGetString(endpoint)
	if err != nil {
		return nil, errcode.Errorf(errcode.ERRE003.Code, errcode.ERRE003.Message)
	}
	release := &Release{}
	if err = json.Unmarshal([]byte(resp), release); err != nil {
		return nil, errcode.Errorf(errcode.ERRE004.Code, errcode.ERRE004.Message)
	}
	return release, nil
}

func GetReleaseByTag(repo, tag string) (*Release, error) {
	endpoint := fmt.Sprintf("https://api.github.com/repos/%s/releases/tags/%s", repo, tag)
	resp, err := util.HttpGetString(endpoint)
	if err != nil {
		return nil, errcode.Errorf(errcode.ERRE005.Code, errcode.ERRE005.Message)
	}
	release := &Release{}
	if err = json.Unmarshal([]byte(resp), release); err != nil {
		return nil, errcode.Errorf(errcode.ERRE006.Code, errcode.ERRE006.Message)
	}
	return release, nil
}
