package fs

var (
	// BasePath is the mother path of all stored files
	// on the file system
	BasePath = "/store/default"
)

func FileExist(file string) error {
	return nil
}

func CreateFile(bytes []byte, file string) error {
	return nil
}

func DeleteFile() error {
	return nil
}

func CreateDir(file string) error {
	return nil
}

func DeleteDir() error {
	return nil
}
