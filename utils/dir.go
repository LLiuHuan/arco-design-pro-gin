package utils

import "os"

//PathExists 文件目录是否存在
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: PathExists
//@description: 文件目录是否存在
//@param: path string
//@return: bool, error
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
