package system

import (
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/system"
	"go.uber.org/zap"
)

type JwtService struct {
}

//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.AdpDb.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//IsBlacklist 判断JWT是否在黑名单内部
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool
func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
	//err := global.GVA_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	//isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	//return !isNotFound
}

//GetRedisJWT 从redis取jwt
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: err error, redisJWT string
func (jwtService *JwtService) GetRedisJWT(userName string) (err error, redisJWT string) {
	//redisJWT, err = global.GVA_REDIS.Get(context.TODO(), userName).Result()
	return err, redisJWT
}

//SetRedisJWT jwt存入redis并设置过期时间
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error
func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	//timer := time.Duration(global.GVA_CONFIG.JWT.ExpiresTime) * time.Second
	//err = global.GVA_REDIS.Set(context.TODO(), userName, jwt, timer).Err()
	return err
}

func LoadAll() {
	var data []string
	err := global.AdpDb.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.AdpLog.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
