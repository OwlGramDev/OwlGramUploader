package server

import (
	"context"
	"fmt"
	"github.com/bramvdbogaerde/go-scp"
	"github.com/cheggaaa/pb"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"net"
	"os"
	"path"
	"publish/consts"
)

func UploadFile(filePath string) error {
	ctx := context.Background()
	var auths []ssh.AuthMethod
	if aConn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		auths = append(auths, ssh.PublicKeysCallback(agent.NewClient(aConn).Signers))
	}
	auths = append(auths, ssh.Password(consts.SshPassword))
	config := ssh.ClientConfig{
		User:            consts.SshUser,
		Auth:            auths,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", consts.SshHost, consts.SshPort), &config)
	if err != nil {
		return err
	}
	defer func(client *ssh.Client) {
		_ = client.Close()
	}(conn)
	client, err := scp.NewClientBySSH(conn)
	if err != nil {
		return err
	}
	err = client.Connect()
	if err != nil {
		return err
	}
	srcFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func(srcFile *os.File) {
		_ = srcFile.Close()
	}(srcFile)
	fi, e := srcFile.Stat()
	if e != nil {
		return e
	}
	fileSize := fi.Size()
	bar := pb.StartNew(int(fileSize))
	bar.Format("[##.]")
	bar.SetUnits(pb.U_BYTES_DEC)
	bar.ShowSpeed = true
	bar.ShowTimeLeft = true
	err = client.Copy(ctx, bar.NewProxyReader(srcFile), path.Join(consts.SshFolderOutput, "bundles.zip"), "0655", fileSize)
	if err != nil {
		return err
	}
	bar.Finish()
	return nil
}
