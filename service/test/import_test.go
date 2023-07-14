package test

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"reflect"
	"testing"
	"time"
	"zerocmf/common/bootstrap/database"
	bsModel "zerocmf/common/bootstrap/model"
	"zerocmf/service/lowcode/model"
	model2 "zerocmf/service/user/model"
)

type hospital struct {
	HospitalName         string `json:"hospital_name"`
	HospitalIntroduction string `json:"hospital_introduction"`
	Province             string `json:"province"`
	City                 string `json:"city"`
	District             string `json:"district"`
	HospitalAddress      string `json:"hospital_address"`
	HospitalPhone        string `json:"hospital_phone"`
	HospitalSelling      string `json:"hospital_selling"`
	HospitalWebsite      string `json:"hospital_website"`
}

type mapField struct {
	ImportName string
	ExportName string
	Value      string
	DoFunc     func(record map[string]interface{}) interface{}
}

func ConvertToInt(i interface{}) (int, error) {
	switch v := i.(type) {
	case int:
		return v, nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	default:
		return 0, fmt.Errorf("无法将接口转换为int类型，接收到的类型为 %v", reflect.TypeOf(i))
	}
}

func TestImportHospital(t *testing.T) {
	dbConf := database.Database{
		Type:     "mysql",
		Host:     "101.132.166.55",
		Database: "kf.iximei.cn",
		Username: "root",
		Password: "tdx2023.",
		Port:     3306,
		Charset:  "utf8mb4",
		Prefix:   "hj_",
	}
	db := database.NewGormDb(dbConf)

	var h []map[string]interface{}
	db.Table("hj_hospital").Find(&h)

	mgoDb, _ := database.NewMongoDB(database.Mongo{
		Host:     "localhost",
		DbName:   "tenant",
		Username: "root",
		Password: "123456",
		Port:     27017,
	}, "123")

	collection := mgoDb.Collection("formData")

	objectID, _ := primitive.ObjectIDFromHex("646c5ede34ec1add92a36551")

	newDbConf := database.Database{
		Type:     "mysql",
		Host:     "localhost",
		Username: "root",
		Database: "site_123_user",
		Password: "123456",
		Port:     3306,
		Charset:  "utf8mb4",
		Prefix:   "cmf_",
		AuthCode: "KFHlk2ubIlMr5ltqaD",
	}

	newDb := database.NewGormDb(newDbConf)

	var mapFields = []mapField{
		{
			ImportName: "hospital_name",
			ExportName: "userLogin",
		},
		{
			ExportName: "userPass",
			DoFunc: func(record map[string]interface{}) interface{} {
				return "123456"
			},
		},
		{
			ExportName: "userId",
			DoFunc: func(record map[string]interface{}) interface{} {
				return "1"
			},
		},
		{
			ImportName: "hospital_name",
			ExportName: "hospitalName",
		},
		{
			ImportName: "hospital_phone",
			ExportName: "hospitalTel",
		},
		{
			ExportName: "cascader",
			DoFunc: func(record map[string]interface{}) interface{} {
				one := record["province"]
				two := record["city"]
				three := record["district"]

				oneInt, _ := ConvertToInt(one)
				twoInt, _ := ConvertToInt(two)
				threeInt, _ := ConvertToInt(three)

				return []int{oneInt, twoInt, threeInt}
			},
		},
		{
			ImportName: "hospital_address",
			ExportName: "hospitalAddress",
		},
		{
			ImportName: "hospital_selling",
			ExportName: "hospitalAvgPrice",
		},
		{
			ImportName: "hospital_website",
			ExportName: "hospitalWebsite",
		},
		{
			ImportName: "hospital_nature",
			ExportName: "hospitalNature",
		},
		{
			ImportName: "doctor_name",
			ExportName: "doctorName",
		},
		{
			ImportName: "doctor_phone",
			ExportName: "doctorTel",
		},
		{
			ImportName: "doctor_qq",
			ExportName: "doctorQQ",
		},
		{
			ImportName: "reception_name",
			ExportName: "receptionName",
		},
		{
			ImportName: "reception_phone",
			ExportName: "receptionTel",
		},
		{
			ImportName: "reception_qq",
			ExportName: "receptionQQ",
		},
		{
			ImportName: "bus_station",
			ExportName: "busStation",
		},
		{
			ImportName: "bus_address",
			ExportName: "busAddress",
		},
		{
			ImportName: "subway_station",
			ExportName: "subwayStation",
		},
		{
			ImportName: "subway_address",
			ExportName: "subwayAddress",
		},
		{
			ImportName: "taxi_fare",
			ExportName: "taxiFare",
		},
		{
			ImportName: "vip_discount",
			ExportName: "vipDiscount",
		},
		{
			ImportName: "return_point",
			ExportName: "returnPoint",
		},
		{
			ImportName: "hospital_introduction",
			ExportName: "hospitalDesc",
		},
	}

	var bulkOps []mongo.WriteModel
	for _, v := range h {
		filter := bson.M{"formId": objectID, "schema.fieldId": "hospitalName", "schema.fieldData.value": v["hospital_name"]}
		var columns = make([]model.ColumnsProps, len(mapFields))

		for k1, field := range mapFields {

			var value = v[field.ImportName]
			if field.DoFunc != nil {
				value = field.DoFunc(v)
			}

			columns[k1].FieldId = field.ExportName
			columns[k1].Label = ""
			columns[k1].ComponentName = "Form.Item"
			columns[k1].FieldData = &model.FieldData{
				Text:  "",
				Value: value,
			}
		}

		t1, _ := ConvertToInt(v["create_time"])
		t2, _ := ConvertToInt(v["update_time"])

		var createAt = int64(t1)
		var updateAt = int64(t2)

		update := bson.M{
			"$set": model.FormData{
				FormId: objectID,
				Schema: columns,
				User: model2.User{
					Id:        1,
					UserLogin: "admin",
				},
				CreateAt: createAt,
				UpdateAt: updateAt,
			},
		}

		upsert := true
		updateModel := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(upsert)
		bulkOps = append(bulkOps, updateModel)

		if v["hospital_name"] != "" {
			user := model2.User{
				UserType:  1,
				UserLogin: v["hospital_name"].(string),
				UserPass:  "3ffc0ef0ced4e6824cc61b5afdedcff4",
			}
			tx := newDb.Where("user_login = ?", v["hospital_name"]).FirstOrCreate(&user)

			if tx.Error != nil {
				continue
			}

			userRole := model2.RoleUser{
				RoleId: 2,
				UserId: user.Id,
			}
			newDb.Where("role_id = ? AND user_id = ?", 3, user.Id).FirstOrCreate(&userRole)

		}

	}

	// 执行BulkWrite操作
	bulkWriteResult, err := collection.BulkWrite(context.Background(), bulkOps)
	if err != nil {
		log.Fatal(err)
	}

	// 打印插入和更新的文档数目
	fmt.Printf("Inserted %d documents\n", bulkWriteResult.InsertedCount)
	fmt.Printf("Updated %d documents\n", bulkWriteResult.ModifiedCount)

}

