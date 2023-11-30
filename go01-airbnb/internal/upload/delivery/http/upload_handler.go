package uploadhttp

import (
	"fmt"
	"go01-airbnb/pkg/common"
	"go01-airbnb/pkg/upload"

	"github.com/gin-gonic/gin"
)

type uploadHandler struct {
	s3Provider upload.UploadProvider
}

func NewUploadHandler(s3Provider upload.UploadProvider) *uploadHandler {
	return &uploadHandler{s3Provider}
}

func (hdl *uploadHandler) Upload() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		// Save file trực tiếp vào server của mình
		// if err := c.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%s", fileHeader.Filename)); err != nil {
		// 	panic(err)
		// }
		// c.JSON(http.StatusOK, common.Response(common.Image{
		// 	Url: "http://localhost:4000/static/" + fileHeader.Filename,
		// }))

		folder := c.DefaultPostForm("folder", "img")
		fileName := fileHeader.Filename // happy.png

		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrBadRequest(err))
		}
		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrBadRequest(err))
		}

		// Lấy width, height của image
		// w, h, err := getImageDemension(dataBytes)
		// if err != nil {
		// 	// File không phải là hình ảnh
		// 	panic(common.ErrBadRequest(err))
		// }

		img, err := hdl.s3Provider.UploadFile(
			c.Request.Context(),
			dataBytes,
			fmt.Sprintf("%s/%s", folder, fileName),
		)

		if err != nil {
			// Không thẻ upload file
			panic(common.ErrBadRequest(err))
		}

		// img.Width = w
		// img.Height = h

		c.JSON(200, common.Response(img))
	}
}

// func getImageDemension(dataBytes []byte) (int, int, error) {
// 	fileBytes := bytes.NewBuffer(dataBytes)
// 	img, _, err := image.DecodeConfig(fileBytes)
// 	if err != nil {
// 		return 0, 0, err
// 	}

// 	return img.Width, img.Height, nil
// }
