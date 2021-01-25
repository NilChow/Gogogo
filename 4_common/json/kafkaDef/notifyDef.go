package kafkaDef

type KafkaDefCMD int32

const (
	CMD_None				= 0
	CMD_ChartSyncNotify		= 1
	CMD_TradeNotify			= 2
	CMD_OptionalInfoPush	= 3
	CMD_MsgCenterPush		= 4
	CMD_BlockChangePush		= 5
)

type MsgHead struct {
	CMD			uint32	`json:"CMD,omitempty"`
	UserID		string	`json:"UserID,omitempty"`
	AccountID	uint64	`json:"AccountID,omitempty"`
}

type MH struct {
	Head MsgHead `json:"Head,omitempty"`
}

type Contract struct {
	MarketType	uint64		`json:"MarketType,omitempty"`
	Symbol		string		`json:"Symbol,omitempty"`
}

type OptionalList struct {
	Contracts	[]Contract	`json:"Contracts,omitempty"`
}

////////////////////////// [云同步] //////////////////////////

// 图表同步通知
type ChartSyncNotify struct {
	Head	MsgHead	`json:"Head,omitempty"`
	Uuid	string	`json:"Uuid,omitempty"`
}

// 提醒同步通知
type TradeNotifySyncNotify struct {
	Head		MsgHead	`json:"Head,omitempty"`
	NotifyNo	uint32	`json:"NotifyNo,omitempty"`
}

type OptionalInfo struct {
	Contracts	Contract
	Stat		uint32
}

// 自选列表变更推送
type OptionalInfoPush struct {
	Head	MsgHead	`json:"Head,omitempty"`
	Infos	[]*OptionalInfo
}

////////////////////////// [云同步] //////////////////////////



///////////////////////// [消息中心] /////////////////////////

// 屏蔽更新推送
type BlockChangePush struct {
	Head		MsgHead	`json:"Head,omitempty"`
	ContactID	string
	Stat		uint32
}

///////////////////////// [消息中心] /////////////////////////