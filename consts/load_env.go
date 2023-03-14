package consts

import (
	"github.com/joho/godotenv"
	"os"
	"path"
	"runtime"
	"strconv"
)

var (
	CurrentPath     string
	SshHost         string
	SshPort         int
	SshUser         string
	SshPassword     string
	SshFolderOutput string
	PublisherToken  string
	JavaPath        string
)

func LoadEnv() {
	_, filename, _, _ := runtime.Caller(0)
	CurrentPath = path.Join(path.Dir(filename), "..")
	err := godotenv.Load(path.Join(CurrentPath, ".env"))
	if err != nil {
		panic("Error loading .env file")
	}
	PublisherToken = os.Getenv("PUBLISHER_TOKEN")
	SshHost = os.Getenv("SSH_HOST")
	SshPort, _ = strconv.Atoi(os.Getenv("SSH_PORT"))
	SshUser = os.Getenv("SSH_USER")
	SshPassword = os.Getenv("SSH_PASSWORD")
	SshFolderOutput = os.Getenv("SSH_FOLDER_OUTPUT")
	JavaPath = os.Getenv("JAVA_PATH")
}
