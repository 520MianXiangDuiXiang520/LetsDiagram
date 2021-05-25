package utils

import (
	"bytes"
	"encoding/base64"
	"github.com/nfnt/resize"
	"github.com/pkg/errors"
	"image"
	"image/png"
	_ "image/png"
	`log`
)

func CompressCover(b64 string) (string, error) {
	if len(b64) < 22 {
		log.Printf("b64 does not start with data:image/png;base64,:%s", b64)
		return b64, errors.New("b64 does not start with data:image/png;base64,")
	}
	b64 = b64[22:]
	dimg, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "", errors.Wrap(err, "b64 cannot be decoded!")
	}
	bf := bytes.NewBuffer(dimg)
	img, _, err := image.Decode(bf)
	if err != nil {
		return "", errors.Wrap(err, "failed to decode img")
	}
	afterResize := resize.Resize(250, 0, img, resize.Lanczos3)
	resizeImgBuf := bytes.NewBuffer([]byte{})
	err = png.Encode(resizeImgBuf, afterResize)
	if err != nil {
		return "", errors.Wrap(err, "failed to encode jpeg")
	}
	res := base64.StdEncoding.EncodeToString(resizeImgBuf.Bytes())
	res = "data:image/png;base64," + res
	return res, nil
}
