package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:EdgesController"] = append(beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:EdgesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:EdgesController"] = append(beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:EdgesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:EdgesController"] = append(beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:EdgesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:EdgesController"] = append(beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:EdgesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:EdgesController"] = append(beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:EdgesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:UsersController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:UsersController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:UsersController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:UsersController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/quchunguang/snsrobot-go/snsrobotdapi/controllers:UsersController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