func TestImportCustom(t *testing.T) {
	dbConf := database.Database{
		Type:     "mysql",
		Host:     "101.132.166.55",
		Database: "kf.iximei.cn",
		Username: "root",
		Password: "tdx2023.",
		Port:     3306,
		Charset:  "utf8mb4",
		Prefix:   "hj_",
	}
	db := database.NewGormDb(dbConf)
	var importData []map[string]interface{}
	db.Table("hj_custom").Find(&importData)

	mgoDb, _ := database.NewMongoDB(database.Mongo{
		Host:     "124.223.74.180",
		DbName:   "tenant",
		Username: "root",
		Password: "123456",
		Port:     27017,
	}, "123")

	collection := mgoDb.Collection("formData")

	objectID, _ := primitive.ObjectIDFromHex("646c5f0034ec1add92a36559")

	var mapFields = []mapField{
		{
			ImportName: "id",
			ExportName: "oid",
		},
		{
			ImportName: "name",
			ExportName: "name",
		},
		{
			ImportName: "birthday",
			ExportName: "birthday",
		},
		{
			ImportName: "gender",
			ExportName: "gender",
		},
		{
			ImportName: "address",
			ExportName: "address",
		},
		{
			ImportName: "division",
			ExportName: "division",
		},
		{
			ImportName: "address",
			ExportName: "address",
		},
		{
			ImportName: "telphone",
			ExportName: "tel",
		},
		{
			ImportName: "mobile",
			ExportName: "phone",
		},
		{
			ImportName: "qq",
			ExportName: "qq",
		},
		{
			ImportName: "wechat",
			ExportName: "wechat",
		},
		{
			ImportName: "project",
			ExportName: "project",
		},
		{
			ImportName: "plastic",
			ExportName: "status",
		},
		{
			ImportName: "remark",
			ExportName: "remark",
		},
	}

	var bulkOps []mongo.WriteModel
	for _, v := range importData {
		filter := bson.M{"formId": objectID, "schema.fieldId": "oid", "schema.fieldData.value": v["id"]}
		var columns = make([]model.ColumnsProps, len(mapFields))

		for k1, field := range mapFields {

			var value = v[field.ImportName]
			if field.DoFunc != nil {
				value = field.DoFunc(v)
			}

			columns[k1].FieldId = field.ExportName
			columns[k1].Label = ""
			columns[k1].ComponentName = "Form.Item"
			columns[k1].FieldData = &model.FieldData{
				Text:  "",
				Value: value,
			}
		}

		t1, _ := ConvertToInt(v["create_time"])
		t2, _ := ConvertToInt(v["update_time"])

		var createAt = int64(t1)
		var updateAt = int64(t2)

		update := bson.M{
			"$set": model.FormData{
				FormId: objectID,
				Schema: columns,
				User: model2.User{
					Id:        1,
					UserLogin: "admin",
				},
				CreateAt: createAt,
				UpdateAt: updateAt,
			},
		}

		upsert := true
		updateModel := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(upsert)
		bulkOps = append(bulkOps, updateModel)
	}

	// 执行BulkWrite操作
	bulkWriteResult, err := collection.BulkWrite(context.Background(), bulkOps)
	if err != nil {
		log.Fatal(err)
	}

	// 打印插入和更新的文档数目
	fmt.Printf("Inserted %d documents\n", bulkWriteResult.InsertedCount)
	fmt.Printf("Updated %d documents\n", bulkWriteResult.ModifiedCount)
}

