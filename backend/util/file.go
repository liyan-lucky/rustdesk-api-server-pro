package util

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func Unzip(file, dst string) error {
	zipFile, err := zip.OpenReader(file)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	cleanDst, err := filepath.Abs(filepath.Clean(dst))
	if err != nil {
		return err
	}

	if !FileExists(cleanDst) {
		if err := os.MkdirAll(cleanDst, 0755); err != nil {
			return err
		}
	}

	for _, f := range zipFile.File {
		dstPath, err := safeZipDestination(cleanDst, f.Name)
		if err != nil {
			return err
		}
		if f.FileInfo().Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("zip entry %q is a symlink", f.Name)
		}

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(dstPath, safeDirMode(f.Mode())); err != nil {
				return err
			}
			fmt.Println("unzipped", dstPath)
			continue
		}
		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(dstPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, safeFileMode(f.Mode()))
		if err != nil {
			return err
		}
		fileInArchive, err := f.Open()
		if err != nil {
			dstFile.Close()
			return err
		}

		_, copyErr := io.Copy(dstFile, fileInArchive)
		closeArchiveErr := fileInArchive.Close()
		closeDstErr := dstFile.Close()
		if copyErr != nil {
			return copyErr
		}
		if closeArchiveErr != nil {
			return closeArchiveErr
		}
		if closeDstErr != nil {
			return closeDstErr
		}

		fmt.Println("unzipped", dstPath)
	}
	return nil
}

func safeZipDestination(dst, name string) (string, error) {
	if strings.TrimSpace(name) == "" {
		return "", errors.New("empty zip entry name")
	}
	if filepath.IsAbs(name) {
		return "", fmt.Errorf("zip entry %q uses absolute path", name)
	}

	cleanDst := filepath.Clean(dst)
	target := filepath.Clean(filepath.Join(cleanDst, name))
	rel, err := filepath.Rel(cleanDst, target)
	if err != nil {
		return "", err
	}
	if rel == "." || rel == "" {
		return target, nil
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
		return "", fmt.Errorf("zip entry %q escapes destination", name)
	}
	return target, nil
}

func safeDirMode(mode os.FileMode) os.FileMode {
	perm := mode.Perm() & 0755
	if perm == 0 {
		return 0755
	}
	return perm
}

func safeFileMode(mode os.FileMode) os.FileMode {
	perm := mode.Perm() & 0755
	if perm == 0 {
		return 0644
	}
	return perm
}

func MoveFiles(src, dst string) error {
	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, file := range files {
		srcPath := path.Join(src, file.Name())
		dstPath := path.Join(dst, file.Name())
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
		if err := os.Rename(srcPath, dstPath); err != nil {
			return err
		}
	}
	return nil
}
