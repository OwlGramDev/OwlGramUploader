package main

import (
	"fmt"
	"path"
	"publish/android"
	"publish/consts"
	"publish/file_manager"
	"publish/http"
	"publish/server"
)

func main() {
	consts.LoadEnv()
	fmt.Println("Building Bundle...")
	err := android.BuildGradlew(":TMessagesProj_App:bundlePlay")
	if err != nil {
		panic(err)
	}
	fmt.Println("Building APKs...")
	err = android.BuildGradlew(":TMessagesProj_App:assembleRelease")
	if err != nil {
		panic(err)
	}
	files, err := file_manager.GetFiles(consts.BuildOutputsPath)
	if err != nil {
		panic(err)
	}
	zipFile := path.Join(consts.CurrentPath, "temp", "publish.zip")
	err = file_manager.ZipFiles(files, zipFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("Sending to NAOS Server...")
	err = server.UploadFile(zipFile)
	if err != nil {
		panic(err)
	}
	err = http.Notify()
	if err != nil {
		fmt.Println("\nError when sending to OwlGram Servers, ", err)
	} else {
		fmt.Println("\nSent to OwlGram Servers!")
	}
}