func TestImportDispatch(t *testing.T) {
	dbConf := database.Database{
		Type:     "mysql",
		Host:     "101.132.166.55",
		Database: "kf.iximei.cn",
		Username: "root",
		Password: "tdx2023.",
		Port:     3306,
		Charset:  "utf8mb4",
		Prefix:   "hj_",
	}
	db := database.NewGormDb(dbConf)

	var h []map[string]interface{}
	db.Debug().Select("d.*,h.hospital_name,c.name as custom_name,c.birthday,c.plastic").Table("hj_dispatch d").
		Joins(" INNER JOIN hj_hospital h ON d.hospital_id = h.id").
		Joins(" INNER JOIN hj_custom c ON d.custom_id = c.id").Order("d.id").Find(&h)

	mgoDb, _ := database.NewMongoDB(database.Mongo{
		Host:     "localhost",
		DbName:   "tenant",
		Username: "root",
		Password: "123456",
		Port:     27017,
	}, "123")

	collection := mgoDb.Collection("formData")

	objectID, _ := primitive.ObjectIDFromHex("646c5f2e34ec1add92a36561")

	var mapFields = []mapField{
		{
			ImportName: "id",
			ExportName: "importId",
		},
		{
			DoFunc: func(record map[string]interface{}) interface{} {
				//fmt.Println("record", record)
				formID, _ := primitive.ObjectIDFromHex("646c5f0034ec1add92a36559")
				// 查询新系统的客户id
				var formData model.FormData

				mgoDb.FindOne(collection, bson.M{
					"formId":                 formID,
					"schema.fieldId":         "oid",
					"schema.fieldData.value": record["id"],
				}, &formData)

				return formData.Id
			},
			ExportName: "customerId",
		},
		{
			DoFunc: func(record map[string]interface{}) interface{} {
				formID, _ := primitive.ObjectIDFromHex("646c5ede34ec1add92a36551")
				// 查询新系统的客户id
				var formData model.FormData

				mgoDb.FindOne(collection, bson.M{
					"formId":                 formID,
					"schema.fieldId":         "hospitalName",
					"schema.fieldData.value": record["hospital_name"],
				}, &formData)

				return formData.Id
			},
			ExportName: "hospitalId",
		},
		{
			ImportName: "hospital_name",
			ExportName: "hospitalName",
		},
		{
			ImportName: "custom_name",
			ExportName: "customerName",
		},
		{
			ImportName: "",
			ExportName: "customerAge",
		},
		{
			ImportName: "plastic",
			ExportName: "project",
		},
		{
			ImportName: "",
			ExportName: "remark",
		},
		{
			ImportName: "",
			ExportName: "hospitalWechat",
		},
		{
			ImportName: "",
			ExportName: "hospitalQQ",
		},
		{
			ImportName: "",
			ExportName: "picture",
		},
		{
			ImportName: "",
			ExportName: "content",
		},
	}

	var bulkOps []mongo.WriteModel

	for _, v := range h {

		filter := bson.M{"formId": objectID, "schema.fieldId": "importId", "schema.fieldData.value": v["id"]}
		var columns = make([]model.ColumnsProps, len(mapFields))

		for k1, field := range mapFields {

			var value = v[field.ImportName]

			if field.DoFunc != nil {
				value = field.DoFunc(v)
			}

			columns[k1].FieldId = field.ExportName
			columns[k1].Label = ""
			columns[k1].ComponentName = "Form.Item"
			columns[k1].FieldData = &model.FieldData{
				Text:  "",
				Value: value,
			}
		}

		t1, _ := ConvertToInt(v["create_time"])
		t2, _ := ConvertToInt(v["update_time"])

		var createAt = int64(t1)
		var updateAt = int64(t2)

		update := bson.M{
			"$set": model.FormData{
				FormId: objectID,
				Schema: columns,
				User: model2.User{
					Id:        1,
					UserLogin: "admin",
				},
				CreateAt: createAt,
				UpdateAt: updateAt,
			},
		}

		upsert := true

		updateModel := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(upsert)
		bulkOps = append(bulkOps, updateModel)

	}

	// 执行BulkWrite操作
	bulkWriteResult, err := collection.BulkWrite(context.Background(), bulkOps)
	if err != nil {
		log.Fatal(err)
	}

	// 打印插入和更新的文档数目
	fmt.Println("Inserted documents", bulkWriteResult.InsertedCount)
	fmt.Printf("Updated %d documents\n", bulkWriteResult.ModifiedCount)

}

