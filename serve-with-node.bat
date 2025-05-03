@echo off
echo Installing http-server if not already installed...
npm list -g http-server >nul 2>&1 || npm install -g http-server

echo Starting local web server for your homepage...
cd gravitational-particles.go-main
http-server -o demo.html
echo Server stopped.
pause 