/**
** @创建时间: 2021/11/24 12:52
** @作者　　: return
** @描述　　:
 */

package migrate

import (
	"context"
	pb "gincmf/app/grpc/migrate"
	"google.golang.org/grpc"
	"log"
	"time"
)

func init()  {
	//bootstrap.Db().AutoMigrate(Tenant{}) // 租户数据库迁移

	conn, err := grpc.Dial("localhost:8801", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMigrateClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Migrate(ctx, &pb.MigrateRequest{TenantId: "666666"})
	if err != nil {
		log.Fatalf("could not migrate: %v", err)
	}
	log.Printf("Success: %s", r.GetMessage())

}
