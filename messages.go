package spice

import "fmt"

type RedLinkMess struct {
	Magick         uint32
	VersionMajor   uint32
	VersionMinor   uint32
	Size           uint32
	ConnectionID   uint32
	ChannelType    ChannelType
	ChannelID      uint8
	CommonCAPSNum  uint32
	ChannelCAPSNum uint32
	CAPSOffset     uint32
}

func (msg *RedLinkMess) String() string {
	return fmt.Sprintf("%v", msg)
}

type RedLinkReply struct {
	Magick         uint32
	VersionMajor   uint32
	VersionMinor   uint32
	Size           uint32
	Error          Error
	Pubkey         [RedTicketPubkeyBytes]byte
	CommonCAPSNum  uint32
	ChannelCAPSNum uint32
	CAPSOffset     uint32
}

func (msg *RedLinkReply) String() string {
	return fmt.Sprintf("Magic:0x%x, Major: %d, Minor: %d, Size: %d, Error: %s Pubkey: %v, CommonCAPSNum: %d, ChannelCAPSNum: %d, CAPSOffset: %d", msg.Magick, msg.VersionMajor, msg.VersionMinor, msg.Size, msg.Error, msg.Pubkey, msg.CommonCAPSNum, msg.ChannelCAPSNum, msg.CAPSOffset)
}

type LinkResult uint32

type RedDataHeader struct {
	Serial  uint64
	Type    MessageType
	Size    uint32
	SubList uint32
}

type RedSubMessageList struct {
	Size        uint16
	SubMessages []uint32
}

type RedSubMessage struct {
	Type uint16
	Size uint32
}

type MessageType uint16

const (
	_ MessageType = iota

	RedMigrateMsgType
	RedMigrateDataMsgType
	RedSetAckMsgType
	RedPingMsgType
	RedWaitForChannelsMsgType
	RedDisconnectingMsgType
	RedNotifyMsgType

	RedFirstAvailMessageMsgType = 101
)

const (
	_ MessageType = iota

	RedcAckSyncMsgType
	RedcAckMsgType
	RedcPongMsgType
	RedcMigrateFlushMarkMsgType
	RedcMigrateDataMsgType
	RedcDisconnectingMsgType

	RedcFirstAvailMessagemsgType = 101
)
