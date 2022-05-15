SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
@REM # android/amd64
@REM darwin/arm64
@REM darwin/amd64
go build  -o ./bin/meituan_windows_amd64.exe .