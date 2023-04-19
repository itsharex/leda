package app

import (
	"errors"
	"log"
	"sync"

	"codebdy.com/leda/services/models/consts"
	"codebdy.com/leda/services/models/entify/model"
	"codebdy.com/leda/services/models/entify/model/graph"
	"codebdy.com/leda/services/models/entify/model/meta"
	"codebdy.com/leda/services/models/modules/app/schema"
	"codebdy.com/leda/services/models/modules/app/schema/parser"
	"codebdy.com/leda/services/models/service"
)

//节省开支，运行时使用，初始化时请使用orm.IsEntityExists
var Installed = false

type App struct {
	AppId  uint64
	Model  *model.Model
	Schema schema.AppGraphqlSchema
	Parser *parser.ModelParser
}

type AppLoader struct {
	appId  uint64
	app    *App
	loaded bool
	sync.Mutex
}

func (l *AppLoader) load(force bool) {
	l.Lock()
	defer l.Unlock()
	if !l.loaded || force {
		log.Println("加载", l.appId)
		l.app = NewApp(l.appId)
		if l.app == nil {
			log.Panic(errors.New("Cant load app"))
		}
		l.loaded = true
	}
}

func Get(appId uint64) (*App, error) {
	if appId == 0 {
		return GetSystemApp(), nil
	}

	if result, ok := appLoaderCache.Load(appId); ok {
		if !result.(*AppLoader).loaded {
			result.(*AppLoader).load(false)
		}
		return result.(*AppLoader).app, nil
	} else {
		appLoader := &AppLoader{
			appId:  appId,
			loaded: false,
		}
		appLoaderCache.Store(appId, appLoader)
		appLoader.load(false)
		if appId == 0 {
			model.SystemModel = appLoader.app.Model
		}
		return appLoader.app, nil
	}
}

func (a *App) GetEntityByName(name string) *graph.Entity {
	return a.Model.Graph.GetEntityByName(name)
}

func (a *App) GetEntityByInnerId(innerId uint64) *graph.Entity {
	return a.Model.Graph.GetEntityByInnerId(innerId)
}

func NewApp(appId uint64) *App {
	systemApp := GetSystemApp()
	if appId == 0 {
		return systemApp
	}

	s := service.NewSystem()
	appData := s.QueryById(
		systemApp.GetEntityByName(meta.APP_ENTITY_NAME),
		appId,
	)

	if appData == nil {
		return nil
	}
	appMeta := s.QueryById(
		systemApp.GetEntityByName(meta.META_ENTITY_NAME),
		appData.(map[string]interface{})["metaId"].(uint64),
	)

	if appMeta == nil {
		return nil
	}

	publishedMeta := appMeta.(map[string]interface{})[consts.META_PUBLISHED_CONTENT]
	var content *meta.MetaContent
	if publishedMeta != nil {
		content = DecodeContent(publishedMeta)
	}

	content = MergeServiceModels(content)

	model := model.New(content, appId)
	schema := schema.New(model)

	return &App{
		AppId:  appId,
		Model:  model,
		Schema: schema,
		Parser: schema.Parser(),
	}
}
