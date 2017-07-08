package spice

const (
	RedMagick = 1363428690 // "REDQ" in uint32 LittleEndian
)

const (
	RedVersionMajor uint32 = 0x2
	RedVersionMinor uint32 = 0x2
)

type Compatibily uint32

//go:generate stringer -type=ChannelType

type ChannelType uint8

const (
	_                          = iota
	RedChannelMain ChannelType = iota
	RedChannelDisplay
	RedChannelInputs
	RedChannelCursor
	RedChannelPlayback
	RedChannelRecord
)

//go:generate stringer -type=Error

type Error uint32

const (
	RedErrorOK Error = iota
	RedErrorError
	RedErrorInvalidMagick
	RedErrorInvalidData
	RedErrorVersionMismatch
	RedErrorNeedSecured
	RedErrorNeedUnsecured
	RedErrorPermissionDenied
	RedErrorBadConnectionID
	RedErrorChannelNotAvailable
)

//go:generate stringer -type=Warn

type Warn uint32

const (
	RedWarnGeneral Warn = iota
)

//go:generate stringer -type=Info

type Info uint32

const (
	RedInfoGeneral Info = iota
)

const (
	RedTicketPubkeyBytes = 162
)
