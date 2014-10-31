package routers

import (
	"github.com/astaxie/beego"
	"github.com/xulei8/haiyandb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/gocrm_login", &controllers.MainController{})
	beego.Router("/gocrm_help", &controllers.HelpController{})
	beego.Router("/gocrm_crmadmin", &controllers.CrmAdmin{})
	beego.Router("/gocrm_agent", &controllers.Agent{})
	beego.Router("/gocrm_agentcall", &controllers.AgentCall{})
	beego.Router("/gocrm_report", &controllers.MainController{})
	beego.Router("/gocrm_crmimport", &controllers.CrmImport{})
	beego.Router("/gocrm_agentcallaction", &controllers.CallAction{})
	beego.Router("/gocrm_agentcalllog", &controllers.CallLog{})

	beego.Router("/data_save/", &controllers.NewSQLSave{})
	beego.Router("/data_get/", &controllers.NewSQLGet{})
	beego.Router("/data_delete/", &controllers.NewSQLDelete{})
	beego.Router("/data_update/", &controllers.NewSQLUpdate{})
	beego.Router("/data_select_values/", &controllers.NewSQLSelectValue{})

}
