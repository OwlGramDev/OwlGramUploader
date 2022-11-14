package file_manager

import (
	"archive/zip"
	"fmt"
	"github.com/cheggaaa/pb"
	"io"
	"os"
)

func ZipFiles(files []string, output string) error {
	archive, err := os.Create(output)
	if err != nil {
		return err
	}
	defer func(archive *os.File) {
		_ = archive.Close()
	}(archive)
	zipWriter := zip.NewWriter(archive)
	for _, file := range files {
		zipFile, err := os.Open(file)
		if err != nil {
			return err
		}
		//goland:noinspection GoDeferInLoop
		defer func(zipFile *os.File) {
			_ = zipFile.Close()
		}(zipFile)
		info, err := zipFile.Stat()
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("Zipping %s...", info.Name()))
		bar := pb.StartNew(int(info.Size()))
		bar.SetUnits(pb.U_BYTES_DEC)
		bar.ShowSpeed = true
		bar.ShowTimeLeft = true
		bar.Format("[##.]")
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Method = zip.Deflate
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		if _, err := io.Copy(bar.NewProxyWriter(writer), zipFile); err != nil {
			return err
		}
		bar.Finish()
		fmt.Println()
	}
	err = zipWriter.Close()
	if err != nil {
		return err
	}
	return nil
}
