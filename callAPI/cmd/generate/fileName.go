package generate

import (
	"time"
)

func FileName() string {
	// 現在時刻を取得
	now := time.Now()

	// 指定のフォーマットでファイル名を生成
	return "./logs/" + now.Format("2006-01-02-150405_") + ".md"
}
