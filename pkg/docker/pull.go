package docker

import (
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func Pull(tag string) error {
	log.Info("Downloading image...")
	cmd := exec.Command("docker", "pull", tag)

	stderrDone, err := pipeToWithDockerChecks(cmd.StderrPipe, log.Debug)
	if err != nil {
		return err
	}

	err = cmd.Run()
	<-stderrDone
	if err != nil {
		return fmt.Errorf("Failed to download model: %w", err)
	}

	return err
}