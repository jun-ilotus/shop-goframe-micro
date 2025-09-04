package admin_info

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"service/app/admin/internal/dao"
	"service/app/admin/internal/model/entity"
	"service/utility"
	"time"
)

func Login(ctx context.Context, name, password string) (token string, expire time.Time, err error) {
	if name == "" || password == "" {
		return "", time.Time{}, errors.New("name or password is empty")
	}

	adminRecord, err := dao.AdminInfo.Ctx(ctx).Where("name", name).One()
	if err != nil {
		return "", time.Time{}, errors.New("系统错误")
	}
	if adminRecord.IsEmpty() {
		return "", time.Time{}, errors.New("用户不存在")
	}

	var admin entity.AdminInfo
	if err = adminRecord.Struct(&admin); err != nil {
		g.Log().Errorf(ctx, "用户数据解析失败：%v", err)
		return "", time.Time{}, errors.New("系统错误")
	}

	encryptedInput := utility.EncryptPassword(password, admin.UserSalt)
	if encryptedInput != admin.Password {
		return "", time.Time{}, errors.New("密码错误")
	}

	return utility.GenerateToken(uint32(admin.Id))
}

func Register(ctx context.Context, name, password string) (*entity.AdminInfo, error) {
	if name == "" {
		return nil, errors.New("用户名不能为空")
	}
	if len(password) < 6 {
		return nil, errors.New("密码长度至少为6位")
	}

	count, err := dao.AdminInfo.Ctx(ctx).Where("name", name).Count()
	if err != nil {
		return nil, errors.New("检查用户名失败")
	}
	if count > 0 {
		return nil, errors.New("用户名已存在")
	}

	salt := utility.GenerateSalt(10)
	encryptedPassword := utility.EncryptPassword(password, salt)

	now := gtime.Now()
	admin := &entity.AdminInfo{
		Name:      name,
		Password:  encryptedPassword,
		RoleIds:   "2",
		UserSalt:  salt,
		IsAdmin:   0,
		CreatedAt: now,
		UpdatedAt: now,
	}

	id, err := dao.AdminInfo.Ctx(ctx).InsertAndGetId(admin)
	if err != nil {
		g.Log().Errorf(ctx, "创建用户失败: %v", err)
		return nil, errors.New("创建用户失败")
	}
	admin.Id = int(id)
	return admin, nil
}
