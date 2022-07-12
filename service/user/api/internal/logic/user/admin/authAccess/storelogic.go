package authAccess

import (
	"context"
	"errors"
	"zerocmf/common/bootstrap/casbin"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/user/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"strconv"
	"time"

	"zerocmf/service/user/api/internal/svc"
	"zerocmf/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreLogic {
	return &StoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreLogic) Store(req *types.AccessStore) (resp types.Response) {
	form := access{}
	copier.Copy(&form, &req)
	role, err := save(form, l.svcCtx)

	if err != nil {
		resp.Error(err.Error(),nil)
		return
	}
	resp.Success("操作成功！",role)
	return
}

type access struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Remark     string   `json:"remark"`
	RoleAccess []string `json:"role_access"`
}

func save(req access, c *svc.ServiceContext) (result model.Role, err error) {

	form := req
	db := c.Db
	// 角色信息
	role := model.Role{
		Name:     form.Name,
		Remark:   form.Remark,
		CreateAt: time.Now().Unix(),
		Status:   1,
	}
	e, err := casbin.NewEnforcer("")
	//	存入casbin
	if err != nil {
		return
	}

	editId := req.Id

	if editId == "" {
		tx := db.Create(&role)
		if tx.Error != nil {
			err = tx.Error
			return
		}
		id := strconv.Itoa(role.Id)
		for _, v := range form.RoleAccess {
			e.AddPolicy(id, v,"*")
		}
	} else {
		roleItem := model.Role{}
		tx := db.Where("id = ? AND status = 1", editId).First(&roleItem)
		if tx.Error != nil {
			if tx.Error == gorm.ErrRecordNotFound {
				err = errors.New("该角色不存在或已删除！")
				return
			}
			err = tx.Error
			return
		}
		role.Id = roleItem.Id
		tx = db.Save(&role)
		if tx.Error != nil {
			err = tx.Error
			return
		}

		id := strconv.Itoa(role.Id)

		// 新增修改策略

		// 获取全部角色策略
		roles := e.GetFilteredPolicy(0, id)
		existAccess := make([]string, 0)
		for _, v := range roles {
			if len(v) > 1 {
				existAccess = append(existAccess, v[1])
			}
		}

		/*
		*  新增：[1,2,3]
		*  原有：[3,4,5]
		*  筛选去除的规则：[4，5]
		*  筛选新增的规则：[1，2]

		* 新增：[1,2,3]
		* 原有：[]
		* 筛选去除的规则：[]
		* 筛选新增的规则：[1，2，3]

		* 新增：[]
		* 原有：[1，2，3]
		* 筛选去除的规则：[1，2，3]
		* 筛选新增的规则：[]
		 */

		alreadyDel := make([]string, 0)
		// 判断是否需要被删除
		for _, v := range existAccess {
			if util.ToLowerInArray(v, form.RoleAccess) == false {
				alreadyDel = append(alreadyDel, v)
			}
		}
		// 如果新增为空，则全部删除
		if len(form.RoleAccess) == 0 {
			alreadyDel = existAccess
		}

		alreadyAdd := make([]string, 0)
		for _, v := range form.RoleAccess {
			if util.ToLowerInArray(v, existAccess) == false {
				alreadyAdd = append(alreadyAdd, v)
			}
		}

		// 如果数据库不存在，则为新增
		if len(existAccess) == 0 {
			alreadyAdd = form.RoleAccess
		}

		// 开始删除策略
		rules := make([][]string, 0)
		for _, v := range alreadyDel {
			rules = append(rules, []string{id, v,"*"})
		}
		if len(rules) > 0 {
			e.RemovePolicies(rules)
		}

		// 开始新增策略
		rules = make([][]string, 0)
		for _, v := range alreadyAdd {
			rules = append(rules, []string{id, v,"*"})
		}
		if len(rules) > 0 {
			e.AddPolicies(rules)
		}

	}

	result = role

	return
}
