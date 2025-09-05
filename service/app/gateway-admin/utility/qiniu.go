package utility

import (
	"bytes"
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/qiniu/go-sdk/v7/auth/qbox"

	"github.com/qiniu/go-sdk/v7/storage"
)

func UploadToQiniu(ctx context.Context, fileContent []byte, fileName string) (string, string, error) {
	cfg := g.Cfg().MustGet(ctx, "qiniu")
	if cfg.IsEmpty() {
		return "", "", errors.New("七牛云配置缺失")
	}

	// 解析配置
	qiniuCfg := cfg.Map()
	accessKey := qiniuCfg["accessKey"].(string)
	secretKey := qiniuCfg["secretKey"].(string)
	bucket := qiniuCfg["bucket"].(string)
	domain := qiniuCfg["domain"].(string)

	// 生成上传凭证
	putPolicy := storage.PutPolicy{Scope: bucket}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	// 配置上传参数（华南区）
	cfgUpload := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseHTTPS:      true,
		UseCdnDomains: false,
	}

	// 获取文件扩展名
	fileExt := gfile.ExtName(fileName)
	if fileExt == "" {
		fileExt = "jpg" // 默认扩展名
	}

	// 生成文件名
	key := grand.S(16) + "." + fileExt // 16位随机字符串

	// 创建表单上传器
	formUploader := storage.NewFormUploader(&cfgUpload)
	ret := storage.PutRet{}

	// 上传文件 - 使用标准库的 bytes.Reader
	err := formUploader.Put(
		context.Background(),
		&ret,
		upToken,
		key,
		bytes.NewReader(fileContent),
		int64(len(fileContent)),
		nil,
	)
	if err != nil {
		return "", "", err
	}

	// 返回完整访问URL
	return domain + "/" + key, key, nil
}
