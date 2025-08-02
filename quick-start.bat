@echo off
echo 🏥 Kansas Healthcare Analytics Platform - Quick Start
echo ================================================
echo.
echo This will automatically set up and run the healthcare demo.
echo Please wait while we prepare everything...
echo.

REM Check if PowerShell is available
powershell -Command "Write-Host 'PowerShell is available'" >nul 2>&1
if errorlevel 1 (
    echo ❌ PowerShell is not available. Please run setup-demo.ps1 manually.
    pause
    exit /b 1
)

REM Run the PowerShell setup script
echo 🚀 Starting automated setup...
powershell -ExecutionPolicy Bypass -File "%~dp0setup-demo.ps1"

pause