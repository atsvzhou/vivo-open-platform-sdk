package vivo

import (
	"fmt"
	"testing"
)

func TestGetAppDetail(t *testing.T) {
	client := NewVivoClient("XXX", "XXX")
	params := &GetAppDetailParams{
		PackageName:  "XXXX",
		Format:       "json",
		Method:       "app.query.details",
		SignMethod:   "hmac",
		TargetAppKey: "developer",
		V:            "1.0",
	}
	res, err := client.GetAppDetail(params)
	if err != nil {
		return
	}
	fmt.Println("111", res.Data.VersionName)
}

func TestPublishAppVersion(t *testing.T) {
	client := NewVivoClient("xxx", "xxx")
	params := &PublishVersionParams{
		ApkMd5:       "xxx",
		ApkUrl:       "xxx",
		Format:       "json",
		Method:       "app.update.app",
		OnlineType:   "1",
		PackageName:  "xxx",
		Remark:       "xxx",
		SignMethod:   "hmac",
		TargetAppKey: "developer",
		UpdateDesc:   "调整模板来源数据和管理方式。",
		V:            "1.0",
		VersionCode:  "xxx",
	}
	res, err := client.PublishVersion(params)
	if err != nil {
		return
	}
	fmt.Println("222", res)
}
