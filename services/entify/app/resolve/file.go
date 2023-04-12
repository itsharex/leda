package resolve

import (
	"fmt"

	"codebdy.com/leda/services/entify/config"
	"codebdy.com/leda/services/entify/consts"
	"codebdy.com/leda/services/entify/contexts"
	"codebdy.com/leda/services/entify/storage"
	"codebdy.com/leda/services/entify/utils"
	"github.com/graphql-go/graphql"
)

func GetFileUrl(fileInfo storage.FileInfo, p graphql.ResolveParams) (interface{}, error) {
	if config.Storage() == consts.LOCAL {
		return fmt.Sprintf(
			"http://%s/%s/%s",
			contexts.Values(p.Context).Host,
			consts.STATIC_PATH,
			fileInfo.Path,
		), nil
	} else {
		return fileInfo.Path, nil
	}
}

func FileUrlResolve(p graphql.ResolveParams) (interface{}, error) {
	defer utils.PrintErrorStack()
	if p.Source != nil {
		fileInfo := p.Source.(storage.FileInfo)
		return GetFileUrl(fileInfo, p)
	}
	return nil, nil
}
