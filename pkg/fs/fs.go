package fs

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func Copy(src, dst string) error {
	return errors.Join(
		checkExists(dst),
		CopyFile(src, dst),
	)
}

func Write(path, data string) error {
	return errors.Join(
		checkExists(path),
		WriteFile(path, []byte(data)),
	)
}

func WriteFile(
	path string,
	data []byte,
	perm ...int,
) error {
	if len(perm) == 0 {
		perm = append(perm, 0644)
	}
	if err := os.WriteFile(path, data, fs.FileMode(perm[0])); err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}
	return nil
}

func CopyFile(
	src, dst string,
	perm ...int,
) error {
	if len(perm) == 0 {
		perm = append(perm, 0755)
	}

	_src, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open src file: %w", err)
	}
	defer _src.Close()

	_dst, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create dst file: %w", err)
	}
	defer _dst.Close()

	if _, err := io.Copy(_dst, _src); err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	if err := os.Chmod(dst, fs.FileMode(perm[0])); err != nil {
		return fmt.Errorf("failed to set permissions: %w", err)
	}

	return nil
}

func checkExists(dst string) error {
	dir := filepath.Dir(dst)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	if _, err := os.Stat(dst); err == nil {
		if err := os.Remove(dst); err != nil {
			return fmt.Errorf("failed to remove existing dst file: %w", err)
		}
	}
	return nil
}
