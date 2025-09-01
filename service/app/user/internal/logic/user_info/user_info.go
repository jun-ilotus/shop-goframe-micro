package user_info

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"service/app/user/internal/dao"
	"service/app/user/internal/model/entity"
	"service/utility"
	"time"
)

func Login(ctx context.Context, name, password string) (token string, expire int, userInfo *entity.UserInfo, err error) {
	// 1. 参数校验
	if name == "" || password == "" {
		return "", 0, nil, errors.New("name or password is empty")
	}

	// 2.查询用户
	userRecord, err := dao.UserInfo.Ctx(ctx).Where("name", name).One()
	if err != nil {
		g.Log().Errorf(ctx, "查询用户失败: %v", err)
		return "", 0, nil, errors.New("系统错误")
	}
	if userRecord.IsEmpty() {
		return "", 0, nil, errors.New("用户不存在")
	}

	// 3.转换为实体
	var user entity.UserInfo
	if err := userRecord.Struct(&user); err != nil {
		g.Log().Errorf(ctx, "用户数据解析失败: %v", err)
		return "", 0, nil, errors.New("系统错误")
	}

	// 4.验证密码
	encryptedInput := utility.EncryptPassword(password, user.UserSalt)
	if encryptedInput != user.Password {
		return "", 0, nil, errors.New("password error")
	}

	// 5.生成JWT Token
	token, expireTime, err := utility.GenerateToken(ctx, user.Id)
	if err != nil {
		return "", 0, nil, errors.New("生成token错误")
	}

	return token, int(expireTime.Sub(time.Now()).Seconds()), &user, nil
}

func Register(ctx context.Context, req *entity.UserInfo) (*entity.UserInfo, error) {
	// 1. 参数校验
	if req.Name == "" {
		return nil, errors.New("用户名不能为可空")
	}
	if len(req.Password) < 6 {
		return nil, errors.New("密码长度至少6位")
	}

	// 2. 检查用户名是否已存在
	count, err := dao.UserInfo.Ctx(ctx).Where("name", req.Name).Count()
	if err != nil {
		return nil, errors.New("检查用户名失败")
	}
	if count > 0 {
		return nil, errors.New("用户名已存在")
	}

	// 3.生成随机盐值（10位）
	req.UserSalt = utility.GenerateSalt(10)

	// 4.使用双重MD5加密密码
	req.Password = utility.EncryptPassword(req.Password, req.UserSalt)

	// 5.设置默认值
	now := gtime.Now()
	req.Status = 1
	req.CreatedAt = now
	req.UpdatedAt = now

	// 6.保存到数据库
	id, err := dao.UserInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "创建用户失败：%v", err)
		return nil, errors.New("创建用户失败")
	}

	// 7.设置ID并返回
	req.Id = int(id)
	return req, nil
}

func UpdatePassword(ctx context.Context, userId int, newPassword, secretAnswer string) error {
	// 1. 查询用户
	userRecord, err := dao.UserInfo.Ctx(ctx).Where("id", userId).One()
	if err != nil {
		g.Log().Errorf(ctx, "查询用户失败：%v", err)
		return errors.New("系统错误")
	}
	if userRecord.IsEmpty() {
		return errors.New("用户不存在")
	}

	// 2. 转换为实体
	var user entity.UserInfo
	if err := userRecord.Struct(&user); err != nil {
		g.Log().Errorf(ctx, "用户数据解析失败：%v", err)
		return errors.New("系统错误")
	}

	// 3. 验证密保答案
	if user.SecretAnswer != secretAnswer {
		return errors.New("密保答案错误")
	}

	// 4. 生成新密码
	newEncryptedPassword := utility.EncryptPassword(newPassword, user.UserSalt)

	// 5. 更新密码
	_, err = dao.UserInfo.Ctx(ctx).Where("id", userId).Update(g.Map{
		"password":   newEncryptedPassword,
		"updated_at": gtime.Now(),
	})
	if err != nil {
		g.Log().Errorf(ctx, "更新密码失败：%v", err)
		return errors.New("系统错误")
	}

	return nil
}

func GetUserInfo(ctx context.Context, userId int) (*entity.UserInfo, error) {
	// 1. 查询用户
	userRecord, err := dao.UserInfo.Ctx(ctx).Where("id", userId).One()
	if err != nil {
		g.Log().Errorf(ctx, "查询用户失败：%v", err)
		return nil, errors.New("系统错误")
	}
	if userRecord.IsEmpty() {
		return nil, errors.New("用户不存在")
	}

	var user entity.UserInfo
	if err := userRecord.Struct(&user); err != nil {
		g.Log().Errorf(ctx, "用户数据解析失败：%v", err)
		return nil, errors.New("系统错误")
	}

	return &user, nil
}
