package controller

import (
	"crypto/subtle"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Isiah998/new-api/common"
	"github.com/Isiah998/new-api/constant"
	"github.com/Isiah998/new-api/model"
	"github.com/gin-gonic/gin"
)

var setupMu sync.Mutex

type Setup struct {
	Status             bool   `json:"status"`
	RootInit           bool   `json:"root_init"`
	DatabaseType       string `json:"database_type"`
	SetupTokenRequired bool   `json:"setup_token_required"`
}

type SetupRequest struct {
	Username           string `json:"username"`
	Password           string `json:"password"`
	ConfirmPassword    string `json:"confirmPassword"`
	SelfUseModeEnabled bool   `json:"SelfUseModeEnabled"`
	DemoSiteEnabled    bool   `json:"DemoSiteEnabled"`
	SetupToken         string `json:"setupToken"`
}

func GetSetup(c *gin.Context) {
	setup := Setup{
		Status: constant.Setup,
	}
	if constant.Setup {
		c.JSON(200, gin.H{
			"success": true,
			"data":    setup,
		})
		return
	}
	setup.RootInit = model.RootUserExists()
	setup.DatabaseType = string(common.MainDatabaseType())
	setup.SetupTokenRequired = strings.TrimSpace(os.Getenv("SETUP_TOKEN")) != ""
	c.JSON(200, gin.H{
		"success": true,
		"data":    setup,
	})
}

func PostSetup(c *gin.Context) {
	setupMu.Lock()
	defer setupMu.Unlock()

	// Check if setup is already completed
	if constant.Setup {
		c.JSON(200, gin.H{
			"success": false,
			"message": "系统已经初始化完成",
		})
		return
	}

	// Check if root user already exists
	rootExists := model.RootUserExists()
	var root *model.User

	var req SetupRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "请求参数有误",
		})
		return
	}

	configuredToken := strings.TrimSpace(os.Getenv("SETUP_TOKEN"))
	providedToken := strings.TrimSpace(req.SetupToken)
	if headerToken := strings.TrimSpace(c.GetHeader("X-Setup-Token")); headerToken != "" {
		providedToken = headerToken
	}
	if !setupRequestAuthorized(c.ClientIP(), configuredToken, providedToken) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "初始化请求未通过安全校验；远程初始化必须配置并提供 SETUP_TOKEN",
		})
		return
	}

	// If root doesn't exist, validate and create admin account
	if !rootExists {
		// Keep setup validation aligned with the model username limit.
		if strings.TrimSpace(req.Username) == "" || len(req.Username) > model.UserNameMaxLength {
			c.JSON(200, gin.H{
				"success": false,
				"message": "用户名不能为空且长度不能超过20个字符",
			})
			return
		}
		// Validate password
		if req.Password != req.ConfirmPassword {
			c.JSON(200, gin.H{
				"success": false,
				"message": "两次输入的密码不一致",
			})
			return
		}

		if len(req.Password) < 8 {
			c.JSON(200, gin.H{
				"success": false,
				"message": "密码长度至少为8个字符",
			})
			return
		}

		// Create root user
		hashedPassword, err := common.Password2Hash(req.Password)
		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"message": "系统错误: " + err.Error(),
			})
			return
		}
		rootUser := &model.User{
			Username:    req.Username,
			Password:    hashedPassword,
			Role:        common.RoleRootUser,
			Status:      common.UserStatusEnabled,
			DisplayName: "Root User",
			AccessToken: nil,
			Quota:       100000000,
		}
		root = rootUser
	}

	setup := model.Setup{
		Version:       common.Version,
		InitializedAt: time.Now().Unix(),
	}
	err = model.InitializeSystem(setup, root, map[string]string{
		"SelfUseModeEnabled": boolToString(req.SelfUseModeEnabled),
		"DemoSiteEnabled":    boolToString(req.DemoSiteEnabled),
	})
	if err != nil {
		common.SysLog("system initialization failed: " + err.Error())
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"message": "系统初始化失败，可能已由其他请求完成，请刷新后重试",
		})
		return
	}

	constant.Setup = true

	c.JSON(200, gin.H{
		"success": true,
		"message": "系统初始化成功",
	})
}

func setupRequestAuthorized(clientIP, configuredToken, providedToken string) bool {
	if configuredToken != "" {
		return subtle.ConstantTimeCompare([]byte(configuredToken), []byte(providedToken)) == 1
	}
	parsedIP := net.ParseIP(clientIP)
	return parsedIP != nil && (parsedIP.IsLoopback() || parsedIP.IsPrivate())
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
