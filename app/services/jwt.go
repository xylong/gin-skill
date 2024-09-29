package services

import (
	"context"
	"fmt"
	"gin-skill/app/dao"
	"gin-skill/app/models"
	"gin-skill/global"
	"gin-skill/utils"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/cast"
	"time"
)

type jwtService struct {
}

var (
	JwtService = new(jwtService)
)

// JwtUser 所有需要颁发 token 的用户模型必须实现这个接口
type JwtUser interface {
	GetUid() string
}

// CustomClaims 自定义 Claims
type CustomClaims struct {
	jwt.StandardClaims
}

const (
	TokenType    = "bearer"
	AppGuardName = "app"
)

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// CreateToken 生成 Token
func (s *jwtService) CreateToken(GuardName string, user JwtUser) (tokenData TokenOutPut, err error, token *jwt.Token) {
	token = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + global.App.Config.Jwt.Expire,
				Id:        user.GetUid(),
				Issuer:    GuardName, // 用于在中间件中区分不同客户端颁发的 token，避免 token 跨端使用
				NotBefore: time.Now().Unix() - 1000,
			},
		},
	)

	tokenStr, err := token.SignedString([]byte(global.App.Config.Jwt.Secret))

	tokenData = TokenOutPut{
		tokenStr,
		int(global.App.Config.Jwt.Expire),
		TokenType,
	}
	return
}

func (s *jwtService) getBlackListKey(token string) string {
	return "jwt_black_list:" + utils.MD5([]byte(token))
}

// JoinBlackList token 加入黑名单
func (s *jwtService) JoinBlackList(token *jwt.Token) (err error) {
	nowUnix := time.Now().Unix()
	timer := time.Duration(token.Claims.(*CustomClaims).ExpiresAt-nowUnix) * time.Second
	// 将 token 剩余时间设置为缓存有效期，并将当前时间作为缓存 value 值
	return global.App.Redis.SetNX(context.Background(), s.getBlackListKey(token.Raw), nowUnix, timer).Err()
}

// IsInBlacklist token 是否在黑名单中
func (s *jwtService) IsInBlacklist(tokenStr string) bool {
	joinUnixStr, err := global.App.Redis.Get(context.Background(), s.getBlackListKey(tokenStr)).Result()
	if joinUnixStr == "" || err != nil {
		return false
	}

	// JwtBlacklistGracePeriod 为黑名单宽限时间，避免并发请求失效
	return time.Now().Unix()-cast.ToInt64(joinUnixStr) > global.App.Config.Jwt.JwtBlacklistGracePeriod
}

// GetUserInfo 获取用户信息
func (s *jwtService) GetUserInfo(guardName string, id string) (error, JwtUser) {
	var (
		err  error
		user *models.User
	)

	switch guardName {
	case AppGuardName:
		u := dao.User
		user, err = u.Where(u.ID.Eq(cast.ToInt64(id))).Preload(dao.User.Profile).First()
		return err, user
	default:
		return fmt.Errorf("guard " + guardName + " does not exist"), nil
	}
}
