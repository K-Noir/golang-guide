package sdk_release

import (
	"fmt"
	"github.com/mao888/go-utils/version"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
)

func VersionNumberToOrdinal() {
	// 1、从 sdk_release_record 表中查出原数据
	sdkReleaseRecord := make([]*SdkReleaseRecord, 0)
	err := db2.MySQLClientAdmin.Table("sdk_release_record").Find(&sdkReleaseRecord).Error
	if err != nil {
		fmt.Println("从mysql查询 sdk_release_record 错误：", err)
	}

	// 2、根据Id更新
	for _, record := range sdkReleaseRecord {
		record.VersionOrdinal = version.VersionOrdinal(record.VersionNumber)
		fmt.Println(record.VersionOrdinal)
		err := db2.MySQLClientAdmin.Table("sdk_release_record").Where("id = ?", record.ID).
			UpdateColumn("version_ordinal", record.VersionOrdinal).Error
		if err != nil {
			fmt.Println("更新 sdk_release_record version_ordinal 错误：", err)
		}
	}
}