func TestImportMongo(t *testing.T) {

	mgoDb, _ := database.NewMongoDB(database.Mongo{
		Host:     "localhost",
		Username: "root",
		Password: "123456",
		Port:     27017,
	}, "439226114")

	// 顶级
	parentId := saveOne(mgoDb, "账号管理", primitive.ObjectID{})

	// 二级
	roleParentId := saveOne(mgoDb, "角色设置", parentId)

	// 三级
	for k, v := range []string{"角色列表", "添加角色", "编辑角色"} {

		update := model.Form{
			Name:       v,
			ParentId:   roleParentId,
			HideInMenu: 1,
			ListOrder:  10000,
			Status:     1,
			Time: bsModel.Time{
				CreateAt: time.Now().Unix(),
				UpdateAt: time.Now().Unix(),
			},
		}
		filePath := "/Users/return/workspace/mygo/zerocmf/service/lowcode/model/template/admin/role/role.json"
		if k == 1 {
			filePath = "/Users/return/workspace/mygo/zerocmf/service/lowcode/model/template/admin/role/addRole.json"
		} else if k == 2 {
			filePath = "/Users/return/workspace/mygo/zerocmf/service/lowcode/model/template/admin/role/editRole.json"
		}
		importOne(mgoDb, update, filePath)
	}

	update := model.Form{
		Name:       "部门管理",
		ParentId:   parentId,
		HideInMenu: 0,
		ListOrder:  10000,
		Status:     1,
		Time: bsModel.Time{
			CreateAt: time.Now().Unix(),
			UpdateAt: time.Now().Unix(),
		},
	}
	filePath := "/Users/return/workspace/mygo/zerocmf/service/lowcode/model/template/admin/department.json"
	importOne(mgoDb, update, filePath)

	adminParentId := saveOne(mgoDb, "管理员设置", parentId)

	for k, v := range []string{"管理员列表", "添加管理员", "编辑管理员"} {
		update = model.Form{
			Name:       v,
			ParentId:   adminParentId,
			HideInMenu: 1,
			ListOrder:  10000,
			Status:     1,
			Time: bsModel.Time{
				CreateAt: time.Now().Unix(),
				UpdateAt: time.Now().Unix(),
			},
		}
		filePath = "/Users/return/workspace/mygo/zerocmf/service/lowcode/model/template/admin/admin/list.json"
		if k == 1 {
			filePath = "/Users/return/workspace/mygo/zerocmf/service/lowcode/model/template/admin/admin/add.json"
		} else if k == 2 {
			filePath = "/Users/return/workspace/mygo/zerocmf/service/lowcode/model/template/admin/admin/edit.json"
		}
		importOne(mgoDb, update, filePath)
	}

	// 顶级
	settingParentId := saveOne(mgoDb, "系统设置", primitive.ObjectID{})
	for k, v := range []string{"站点设置", "菜单设置"} {
		update = model.Form{
			Name:       v,
			ParentId:   settingParentId,
			HideInMenu: 0,
			ListOrder:  10000,
			Status:     1,
			Time: bsModel.Time{
				CreateAt: time.Now().Unix(),
				UpdateAt: time.Now().Unix(),
			},
		}
		filePath = "/Users/return/workspace/mygo/zerocmf/service/lowcode/model/template/admin/site.json"
		if k == 1 {
			filePath = "/Users/return/workspace/mygo/zerocmf/service/lowcode/model/template/admin/menu.json"
		}
		importOne(mgoDb, update, filePath)
	}

}

