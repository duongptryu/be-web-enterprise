package common

import "os"

func CheckFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		if err != nil {
			return ErrInternal(err)
		}
	}

	return nil
}
