package main

import (
	"Gogogo/4_common/json/kafkaDef"
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Json Test Start...")

	//UnmarshalTest()
	UnmarshalSliceTest()
}

func UnmarshalSliceTest() {
	contract_1 := kafkaDef.Contract{MarketType:8100, Symbol:"EURUSD"}
	contract_2 := kafkaDef.Contract{MarketType:8200, Symbol:"USDCHN"}
	contract_3 := kafkaDef.Contract{MarketType:8200, Symbol:"USDCAD"}
	contract_4 := kafkaDef.Contract{MarketType:8500, Symbol:"XPTUSD"}
	contract_5 := kafkaDef.Contract{MarketType:8500, Symbol:"XAUUSD"}

	contracts := kafkaDef.OptionalList{Contracts:[]kafkaDef.Contract{contract_1, contract_2, contract_3, contract_4, contract_5}}
	b, err := json.Marshal(contracts)
	if err != nil {
		fmt.Printf("[Marshal Error]: [%s]", err)
		return
	} else {
		fmt.Printf("%s\n", b)
	}

	temp := kafkaDef.OptionalList{}
	err = json.Unmarshal(b, &temp)
	if err != nil {
		fmt.Printf("[UnMarshal Error]: [%s]", err)
		return
	} else {
		fmt.Printf("%v\n", temp)
	}
}

func UnmarshalTest() {
	b := NewOptionalInfoPush()

	mh := kafkaDef.MH{}
	err := json.Unmarshal(b, &mh)
	if err != nil {
		fmt.Printf("[Unmarshal Error]: [%s]\n", err)
	}

	fmt.Printf("[%d] [%s] [%d]\n", mh.Head.CMD, mh.Head.UserID, mh.Head.AccountID)
	if mh.Head.CMD == kafkaDef.CMD_ChartSyncNotify {
		data := kafkaDef.ChartSyncNotify{}
		err = json.Unmarshal(b, &data)
		if err != nil {
			fmt.Printf("[Unmarshal_1 Error]: [%s]\n", err)
		} else {
			fmt.Printf("[CMD: %d] [UserID: %s] [AccID: %d] [UUID: %s]\n", data.Head.CMD, data.Head.UserID, data.Head.AccountID, data.Uuid)
		}

	} else if mh.Head.CMD == kafkaDef.CMD_TradeNotify {
		data := kafkaDef.TradeNotifySyncNotify{}
		err = json.Unmarshal(b, &data)
		if err != nil {
			fmt.Printf("[Unmarshal_2 Error]: [%s]\n", err)
		} else {
			fmt.Printf("[CMD: %d] [UserID: %s] [AccID: %d] [NotifyNo: %d]\n", data.Head.CMD, data.Head.UserID, data.Head.AccountID, data.NotifyNo)
		}

	} else if mh.Head.CMD == kafkaDef.CMD_OptionalInfoPush {
		data := kafkaDef.OptionalInfoPush{}
		err = json.Unmarshal(b, &data)
		if err != nil {
			fmt.Printf("[Unmarshal_2 Error]: [%s]\n", err)
		} else {
			fmt.Printf("[CMD: %d] [UserID: %s] [AccID: %d]\n", data.Head.CMD, data.Head.UserID, data.Head.AccountID)

			for i := range data.Infos {
				fmt.Printf("[%d] [%s] [%d]\n", data.Infos[i].Contracts.MarketType, data.Infos[i].Contracts.Symbol, data.Infos[i].Stat)
			}
		}
	}
}

func NewMsgHead(cmd kafkaDef.KafkaDefCMD) kafkaDef.MsgHead {
	return kafkaDef.MsgHead{
		CMD:       uint32(cmd),
		UserID:    "wonima",
		AccountID: 6324,
	}
}

func NewChartSyncNotify() []byte {
	msgHead := NewMsgHead(kafkaDef.CMD_ChartSyncNotify)
	data := kafkaDef.ChartSyncNotify{
		Head: msgHead,
		Uuid: "233333",
	}

	d, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("[ChartSyncNotify Marshal Error]: [%s]\n", err)
		return nil
	}
	return d
}

func NewTradeNotifySyncNotify() []byte {
	msgHead := NewMsgHead(kafkaDef.CMD_TradeNotify)
	data := kafkaDef.TradeNotifySyncNotify{
		Head:     msgHead,
		NotifyNo: 233,
	}

	d, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("[ChartSyncNotify Marshal Error]: [%s]\n", err)
		return nil
	}
	return d
}

func NewOptionalInfoPush() []byte {
	msgHead := NewMsgHead(kafkaDef.CMD_OptionalInfoPush)
	info1 := &kafkaDef.OptionalInfo{
		Contracts: kafkaDef.Contract{
			MarketType	: 8100,
			Symbol		: "EURUSD",
		},
		Stat:      0,
	}
	info2 := &kafkaDef.OptionalInfo{
		Contracts: kafkaDef.Contract{
			MarketType	: 8200,
			Symbol		: "USDCAD",
		},
		Stat:      1,
	}
	infos := []*kafkaDef.OptionalInfo{info1, info2}

	data := kafkaDef.OptionalInfoPush{
		Head:  msgHead,
		Infos: infos,
	}
	d, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("[OptionalInfoPush Marshal Error]: [%s]\n", err)
		return nil
	}
	return d
}