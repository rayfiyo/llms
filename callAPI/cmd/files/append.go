package files

import (
	"fmt"
)

// ファイルに文字列を追記する
func Append(fileName, text string) error {
	// ファイルを開く（存在しない場合は新規作成）
	file, err := Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// ファイルに文字列を追記
	_, err = file.WriteString(text + "\n")
	if err != nil {
		return fmt.Errorf("can't append the file: %w", err)
	}

	return nil
}
