package download

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/smallsha123/php2go/tool"

	"github.com/kr/pretty"
	"github.com/smallsha123/php2go/uniqueid"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

const ExtImg = ".jpg"

// 下载网络图片
func OriginImg(url string, ext string) (string, error) {
	pic := fmt.Sprintf("%s%d%s/", "/tmp/", time.Now().Year(), time.Now().Month().String())
	isExist, _ := tool.IsFileExist(pic)
	if !isExist {
		os.Mkdir(pic, os.ModePerm)
	}

	newId, _ := uniqueid.GenId()
	sid := cast.ToString(newId)
	if len(ext) == 0 {
		ext = ExtImg
	}
	pic += sid + ext

	v, err := http.Get(url)
	if err != nil {
		logx.Errorf("Http get [%v] failed! %v", url, err)
		return "", err
	}
	defer v.Body.Close()
	content, err := ioutil.ReadAll(v.Body)
	if err != nil {
		logx.Errorf("Read http response failed! %v", err)
		return "", err
	}
	err = ioutil.WriteFile(pic, content, 0666)
	fmt.Sprintf("%# v\n", pretty.Formatter(pic))

	return pic, nil
}
