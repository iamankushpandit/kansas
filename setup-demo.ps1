Write-Host "🏥 Kansas Healthcare Analytics Platform - Demo Setup" -ForegroundColor Green
Write-Host "=================================================" -ForegroundColor Green
Write-Host ""

# Check Docker
Write-Host "🔍 Checking Docker..." -ForegroundColor Cyan
try {
    docker --version | Out-Null
    Write-Host "✅ Docker is available" -ForegroundColor Green
} catch {
    Write-Host "❌ Docker not found. Please install Docker Desktop first." -ForegroundColor Red
    Read-Host "Press Enter to exit"
    exit 1
}

# Check if Docker is running
Write-Host "🔍 Checking if Docker is running..." -ForegroundColor Cyan
try {
    docker ps | Out-Null
    Write-Host "✅ Docker is running" -ForegroundColor Green
} catch {
    Write-Host "❌ Docker is not running. Please start Docker Desktop." -ForegroundColor Red
    Read-Host "Press Enter to exit"
    exit 1
}

# Navigate to script directory
$scriptPath = Split-Path -Parent $MyInvocation.MyCommand.Path
Set-Location $scriptPath

Write-Host ""
Write-Host "🏗️ Starting the healthcare platform..." -ForegroundColor Green
Write-Host ""

# Stop existing containers
docker-compose down 2>$null

# Start the application
docker-compose up --build -d

Write-Host ""
Write-Host "🎉 SUCCESS! Application is starting..." -ForegroundColor Green
Write-Host ""
Write-Host "📱 Frontend: http://localhost:4192" -ForegroundColor Cyan
Write-Host "🔧 Backend: http://localhost:3247" -ForegroundColor Cyan
Write-Host ""
Write-Host "⏳ Waiting 15 seconds then opening browser..." -ForegroundColor Yellow

Start-Sleep -Seconds 15
Start-Process "http://localhost:4192"

Write-Host ""
Write-Host "✅ Demo is ready! Browser should open automatically." -ForegroundColor Green
Write-Host "🛑 To stop: run 'docker-compose down'" -ForegroundColor Yellow

Read-Host "Press Enter to exit"