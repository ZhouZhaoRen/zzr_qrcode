package qrcode

import (
	"fmt"
	"image/jpeg"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	uuid "github.com/satori/go.uuid"
)

func main() {
	qrcode := New("https://www.baidu.com/",400,400)
	err := qrcode.GenerateQrcode()
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("chenggong")
}

type Qrcode struct {
	Url    string
	Height int
	Width  int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

// New 定义创建二维码的路径，高度和宽度
func New(url string,height,width int) *Qrcode {
	return &Qrcode{
		Url:    url,
		Height: height,
		Width:  width,
		Ext:    ".jpg",
		Level:  qr.M,
		Mode:   qr.Auto,
	}
}
// GenerateQrcode 生成二维码，默认存放在当前目录下的qrcode文件夹下
func (q *Qrcode) GenerateQrcode() error {
	src := "qrcode/" + uuid.NewV4().String() + ".jpg"
	_, err := os.Stat(src)
	if os.IsNotExist(err) {
		mkdir()
	}
	//
	code, err := qr.Encode(q.Url, q.Level, q.Mode)
	if err != nil {
		return err
	}

	code, err = barcode.Scale(code, q.Width, q.Height)
	if err != nil {
		return err
	}
	file, _ := os.OpenFile(src, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)

	defer file.Close()

	err = jpeg.Encode(file, code, nil)
	if err != nil {
		return err
	}

	return nil
}

func (q *Qrcode) generate(src string) error {
	code, err := qr.Encode(q.Url, q.Level, q.Mode)
	if err != nil {
		return err
	}

	code, err = barcode.Scale(code, q.Width, q.Height)
	if err != nil {
		return err
	}
	file, _ := os.OpenFile(src, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)

	defer file.Close()

	err = jpeg.Encode(file, code, nil)
	if err != nil {
		return err
	}
	return nil
}


func mkdir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/qrcode/", os.ModePerm)
	if err != nil {
		fmt.Println("创建文件夹失败：", err)
	}
}
