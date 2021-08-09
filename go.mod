module m3u8-downloader-extension-native-message

go 1.15

require (
	github.com/mitchellh/go-ps v1.0.0
	github.com/mitchellh/mapstructure v1.4.1
	github.com/satom9to5/m3u8-download/m3u8 v0.0.0-20210527144927-cec838bb1dcb
	github.com/satom9to5/webext v0.0.0-20190226145152-1cd0537acbbc
)

replace (
	github.com/satom9to5/m3u8-download/m3u8 v0.0.0-20210527144927-cec838bb1dcb => ../m3u8-download/m3u8
	github.com/satom9to5/webext v0.0.0-20190226145152-1cd0537acbbc => ../webext
)
