package aliyun

import (
	"os"
	"path"
)

// inject neccessary index.handler adaptor for aliyun function
func (m* Manager) createPython3Function(dir string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// TODO: config
	err = os.Link(path.Join(home, ".jfManager", "ali", "python3", "jointfaas.py") ,path.Join(dir, "jointfaas.py"))
	if err != nil {
		return err
	}
	return nil
}