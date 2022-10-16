package app

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaoapp/gou"
	"github.com/yaoapp/gou/lang"
	"github.com/yaoapp/kun/any"
	"github.com/yaoapp/yao/config"
	"github.com/yaoapp/yao/flow"
	"github.com/yaoapp/yao/i18n"
	"github.com/yaoapp/yao/widgets/login"
)

func TestLoad(t *testing.T) {
	err := Load(config.Conf)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "::Demo Application", Setting.Name)
	assert.Equal(t, "::Demo", Setting.Short)
	assert.Equal(t, "::Another yao application", Setting.Description)
	assert.Equal(t, []string{"demo"}, Setting.Menu.Args)
	assert.Equal(t, "flows.app.menu", Setting.Menu.Process)
	assert.Equal(t, true, Setting.Optional.HideNotification)
	assert.Equal(t, false, Setting.Optional.HideSetting)
}

func TestLoadHK(t *testing.T) {

	err := i18n.Load(config.Conf)
	if err != nil {
		t.Fatal(err)
	}

	err = Load(config.Conf)
	if err != nil {
		t.Fatal(err)
	}

	newSetting, err := i18n.Trans("zh-hk", "app", "app", Setting)
	if err != nil {
		t.Fatal(err)
	}
	setting := newSetting.(*DSL)

	assert.Equal(t, "示例應用", setting.Name)
	assert.Equal(t, "演示", setting.Short)
	assert.Equal(t, "又一個YAO應用", setting.Description)
	assert.Equal(t, []string{"demo"}, setting.Menu.Args)
	assert.Equal(t, "flows.app.menu", setting.Menu.Process)
	assert.Equal(t, true, setting.Optional.HideNotification)
	assert.Equal(t, false, setting.Optional.HideSetting)

	assert.Equal(t, "::Demo Application", Setting.Name)
	assert.Equal(t, "::Demo", Setting.Short)
	assert.Equal(t, "::Another yao application", Setting.Description)
	assert.Equal(t, []string{"demo"}, Setting.Menu.Args)
	assert.Equal(t, "flows.app.menu", Setting.Menu.Process)
	assert.Equal(t, true, Setting.Optional.HideNotification)
	assert.Equal(t, false, Setting.Optional.HideSetting)
}

func TestLoadCN(t *testing.T) {

	err := i18n.Load(config.Conf)
	if err != nil {
		t.Fatal(err)
	}

	err = Load(config.Conf)
	if err != nil {
		t.Fatal(err)
	}

	newSetting, err := i18n.Trans("zh-cn", "app", "app", Setting)
	if err != nil {
		t.Fatal(err)
	}
	setting := newSetting.(*DSL)

	assert.Equal(t, "示例应用", setting.Name)
	assert.Equal(t, "演示", setting.Short)
	assert.Equal(t, "又一个 YAO 应用", setting.Description)
	assert.Equal(t, []string{"demo"}, setting.Menu.Args)
	assert.Equal(t, "flows.app.menu", setting.Menu.Process)
	assert.Equal(t, true, setting.Optional.HideNotification)
	assert.Equal(t, false, setting.Optional.HideSetting)

	assert.Equal(t, "::Demo Application", Setting.Name)
	assert.Equal(t, "::Demo", Setting.Short)
	assert.Equal(t, "::Another yao application", Setting.Description)
	assert.Equal(t, []string{"demo"}, Setting.Menu.Args)
	assert.Equal(t, "flows.app.menu", Setting.Menu.Process)
	assert.Equal(t, true, Setting.Optional.HideNotification)
	assert.Equal(t, false, Setting.Optional.HideSetting)
}

func TestExport(t *testing.T) {

	err := login.Load(config.Conf)
	if err != nil {
		t.Fatal(err)
	}

	err = Load(config.Conf)
	if err != nil {
		t.Fatal(err)
	}

	err = Export()
	if err != nil {
		t.Fatal(err)
	}

	api, has := gou.APIs["widgets.app"]
	assert.True(t, has)
	assert.Equal(t, 3, len(api.HTTP.Paths))

	_, has = gou.ThirdHandlers["yao.app.setting"]
	assert.True(t, has)

	_, has = gou.ThirdHandlers["yao.app.xgen"]
	assert.True(t, has)

	_, has = gou.ThirdHandlers["yao.app.menu"]
	assert.True(t, has)
}

func TestProcessSetting(t *testing.T) {
	loadApp(t)
	res, err := gou.NewProcess("yao.app.Setting").Exec()
	if err != nil {
		t.Fatal(err)
	}

	setting, ok := res.(DSL)
	assert.True(t, ok)
	assert.Equal(t, "Demo Application", setting.Name)
	assert.Equal(t, "Demo", setting.Short)
	assert.Equal(t, "Another yao application", setting.Description)
	assert.Equal(t, []string{"demo"}, setting.Menu.Args)
	assert.Equal(t, "flows.app.menu", setting.Menu.Process)
	assert.Equal(t, true, setting.Optional.HideNotification)
	assert.Equal(t, false, setting.Optional.HideSetting)
}

func TestProcessXgen(t *testing.T) {
	loadApp(t)
	res, err := gou.NewProcess("yao.app.Xgen").Exec()
	if err != nil {
		t.Fatal(err)
	}

	xgen := any.Of(res).MapStr().Dot()
	assert.Equal(t, "__yao", xgen.Get("apiPrefix"))
	assert.Equal(t, "Another yao application", xgen.Get("description"))
	assert.Equal(t, "/api/__yao/login/admin/captcha?type=digit", xgen.Get("login.admin.captcha"))
	assert.Equal(t, "/api/__yao/login/admin", xgen.Get("login.admin.login"))
	assert.Equal(t, "/x/Chart/dashboard", xgen.Get("login.entry.admin"))
	assert.Equal(t, "/x/Table/pet", xgen.Get("login.entry.user"))
	assert.Equal(t, "/api/__yao/login/user/captcha?type=digit", xgen.Get("login.user.captcha"))
	assert.Equal(t, "/api/__yao/login/user", xgen.Get("login.user.login"))
	assert.Equal(t, "/assets/images/login/cover.svg", xgen.Get("login.layout.cover"))
	assert.Equal(t, "/api/__yao/app/icons/app.ico", xgen.Get("favicon"))
	assert.Equal(t, "/api/__yao/app/icons/app.png", xgen.Get("logo"))
	assert.Equal(t, os.Getenv("YAO_ENV"), xgen.Get("mode"))
	assert.Equal(t, "Demo Application", xgen.Get("name"))
	assert.Equal(t, true, xgen.Get("optional.hideNotification"))
	assert.Equal(t, "localStorage", xgen.Get("token"))
}

func TestProcessMenu(t *testing.T) {
	loadApp(t)
	res, err := gou.NewProcess("yao.app.Menu").Exec()
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, res, 2)
}

func TestProcessIcons(t *testing.T) {
	loadApp(t)
	res, err := gou.NewProcess("yao.app.Icons", "app.png").Exec()
	if err != nil {
		t.Fatal(err)
	}
	assert.Greater(t, len(res.(string)), 10)
}

func loadApp(t *testing.T) {

	err := i18n.Load(config.Conf)
	if err != nil {
		t.Fatal(err)
	}
	lang.Pick("en-us").AsDefault()

	err = login.Load(config.Conf)
	if err != nil {
		t.Fatal(err)
	}

	err = flow.Load(config.Conf)
	if err != nil {
		t.Fatal(err)
	}

	err = Load(config.Conf)
	if err != nil {
		t.Fatal(err)
	}

	err = Export()
	if err != nil {
		t.Fatal(err)
	}

}
