package S3

import (
	"context"
	"io"

	"github.com/closure-studio/imgHost/server/utils/consts"
	"github.com/closure-studio/imgHost/server/utils/resp"
	"github.com/closure-studio/imgHost/utils/storage"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

// UploadHandler 处理文件上传
func UploadHandler() fiber.Handler {
	return func(c fiber.Ctx) error {
		// 解析上传文件
		file, err := c.FormFile("file")
		if err != nil {
			return resp.Failed(c, 0, consts.FILE_NOT_PROVIDED)
		}

		// 读取文件内容
		// 打开文件流
		fileBytes, err := file.Open()
		if err != nil {
			return resp.Failed(c, 0, consts.UNABLE_TO_READ_FILE)
		}
		defer fileBytes.Close() // 确保文件句柄关闭

		// 读取完整的文件数据
		fileData, err := io.ReadAll(fileBytes)
		if err != nil {
			return resp.Failed(c, 0, consts.UNABLE_TO_READ_FILE)
		}

		fileName := c.FormValue("filename")
		if fileName == "" {
			fileName = uuid.NewString()
		}

		// 获取 MIME 类型
		contentType := file.Header.Get("Content-Type")
		// 上传到 S3
		err = storage.S3Instance.UploadFile(context.Background(), fileName, fileData, contentType)
		if err != nil {
			return resp.Failed(c, 0, consts.UNABLE_TO_UPLOAD_FILE)
		}
		return resp.Success(c, fileName, consts.UPLOAD_SUCCESS)
	}
}
