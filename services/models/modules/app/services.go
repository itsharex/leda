package app

// var ServiceMetas sync.Map

// //加载微内核
// func LoadServiceMetas() {
// 	clearServiceMetas()

// 	s := service.NewSystem()

// 	services := s.QueryEntity(consts.SERVICE_ENTITY_NAME, graph.QueryArg{}, []string{"id", "metaId"})

// 	for _, service := range services.Nodes {
// 		metaIdData := service["metaId"]
// 		if metaIdData != nil && metaIdData != 0 {
// 			metaData := s.QueryById(consts.META_ENTITY_NAME, metaIdData.(uint64))
// 			if metaData != nil {
// 				metaMap := metaData.(map[string]interface{})
// 				publishedMeta := metaMap[consts.META_PUBLISHED_CONTENT]
// 				if publishedMeta != nil && publishedMeta != "" {
// 					var content *meta.UMLMeta
// 					if publishedMeta != nil {
// 						content = DecodeContent(publishedMeta)
// 					}
// 					ServiceMetas.Store(metaMap["id"], content)
// 				}
// 			}
// 		}
// 	}

// 	ReloadSystemApp()
// }

// func clearServiceMetas() {
// 	ServiceMetas.Range(func(key interface{}, value interface{}) bool {
// 		ServiceMetas.Delete(key)
// 		return true
// 	})
// }
