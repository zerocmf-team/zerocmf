package role

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"zerocmf/service/lowcode/model"

	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.UserAdminRoleShow) (resp types.Response) {

	c := l.svcCtx
	siteId, _ := c.Get("siteId")

	// 选择租户表
	db, err := c.MongoDB(siteId.(string))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	var e *casbin.Enforcer
	e, err = db.NewEnforcer()

	collection := db.Collection("role")

	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		fmt.Println("Invalid ObjectID:", err)
		return
	}

	role := model.Role{}
	err = db.FindOne(collection, bson.M{
		"_id": objectID,
	}, &role)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	if role.Id.IsZero() {
		resp.Error("该角色不存在！", nil)
		return
	}

	var result struct {
		Role   model.Role `json:"role"`
		Access []string   `json:"access"`
	}

	result.Role = role

	id := role.Id.Hex()

	policy := e.GetFilteredPolicy(0, id)
	access := make([]string, 0)
	for _, v := range policy {
		if len(v) > 1 {
			access = append(access, v[1])
		}
	}
	result.Access = access
	resp.Success("获取成功！", result)
	return
}
