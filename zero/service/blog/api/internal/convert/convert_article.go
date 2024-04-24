package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

func ConvertArticleTypes(in *blog.Article) (out *types.Article) {
	jsonconv.ObjectMarshal(in, &out)

	return
}

func ConvertArticlePb(in *types.Article) (out *blog.Article) {
	jsonconv.ObjectMarshal(in, &out)
	return
}

func ConvertArticleDetailsTypes(in *blog.Article) (out *types.ArticleDetailsResp) {
	jsonconv.ObjectMarshal(in, &out)
	return
}
