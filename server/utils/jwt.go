package utils

import (
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	jwt.StandardClaims
	UserId   int64 `json:"userId"`
	UserRole int64 `json:"userRole"`
}

// GenerateToken
//
//	@Description: 生成token
//	@param userId 用户id
//	@param security 加密字符串
//	@param role 用户角色
//	@param expireTime 过期时间
//	@param issure 签发者
//	@return string token
//	@return error 错误
func GenerateToken(userId int64, security string, role int64, expireTime time.Time, issure string) (string, error) {
	signedString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + expireTime.Unix(),
			NotBefore: time.Now().Unix(),
			Issuer:    issure,
		},
		UserId:   userId,
		UserRole: role,
	}).SignedString([]byte(security))
	return signedString, err
}

// TokenVal
//
//	@Description: 验证token
//	@param jwtToken token
//	@return claims token信息
//	@return err 错误
func TokenVal(jwtToken string) (claims *MyClaims, err error) {
	tokenClaims, err := jwt.ParseWithClaims(jwtToken, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.JWT_SECRET), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return claims, err
}
