// Package xingzuo
package xingzuo

import (
	"github.com/FloatTech/floatbox/web"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	engine := control.Register("xingzuo", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief: "星座",
		Help: "今日[白羊座|金牛座|双子座|巨蟹座|狮子座|处女座|天秤座|天蝎座|射手座|摩羯座|水瓶座|双鱼座]",
		OnEnable: func(ctx *zero.Ctx) {
			ctx.Send("插件已启用")
		},
		OnDisable: func(ctx *zero.Ctx) {
			ctx.Send("插件已禁用")
		},
	})
	engine.OnRegex(`^今日(.{1,25})$`).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			xz := ctx.State["regex_matched"].([]string)[1]
			data, err := web.GetData("https://xiaoapi.cn/API/xzys_pic.php?msg=" + xz)
			if err != nil {
				ctx.SendChain(message.Text("ERROR: ", err))
				return
			}
			ctx.SendChain(message.ImageBytes(data))
		})
}
