package controllers

import (
	"encoding/base64"
	"gin-mvc/lib/partner/vietqr"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type VietQRRequest struct {
	AccountNo   string `json:"accountNo"`
	AccountName string `json:"accountName"`
	AcqID       int    `json:"acqId"`
	Amount      int    `json:"amount"`
	AddInfo     string `json:"addInfo"`
	Template    string `json:"template"`
}

func CreateVietQRImage(c *gin.Context) {
	client := vietqr.NewVietQr()
	var req VietQRRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu đầu vào không hợp lệ"})
		return
	}

	params := map[string]interface{}{
		"accountNo":   req.AccountNo,
		"accountName": req.AccountName,
		"acqId":       req.AcqID,
		"amount":      req.Amount,
		"addInfo":     req.AddInfo,
		"template":    req.Template,
	}

	resp, err := client.Call(params, "generate", "POST")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Lỗi gọi API VietQR",
		})
		return
	}

	// Lấy chuỗi base64 từ qrDataURL
	dataMap, ok := resp["data"].(map[string]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Không lấy được dữ liệu ảnh",
		})
		return
	}

	qrDataURL := dataMap["qrDataURL"].(string)

	// Tách phần base64 sau dấu phẩy
	base64Data := strings.Split(qrDataURL, ",")
	if len(base64Data) != 2 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Định dạng base64 không hợp lệ",
		})
		return
	}

	imageData, err := base64.StdEncoding.DecodeString(base64Data[1])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Không thể decode ảnh",
		})
		return
	}

	// Trả ảnh về dưới dạng response image/png
	c.Header("Content-Type", "image/png")
	_, err = c.Writer.Write(imageData)
	if err != nil {
		return
	}
}
