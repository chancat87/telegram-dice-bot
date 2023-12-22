package enums

// BotPrivateChatStatus 代表枚举的自定义类型
type BotPrivateChatStatus struct {
	Value string
	Name  string
}

// 枚举映射
var BotPrivateChatStatusMap = make(map[string]BotPrivateChatStatus)

// 构造函数
func newBotPrivateChatStatus(value string, name string) BotPrivateChatStatus {
	enum := BotPrivateChatStatus{Value: value, Name: name}
	BotPrivateChatStatusMap[value] = enum
	return enum
}

// 使用构造函数定义枚举值等
var (
	WaitGameDrawCycle = newBotPrivateChatStatus("WAIT_GAME_DRAW_CYCLE", "开奖周期设置")
)

// GetBotPrivateChatStatus 通过 value 获取枚举项
func GetBotPrivateChatStatus(value string) (BotPrivateChatStatus, bool) {
	enum, ok := BotPrivateChatStatusMap[value]
	return enum, ok

}