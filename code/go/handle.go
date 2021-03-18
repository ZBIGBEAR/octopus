package _go

import (
	"awesomeProject/octopus/common/util"
	"errors"
	"fmt"
)

func handle(filepath string) (error) {
	isFile, err := util.IsFile(filepath)
	if err != nil {
		return err
	}
	
	if !isFile {
		errors.New(fmt.Sprintf("filePath:%s is not a file", filepath))
	}

	return nil
}