func saveOne(mgoDb database.MongoDB, name string, parentId primitive.ObjectID) (id primitive.ObjectID) {
	collection := mgoDb.Collection("form")

	update := model.Form{
		Name:     name,
		ParentId: parentId,
		Schema:   "",
	}

	form := model.Form{}
	filter := bson.M{"name": update.Name}
	mgoDb.FindOne(collection, filter, &form)

	if form.Id.IsZero() {
		update.CreateAt = time.Now().Unix()
		update.UpdateAt = time.Now().Unix()
		one, err := mgoDb.InsertOne(collection, &update)
		if err != nil {
			panic(err)
		}
		id = one.InsertedID.(primitive.ObjectID)
	} else {
		update.UpdateAt = time.Now().Unix()
		_, err := mgoDb.UpdateOne(collection, filter, bson.M{
			"$set": update,
		})
		if err != nil {
			panic(err)
		}
		id = form.Id
	}
	return
}

func importOne(mgoDb database.MongoDB, update model.Form, filePath string) {

	filter := bson.M{"name": update.Name}

	form := model.Form{}
	collection := mgoDb.Collection("form")
	mgoDb.FindOne(collection, filter, &form)

	// 导入菜单

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("无法读取文件：%s", err)
	}

	var schemaData model.Schema
	formSchema := data
	json.Unmarshal(formSchema, &schemaData)
	formComponents := model.FindComponents(schemaData.ComponentsTree, "Form")
	components := model.FindComponents(formComponents, "Form.Item")

	var columns []model.ColumnsProps
	for _, component := range components {

		props := component.Props

		column := model.ColumnsProps{
			FieldId:       props.Name,
			Label:         props.Label,
			Unique:        props.Unique,
			ComponentName: component.ComponentName,
			Rules:         props.Rules,
		}

		if column.FieldId != "" {
			columns = append(columns, column)
		}
	}

	if form.Id.IsZero() {
		//新增表单
		update.Columns = columns
		update.Schema = string(data)
		//	选择表单
		var one *mongo.InsertOneResult
		one, err = mgoDb.InsertOne(collection, update)
		update.Id = one.InsertedID.(primitive.ObjectID)
	} else {
		update.Columns = columns
		update.Schema = string(data)
		update.UpdateAt = time.Now().Unix()
		mgoDb.UpdateOne(collection, filter, bson.M{
			"$set": update,
		})
	}
}
