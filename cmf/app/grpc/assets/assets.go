/**
** @创建时间: 2022/3/5 13:16
** @作者　　: return
** @描述　　:
 */

package assets

import (
	"context"
	"github.com/gincmf/bootstrap/util"
)

type Assets struct {
	UnimplementedAssetsServer
}

func (s *Assets) GetPrevPath(ctx context.Context, in *AssetsRequest) (assetsReply *AssetsReply, err error) {
	filePath := in.FilePath
	if filePath != "" {
		assetsReply = &AssetsReply{
			Data: &Data{
				PrevPath: util.FileUrl(filePath),
			},
		}
	}
	return
}
