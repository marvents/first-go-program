package db

import "os"

func SaveData(filename string, bytes []byte, perm os.FileMode) error {

	_, err := os.ReadFile(filename)
	if err != nil { return err }

	err2 := os.WriteFile(filename, bytes, perm)

	if err2 != nil { return err }

	return nil
}

func LoadData(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil { return nil, err }

	return data, nil
}