package vivo

type GetAppDetailParams struct {
	PackageName  string `json:"packageName"`
	Format       string `json:"format"`
	Method       string `json:"method"`
	SignMethod   string `json:"sign_method"`
	TargetAppKey string `json:"target_app_key"`
	V            string `json:"v"`
}

type GetAppDetailRes struct {
	Code int      `json:"code"`
	Data VivoData `json:"data"`
}

type VivoData struct {
	VersionName  string `json:"versionName"`
	SaleStatus   int    `json:"saleStatus"`
	Status       int    `json:"status"`
	UnPassReason string `json:"unPassReason"`
}
type PublishVersionParams struct {
	ApkMd5       string `json:"apkMd5"`
	ApkUrl       string `json:"apkUrl"`
	Format       string `json:"format"`
	Method       string `json:"method"`
	OnlineType   string `json:"onlineType"`
	PackageName  string `json:"packageName"`
	Remark       string `json:"remark"`
	SignMethod   string `json:"sign_method"`
	TargetAppKey string `json:"target_app_key"`
	UpdateDesc   string `json:"updateDesc"`
	V            string `json:"v"`
	VersionCode  string `json:"versionCode"`
}

type PublishVersionRes struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	SubCode   string `json:"subCode"`
	Timestamp int64  `json:"timestamp"`
}
