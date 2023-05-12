package zccompress

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path"
	"path/filepath"
)

func CompressDirToTargz(dirPath, targetTarGzPath string) error {
	tarFile, err := os.Create(targetTarGzPath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(tarFile)

	gzWriter := gzip.NewWriter(tarFile)
	defer func(gzWriter *gzip.Writer) {
		err := gzWriter.Close()
		if err != nil {
			panic(err)
		}
	}(gzWriter)

	tarWriter := tar.NewWriter(gzWriter)
	defer func(tarWriter *tar.Writer) {
		err := tarWriter.Close()
		if err != nil {
			panic(err)
		}
	}(tarWriter)

	err = filepath.Walk(dirPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		baseDir := filepath.Base(dirPath)
		relPath, err := filepath.Rel(dirPath, filePath)
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}
		header.Name = path.Join(baseDir, relPath)

		if err = tarWriter.WriteHeader(header); err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				panic(err)
			}
		}(file)

		_, err = io.Copy(tarWriter, file)
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

func UnCompressTargzToDir(sourceTarGzPath, dirPath string) error {
	sourceTarGz, err := os.Open(sourceTarGzPath)
	if err != nil {
		return err
	}
	defer func(sourceTarGz *os.File) {
		err := sourceTarGz.Close()
		if err != nil {
			panic(err)
		}
	}(sourceTarGz)

	gzReader, err := gzip.NewReader(sourceTarGz)
	if err != nil {
		return err
	}
	defer func(gzReader *gzip.Reader) {
		err := gzReader.Close()
		if err != nil {
			panic(err)
		}
	}(gzReader)

	tarReader := tar.NewReader(gzReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		targetPath := filepath.Join(dirPath, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, os.FileMode(header.Mode)); err != nil {
				return err
			}
		case tar.TypeReg:
			file, err := os.OpenFile(targetPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			//goland:noinspection GoDeferInLoop
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					panic(err)
				}
			}(file)

			if _, err := io.Copy(file, tarReader); err != nil {
				return err
			}
		}
	}
	return nil
}
