/**
** @创建时间: 2020/11/4 8:45 下午
** @作者　　: return
** @描述　　:
 */
package bootstrap

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io/ioutil"
	"strings"
)

type qiNiu struct {
	Enabled   bool   `json:"enabled"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	IsHttps   bool   `json:"isHttps"`
	Domain    string `json:"domain"`
	Bucket    string `json:"bucket"`
	Zone      string `json:"zone"`
	IsCdn     bool   `json:"isCdn"`
}

type QiNiu struct {
	Bucket string
}

type Config struct {
	qiNiu `json:"qiniu"`
	BucketManager *storage.BucketManager `json:"-"`
}

var (
	qiNiuConfig *Config
	publicPath  string
)

func StartInit(path string) {
	publicPath = path
	qiNiuConfig = NewQiuNiu()
}

func NewQiuNiu() *Config {

	if qiNiuConfig == nil {
		data, err := ioutil.ReadFile(publicPath)
		if err != nil {
			panic(err.Error())
		}
		//读取的数据为json格式，需要进行解码
		err = json.Unmarshal(data, &qiNiuConfig)
		if err != nil {
			panic(err.Error())
		}

	}

	return qiNiuConfig
}

func QiuNiuConf() *Config {
	return qiNiuConfig
}


func (qn QiNiu) UploadFile(key string, localPath string) (string, error) {

	bucket := qiNiuConfig.Bucket
	if qn.Bucket != "" {
		bucket = qn.Bucket
	}

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}

	mac := qbox.NewMac(qiNiuConfig.AccessKey, qiNiuConfig.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	zone := strings.ToLower(qiNiuConfig.Zone)
	switch zone {
	case "huadong":
		cfg.Zone = &storage.ZoneHuadong
	case "huabei":
		cfg.Zone = &storage.ZoneHuabei
	case "huanan":
		cfg.Zone = &storage.ZoneHuanan
	case "beimei":
		cfg.Zone = &storage.ZoneBeimei
	default:
		cfg.Zone = &storage.ZoneHuadong
	}

	// 是否使用https域名
	cfg.UseHTTPS = qiNiuConfig.IsHttps
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = qiNiuConfig.IsCdn
	resumeUploader := storage.NewResumeUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.RputExtra{}
	err := resumeUploader.PutFile(context.Background(), &ret, upToken, key, localPath, &putExtra)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return ret.Key, nil
}
