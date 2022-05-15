# 定义数组
$oss=@(
    "windows"
)
$archs=@(
    "amd64",
    "arm64"
)
$date=Get-Date
$buildDate=$date

$v=(go version)
foreach($os in $oss){
    foreach($arch in $archs) {
        Write-Host $os-$arch
        SET GOARCH=$arch
        SET GOOS=$os
        $env:GOARCH=$arch
        $env:GOOS=$os
        (go build -ldflags "-X 'main.BuildDate=$buildDate' "  -o ./bin/meituan_$os-$arch .)
    }
}
