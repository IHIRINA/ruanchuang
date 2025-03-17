package server

import (
	"github.com/dingdinglz/ai-swindle-detecter-backend/ai"
	"github.com/dingdinglz/ai-swindle-detecter-backend/setting"
	"github.com/gofiber/fiber/v2"
)

func AIApiRoute(c *fiber.Ctx) error {
	if c.FormValue("text", "") == "" {
		return c.JSON(fiber.Map{"code": -1, "message": "参数不全"})
	}
	if setting.SettingVar.Debug {
		return c.JSON(fiber.Map{"code": 0, "message": "", "type": "中性"})
	}
	res := ai.Run(c.FormValue("text", ""), setting.SettingVar.AIPort)
	if res == "err" {
		return c.JSON(fiber.Map{"code": 1, "message": "ai error"})
	}
	resMap := make(map[string]int)
	resMap["中性"] = 0
	resMap["网络交易及兼职诈骗"] = 1
	resMap["虚假金融及投资诈骗"] = 2
	resMap["身份冒充及威胁诈骗"] = 3
	return c.JSON(fiber.Map{"code": 0, "message": "", "type_id": resMap[res], "type": res})
}
