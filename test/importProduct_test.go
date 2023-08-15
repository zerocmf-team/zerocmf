package test

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"testing"
	"time"
	"zerocmf/common/bootstrap/database"
	"zerocmf/service/shop/model"

	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/xuri/excelize/v2"
)

// AssignStructFields 使用反射将数据映射到结构体的字段
func AssignStructFields(target interface{}, data map[string]interface{}) error {
	targetValue := reflect.ValueOf(target)

	if targetValue.Kind() != reflect.Ptr || targetValue.IsNil() {
		return errors.New("target must be a non-nil pointer to a struct")
	}

	targetValue = targetValue.Elem()

	if targetValue.Kind() != reflect.Struct {
		return errors.New("target must be a pointer to a struct")
	}

	for fieldName, fieldValue := range data {

		field := targetValue.FieldByName(fieldName)

		if !field.IsValid() {
			continue
			//return errors.New("field " + fieldName + " not found")
		}

		if !field.CanSet() {
			continue
			//return errors.New("field " + fieldName + " cannot be set")
		}

		fieldType := field.Type()
		val := reflect.ValueOf(fieldValue)

		typeString := fieldType.String()

		pattern := `^sql\.Null(\w+$)`

		// Compile the regular expression
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(typeString)

		if len(matches) > 1 {
			name := matches[1]
			if val.Type().String() == name {
				field.FieldByName(name).Set(val)
				validField := field.FieldByName("Valid")
				if validField.IsValid() && validField.CanSet() {
					validField.SetBool(true)
				}
			}

		} else if val.Type().ConvertibleTo(fieldType) {
			field.Set(val.Convert(fieldType))
		}

		switch fieldType.String() {
		case "sql.NullString":

		default:

		}

		//val := reflect.ValueOf(fieldValue)
		//
		//if !val.Type().ConvertibleTo(fieldType) {
		//	return errors.New("cannot assign field " + fieldName + " with given value")
		//}
		//if val.Type().ConvertibleTo(fieldType) {
		//	field.Set(val.Convert(fieldType))
		//}
	}

	return nil
}

func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

func TestImport(t *testing.T) {

	f, err := excelize.OpenFile("./商品模板.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	database := database.Database{
		Host:     "localhost",
		Database: "site_408764519_shop",
		Username: "root",
		Password: "123456",
		Port:     3306,
		Charset:  "utf8mb4",
		Prefix:   "cmf_",
		AuthCode: "KFHlk2ubIlMr5ltqaD",
	}
	dsn := database.Dsn()

	//mysql model调用
	conn := sqlx.NewMysql(dsn)

	c := cache.CacheConf{}
	c = append(c, cache.NodeConf{
		RedisConf: redis.RedisConf{
			Host:        "localhost:6379",
			Type:        "node",
			Tls:         false,
			NonBlock:    true,
			PingTimeout: 1,
		},
		Weight: 100,
	})

	now := time.Now().Unix()

	categoryDb := model.NewProductCategoryModel(conn, c)

	ctx := context.Background()

	categoryDb.Find(ctx)

	// 获取 分类 上所有单元格
	rows, err := f.GetRows("分类")
	if err != nil {
		fmt.Println(err)
		return
	}

	categoryMap := map[string]string{"分类名称": "Name", "分类描述": "Desc"}
	var categoryIndex []string

	categoryModelMap := map[string]model.ProductCategory{}

	for rowIndex, row := range rows {
		category := model.ProductCategory{}
		data := map[string]interface{}{}
		for colIndex, colCell := range row {
			if rowIndex == 0 {
				field := categoryMap[colCell]
				categoryIndex = append(categoryIndex, field)
			} else {
				field := categoryIndex[colIndex]
				data[field] = colCell
			}
		}

		if rowIndex == 0 {
			continue
		}

		err = AssignStructFields(&category, data)

		//err = copier.Copy(&category, &data)
		if err != nil {
			fmt.Println("err", err.Error())
			continue
		}

		// 查询当前分类是否存在
		one, err := categoryDb.Where("name = ?", category.Name).First(ctx)
		if err != nil && err != model.ErrNotFound {
			fmt.Println("err", err.Error())
			continue
		}

		// 更新
		if one != nil {
			category.UpdatedAt = now
			copier.CopyWithOption(&category, &one, copier.Option{IgnoreEmpty: true})
			categoryDb.Where("name = ?", category.Name).Update(ctx, &category)
		} else {
			category.CreatedAt = now
			category.UpdatedAt = now
			insert, err := categoryDb.Insert(ctx, &category)
			if err != nil {
				fmt.Println("err", err.Error())
				continue
			}
			category.ProductCategoryId, _ = insert.LastInsertId()
		}

		categoryModelMap[category.Name] = category

	}

	// 获取 商品 上所有单元格
	rows, err = f.GetRows("商品")
	if err != nil {
		fmt.Println(err)
		return
	}

	productMap := map[string]string{
		"商品名称": "ProductName",
		"商品分类": "ProductCategoryName",
		"商品条码": "ProductBarcode",
		"商品图片": "ProductThumbnail",
		"库存单位": "StockUnit",
		"零售价":  "Price",
		"标准价":  "OriginalPrice",
		"成本价":  "CostPrice",
		"价格面议": "PriceNegotiable",
		"隐藏库存": "HideRemainingStock",
		"分享描述": "ShareDescription",
		"商品卖点": "ProductSellingPoint",
	}

	productDb := model.NewProductModel(conn, c)

	var productIndex []string
	productModelMap := map[string]model.Product{}

	for rowIndex, row := range rows {
		product := model.Product{
			Attributes: "[]",
		}
		data := map[string]interface{}{}
		for colIndex, colCell := range row {
			if rowIndex == 0 {
				field := productMap[colCell]
				productIndex = append(productIndex, field)
			} else {
				field := productIndex[colIndex]
				data[field] = colCell
			}
		}

		if rowIndex == 0 {
			continue
		}

		err = AssignStructFields(&product, data)
		if err != nil {
			fmt.Println("err", err.Error())
			continue
		}

		categoryName := data["ProductCategoryName"].(string)
		productCategory := categoryModelMap[categoryName].ProductCategoryId

		product.ProductCategory = productCategory

		one, err := productDb.Where("product_name = ?", product.ProductName).First(ctx)
		if err != nil && err != model.ErrNotFound {
			fmt.Println("err", err.Error())
			continue
		}

		// 更新
		if one != nil {
			product.UpdatedAt = now
			copier.CopyWithOption(&one, &product, copier.Option{IgnoreEmpty: true})
			productDb.Where("product_name = ?", product.ProductName).Update(ctx, one)
		} else {
			product.CreatedAt = now
			product.UpdatedAt = now
			insert, err := productDb.Insert(ctx, &product)
			if err != nil {
				fmt.Println("err", err.Error())
				continue
			}
			product.ProductId, _ = insert.LastInsertId()
		}

		productModelMap[product.ProductName] = product

	}

	fmt.Println("productModelMap", productModelMap)
}
