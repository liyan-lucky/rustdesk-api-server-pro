package util

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"rustdesk-api-server-pro/internal/errcode"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
)

const (
	maxHTTPStringSize = 10 * 1024 * 1024
	maxDownloadSize   = 1024 * 1024 * 1024
)

var _httpProxy = ""

func SetHttpProxy(proxy string) {
	_httpProxy = proxy
}

func HttpClient() (*http.Client, error) {
	client := &http.Client{Timeout: 60 * time.Second}
	if _httpProxy != "" {
		proxyUrl, err := url.Parse(_httpProxy)
		if err != nil {
			return nil, err
		}
		scheme := strings.ToLower(proxyUrl.Scheme)
		allowSchemes := []string{"http", "https", "socks5"}
		if !InArray(allowSchemes, scheme) {
			return nil, errcode.New(errcode.ERRC001.Code, errcode.ERRC001.Message)
		}

		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
	}
	return client, nil
}

func HttpGetString(remoteAddr string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, remoteAddr, nil)
	if err != nil {
		return "", err
	}
	client, err := HttpClient()
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return "", errcode.Errorf(errcode.ERRC003.Code, errcode.ERRC003.Message, resp.StatusCode)
	}
	if resp.ContentLength > maxHTTPStringSize {
		return "", errcode.Errorf(errcode.ERRC002.Code, errcode.ERRC002.Message)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, maxHTTPStringSize+1))
	if err != nil {
		return "", err
	}
	if len(body) > maxHTTPStringSize {
		return "", errcode.New(errcode.ERRC002.Code, errcode.ERRC002.Message)
	}
	return string(body), nil
}

func DownloadFile(remoteAddr, filename string, showConsoleProgressBar bool) error {
	req, err := http.NewRequest(http.MethodGet, remoteAddr, nil)
	if err != nil {
		return err
	}
	httpClient, err := HttpClient()
	if err != nil {
		return err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return errcode.Errorf(errcode.ERRC004.Code, errcode.ERRC004.Message, resp.StatusCode)
	}
	if resp.ContentLength > maxDownloadSize {
		return errcode.New(errcode.ERRC005.Code, errcode.ERRC005.Message)
	}

	dir := filepath.Dir(filename)
	tmpFile, err := os.CreateTemp(dir, filepath.Base(filename)+".*.tmp")
	if err != nil {
		return err
	}
	tmpName := tmpFile.Name()
	cleanup := true
	defer func() {
		if cleanup {
			_ = os.Remove(tmpName)
		}
	}()

	reader := io.LimitReader(resp.Body, maxDownloadSize+1)
	var written int64
	if showConsoleProgressBar {
		bar := progressbar.DefaultBytes(resp.ContentLength, "Downloading")
		written, err = io.Copy(io.MultiWriter(tmpFile, bar), reader)
	} else {
		written, err = io.Copy(tmpFile, reader)
	}
	if err != nil {
		_ = tmpFile.Close()
		return err
	}
	if written > maxDownloadSize {
		_ = tmpFile.Close()
		return errcode.New(errcode.ERRC005.Code, errcode.ERRC005.Message)
	}
	if err = tmpFile.Chmod(0644); err != nil {
		_ = tmpFile.Close()
		return err
	}
	if err = tmpFile.Close(); err != nil {
		return err
	}
	if err = os.Rename(tmpName, filename); err != nil {
		return err
	}
	cleanup = false
	return nil
}
