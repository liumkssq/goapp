# scripts/run_all.ps1
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8
Write-Host "启动所有服务..." -ForegroundColor Yellow

$ErrorActionPreference = 'SilentlyContinue'
$services = @(
    'app/user/rpc',
    'app/social/rpc',
    'app/im/rpc',
    'app/im/im',
    'app/product/rpc',
    'app/article/rpc',
    'app/search/rpc',
    'app/lostfound/rpc',
    'app/task/rpc'
)

foreach ($dir in $services) {
    # 判断该目录下是否有 .go 文件
    if (Test-Path "$dir/*.go") {
        Write-Host "正在启动服务目录：$dir" -ForegroundColor Cyan

        # 遍历该目录下所有的 .go 文件，获取其文件名并拼接成字符串
        $goFiles = Get-ChildItem -Path $dir -Filter *.go | ForEach-Object { $_.Name }
        $filesStr = $goFiles -join " "

        # 构造命令，这里用单引号包裹目录，确保目录带有空格时也能正确识别
        $cmd = "cd '$dir'; go run $filesStr"
        Write-Host "执行命令：$cmd" -ForegroundColor Green

        # 执行命令：这里使用 Start-Process 并指定 -NoNewWindow，如果你希望新窗口显示输出，可去掉此参数
        Start-Process powershell -ArgumentList "-NoProfile", "-Command", $cmd -NoNewWindow
    }
    else {
        Write-Host "未找到目录或 .go 文件：$dir" -ForegroundColor Yellow
    }
}

Write-Host "所有服务已启动" -ForegroundColor Magenta
