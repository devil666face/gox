package start

import (
	"errors"
	"fmt"
	"gox/pkg/cmd"
	"gox/pkg/fs"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	userFlag = "--user"
	replace  = "binpath"
)

var (
	unit = `[Unit]
After=network.target

[Service]
Restart=on-failure
RestartSec=5
RestartPreventExitStatus=666
ExecStart=binpath

[Install]
WantedBy=default.target
`
	binPath = filepath.Join(
		".local", "bin", "gox",
	)
	unitRootPath = filepath.Join(
		"/", "etc", "systemd", "system", "gox.service",
	)
	unitUserPath = filepath.Join(
		".config", "systemd", "user", "gox.service",
	)
)

type systemd struct {
	isPerm                         bool
	unitContent, unitDest, binDest string
	commands                       []*cmd.Command
}

func Systemd(user *user.User, home string) (*systemd, error) {
	_start := &systemd{
		isPerm:   (user.Uid == "0" || os.Getegid() == 0),
		unitDest: unitRootPath,
		binDest:  filepath.Join(home, binPath),
		commands: []*cmd.Command{
			cmd.New("systemctl", "daemon-reload"),
			cmd.New("systemctl", "enable", "gox.service", "--now"),
			cmd.New("systemctl", "restart", "gox.service"),
			cmd.New("loginctl", "enable-linger", user.Username),
		},
	}
	if !_start.isPerm {
		_start.unitDest = filepath.Join(home, unitUserPath)
		for _, command := range _start.commands {
			if strings.HasPrefix(command.String(), "loginctl") {
				continue
			}
			command.Add(userFlag)
		}
	}
	_start.unitContent = strings.ReplaceAll(unit, replace, _start.binDest)
	return _start, nil
}

func (s *systemd) Setup() error {
	return errors.Join(
		s.copyBin(),
		s.writeUnit(),
		func() error {
			for _, command := range s.commands {
				if _, err := command.Run(); err != nil {
					return err
				}
			}
			return nil
		}(),
	)
}

func (s *systemd) Remove() error {
	command := cmd.New("systemctl", "disable", "gox.service", "--now")
	if !s.isPerm {
		command.Add(userFlag)
	}
	return errors.Join(
		func() error {
			if _, err := command.Run(); err != nil {
				return err
			}
			return nil
		}(),
		os.RemoveAll(s.binDest),
		os.RemoveAll(s.unitDest),
	)
}

func (s *systemd) copyBin() error {
	path, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get exe file name: %w", err)
	}
	binPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to get abs path: %w", err)
	}
	if err := fs.Copy(binPath, s.binDest); err != nil {
		return fmt.Errorf("failed to copy: %w", err)
	}
	log.Printf("copy %s to %s", binPath, s.binDest)
	return nil
}

func (s *systemd) writeUnit() error {
	if err := fs.Write(s.unitDest, s.unitContent); err != nil {
		return fmt.Errorf("failed to write unit: %w", err)
	}
	log.Printf(`write
%s
to %s`, s.unitContent, s.unitDest)
	return nil
}
