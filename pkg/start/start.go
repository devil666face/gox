package start

import (
	"errors"
	"fmt"
	"os"
	"os/user"
)

type Start struct {
	systemd *systemd
}

func New() (*Start, error) {
	_user, err := user.Current()
	if err != nil {
		return nil, err
	}
	_home, err := os.UserHomeDir()
	if err != nil {
		_home = _user.HomeDir
	}
	_systemd, err := Systemd(_user, _home)
	if err != nil {
		return nil, fmt.Errorf("systemd fatal error: %w", err)
	}
	return &Start{
		systemd: _systemd,
	}, nil
}

func (s *Start) Setup() error {
	return errors.Join(
		s.systemd.Setup(),
	)
}

func (s *Start) Remove() error {
	return errors.Join(
		s.systemd.Remove(),
	)
}
