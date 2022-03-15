/**
** @创建时间: 2021/11/23 20:27
** @作者　　: return
** @描述　　: 平台租户表
 */

package model

import "gincmf/app/migrate"

/**
 * @Author return <1140444693@qq.com>
 * @Description 租户模型数据
 * @Date 2021/11/24 12:44:54
 * @Param
 * @return
 **/

type Tenant struct {
	migrate.Tenant
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 租户状态日志，包含创建，到期，购买等记录
 * @Date 2021/11/24 12:45:17
 * @Param
 * @return
 **/
