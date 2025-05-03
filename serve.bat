@echo off
echo Starting local web server for your homepage...
cd gravitational-particles.go-main
python -m http.server
echo Server stopped.
pause 