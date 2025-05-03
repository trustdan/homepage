Write-Host "Starting local web server for your homepage..." -ForegroundColor Green
Set-Location -Path ".\gravitational-particles.go-main"
python -m http.server
Write-Host "Server stopped." -ForegroundColor Yellow
Read-Host -Prompt "Press Enter to exit" 