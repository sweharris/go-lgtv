package control

const (
	uriVolumeUp   = "ssap://audio/volumeUp"
	uriVolumeDown = "ssap://audio/volumeDown"
	uriSetVolume  = "ssap://audio/setVolume"
	uriGetVolume  = "ssap://audio/getVolume"
	uriSetMute    = "ssap://audio/setMute"
	uriGetMute    = "ssap://audio/getMute"

	uriGetSoundOutput = "ssap://com.webos.service.apiadapter/audio/getSoundOutput"
	uriSetSoundOutput = "ssap://audio/changeSoundOutput"

	uriPlay        = "ssap://media.controls/play"
	uriPause       = "ssap://media.controls/pause"
	uriStop        = "ssap://media.controls/stop"
	uriRewind      = "ssap://media.controls/rewind"
	uriFastForward = "ssap://media.controls/fastForward"

	uriChannelUp             = "ssap://tv/channelUp"
	uriChannelDown           = "ssap://tv/channelDown"
	uriSetChannel            = "ssap://tv/openChannel"
	uriSwitchInput           = "ssap://tv/switchInput"
	uriGetExternalInputList  = "ssap://tv/getExternalInputList"
	uriGetChannelList        = "ssap://tv/getChannelList"
	uriGetCurrentChannel     = "ssap://tv/getCurrentChannel"
	uriGetChannelProgramInfo = "ssap://tv/getChannelProgramInfo"

	uriWriteToast = "ssap://system.notifications/createToast"

	uriListApps = "ssap://com.webos.applicationManager/listApps"

	uriLaunchApp = "ssap://system.launcher/launch"

	uriTurnOff = "ssap://system/turnOff"

	uriPowerState = "ssap://com.webos.service.tvpower/power/getPowerState"
)
