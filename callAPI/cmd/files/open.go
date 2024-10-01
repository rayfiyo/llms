package files

import (
	"fmt"
	"os"
)

// ファイルを作成し開く
func Open(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("can't open the file: %w", err)
	}
	return file, nil
}
