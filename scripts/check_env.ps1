# Check for Go
Write-Host "Checking for Go..."
if (Get-Command go -ErrorAction SilentlyContinue) {
    go version
} else {
    Write-Warning "Go is not found in PATH."
}

# Check for Flutter
Write-Host "Checking for Flutter..."
if (Get-Command flutter -ErrorAction SilentlyContinue) {
    flutter --version
} else {
    Write-Warning "Flutter is not found in PATH."
    Write-Host "Searching for flutter.bat in C:\Users..."
    $flutterPath = Get-ChildItem -Path C:\Users -Filter flutter.bat -Recurse -ErrorAction SilentlyContinue -Depth 4
    if ($flutterPath) {
        Write-Host "Found Flutter at: $($flutterPath.DirectoryName)"
        Write-Host "Please add $($flutterPath.DirectoryName) to your System Environment Path."
    } else {
        Write-Error "Flutter SDK not found. Please install it from flutter.dev."
    }
}
