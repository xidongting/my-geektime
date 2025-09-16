package utils

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zkep/my-geektime/internal/global"
)

// GetStorageURL 根据配置生成存储文件的URL
// 如果启用了AutoHost，则使用请求的host；否则使用配置的host
func GetStorageURL(c *gin.Context, key string) string {
	if key == "" {
		return ""
	}

	// 输入验证：确保gin.Context不为nil
	if c == nil {
		return global.Storage.GetUrl(key)
	}

	if global.CONF.Storage.AutoHost {
		host := getRequestHost(c)
		return global.Storage.GetUrlWithHost(key, host)
	}

	return global.Storage.GetUrl(key)
}

// getRequestHost 从请求中获取完整的host（包含协议）
// 支持反向代理场景，自动检测HTTPS协议
func getRequestHost(c *gin.Context) string {
	scheme := getScheme(c)
	host := c.Request.Host
	return fmt.Sprintf("%s://%s", scheme, host)
}

// getScheme 获取请求协议，支持反向代理场景
func getScheme(c *gin.Context) string {
	// 优先检查 X-Forwarded-Proto 头（反向代理设置）
	if proto := c.GetHeader("X-Forwarded-Proto"); proto != "" {
		return strings.ToLower(proto)
	}

	// 检查 X-Forwarded-Ssl 头
	if ssl := c.GetHeader("X-Forwarded-Ssl"); ssl == "on" {
		return "https"
	}

	// 检查 X-Url-Scheme 头
	if scheme := c.GetHeader("X-Url-Scheme"); scheme != "" {
		return strings.ToLower(scheme)
	}

	// 最后检查 TLS 连接
	if c.Request.TLS != nil {
		return "https"
	}

	return "http"
}
