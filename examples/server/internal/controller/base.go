package controller

import (
	"errors"
	"examples/server/config"
	"examples/server/internal/constants"
	"examples/server/internal/model"
	"net/http"

	"github.com/rulego/rulego/endpoint/rest"

	"github.com/golang-jwt/jwt"
	"github.com/rulego/rulego/api/types"
	endpointApi "github.com/rulego/rulego/api/types/endpoint"
)

var ErrIllegalToken = errors.New("illegal token")

var Base = &base{}

type base struct {
}

type RuleGoClaim struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// userNotFound 用户不存在
func userNotFound(username string, exchange *endpointApi.Exchange) bool {
	_ = "STUB: not implemented"
	return false
}

// unauthorized 用户未授权
func unauthorized(username string, exchange *endpointApi.Exchange) bool {
	_ = "STUB: not implemented"
	return false
}

// GetRuleGoFunc 动态获取指定用户规则链池
func GetRuleGoFunc(exchange *endpointApi.Exchange) types.RuleEnginePool {
	_ = "STUB: not implemented"
	return *new(types.RuleEnginePool)
}

var AuthProcess = func(router endpointApi.Router, exchange *endpointApi.Exchange) bool {
	var metadata *types.Metadata
	if r, ok := exchange.In.(*rest.RequestMessage); ok {
		metadata = r.Metadata
	} else if r, ok := exchange.In.(endpointApi.HeaderModifier); ok {
		metadata = r.GetMetadata()
	}

	// 先从 header 获取 authorization
	authorization := exchange.In.Headers().Get(constants.KeyAuthorization)

	if !config.Get().RequireAuth && authorization == "" {
		//允许匿名访问
		if metadata != nil {
			metadata.PutValue(constants.KeyUsername, config.C.DefaultUsername)
		}
		return true
	}
	username := getUsernameApiKey(authorization) // "Bearer api_key" 方式
	if username != "" {
		if metadata != nil {
			metadata.PutValue(constants.KeyUsername, username)
		}
		return true
	} else {
		claim, err := parseToken(authorization) // "Bearer jwt" 方式
		if err != nil {
			exchange.Out.SetStatusCode(http.StatusUnauthorized)
			exchange.Out.SetBody([]byte(err.Error()))
			return false
		}
		if metadata != nil {
			metadata.PutValue(constants.KeyUsername, claim.Username)
		}
		return true
	}

}

func GetComponentsFromMarketplace(baseUrl, keywords string, root *bool, currentPage, size int) (ComponentList, error) {
	_ = "STUB: not implemented"
	// 构造查询参数
	return *new(ComponentList), nil
}

// 拼接完整的 URL

// 发送 GET 请求

func parseToken(token string) (*RuleGoClaim, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *base) Login(url string) endpointApi.Router {
	_ = "STUB: not implemented"
	return *new(endpointApi.Router)
}

// 设置 Token 过期时间
// 设置 Token 的签发者

func createToken(claim jwt.Claims) (*string, error) {
	_ = "STUB: not implemented"
	// 创建 JWT Token
	return nil, nil
}

func validatePassword(user model.User) bool { _ = "STUB: not implemented"; return false }

func getUsernameApiKey(token string) string { _ = "STUB: not implemented"; return "" }
