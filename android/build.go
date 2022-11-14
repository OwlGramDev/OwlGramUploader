package android

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path"
	"publish/consts"
	"strings"
)

func BuildGradlew(buildVariant string) error {
	var errRes bytes.Buffer
	cmd := exec.Command(path.Join(consts.ProjectPath, "gradlew"), strings.Split(buildVariant, " ")...)
	cmd.Stderr = &errRes
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "JAVA_HOME="+consts.JavaPath)
	cmd.Dir = consts.ProjectPath
	err := cmd.Run()
	if err != nil {
		if len(errRes.String()) > 0 {
			return errors.New(errRes.String())
		}
		return err
	}
	return nil
}
