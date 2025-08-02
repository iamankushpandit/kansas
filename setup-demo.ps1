# Kansas Healthcare Analytics Platform - Automated Demo Setup
# This script installs Docker Desktop and runs the healthcare analytics application
# Designed for non-technical users at hackathons

Write-Host "🏥 Kansas Healthcare Analytics Platform - Demo Setup" -ForegroundColor Green
Write-Host "=================================================" -ForegroundColor Green
Write-Host ""

# Check if running as Administrator
if (-NOT ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")) {
    Write-Host "❌ This script needs Administrator privileges to install Docker Desktop." -ForegroundColor Red
    Write-Host "Right-click on PowerShell and select 'Run as Administrator', then run this script again." -ForegroundColor Yellow
    Read-Host "Press Enter to exit"
    exit 1
}

# Function to check if Docker Desktop is installed
function Test-DockerInstalled {
    try {
        $dockerPath = Get-Command docker -ErrorAction SilentlyContinue
        return $dockerPath -ne $null
    }
    catch {
        return $false
    }
}

# Function to check if Docker Desktop is running
function Test-DockerRunning {
    try {
        docker version | Out-Null
        return $true
    }
    catch {
        return $false
    }
}

# Function to download and install Docker Desktop
function Install-DockerDesktop {
    Write-Host "📥 Downloading Docker Desktop..." -ForegroundColor Yellow
    
    $dockerUrl = "https://desktop.docker.com/win/main/amd64/Docker%20Desktop%20Installer.exe"
    $dockerInstaller = "$env:TEMP\DockerDesktopInstaller.exe"
    
    try {
        Invoke-WebRequest -Uri $dockerUrl -OutFile $dockerInstaller -UseBasicParsing
        Write-Host "✅ Docker Desktop downloaded successfully" -ForegroundColor Green
        
        Write-Host "🔧 Installing Docker Desktop..." -ForegroundColor Yellow
        Write-Host "   This may take 5-10 minutes. Please wait..." -ForegroundColor Cyan
        
        Start-Process -FilePath $dockerInstaller -ArgumentList "install", "--quiet" -Wait
        
        Write-Host "✅ Docker Desktop installed successfully" -ForegroundColor Green
        Write-Host "🔄 Please restart your computer and run this script again." -ForegroundColor Yellow
        Read-Host "Press Enter to exit"
        exit 0
    }
    catch {
        Write-Host "❌ Failed to download or install Docker Desktop: $($_.Exception.Message)" -ForegroundColor Red
        Write-Host "Please download Docker Desktop manually from: https://www.docker.com/products/docker-desktop" -ForegroundColor Yellow
        Read-Host "Press Enter to exit"
        exit 1
    }
}

# Function to wait for Docker Desktop to start
function Wait-ForDocker {
    Write-Host "⏳ Waiting for Docker Desktop to start..." -ForegroundColor Yellow
    $timeout = 120 # 2 minutes timeout
    $elapsed = 0
    
    while (-not (Test-DockerRunning) -and $elapsed -lt $timeout) {
        Start-Sleep -Seconds 5
        $elapsed += 5
        Write-Host "   Still waiting... ($elapsed/$timeout seconds)" -ForegroundColor Cyan
    }
    
    if (Test-DockerRunning) {
        Write-Host "✅ Docker Desktop is running!" -ForegroundColor Green
        return $true
    } else {
        Write-Host "❌ Docker Desktop failed to start within $timeout seconds" -ForegroundColor Red
        return $false
    }
}

# Main installation process
Write-Host "🔍 Checking Docker Desktop installation..." -ForegroundColor Cyan

if (-not (Test-DockerInstalled)) {
    Write-Host "❌ Docker Desktop not found. Installing..." -ForegroundColor Yellow
    Install-DockerDesktop
} else {
    Write-Host "✅ Docker Desktop is installed" -ForegroundColor Green
}

# Check if Docker is running
if (-not (Test-DockerRunning)) {
    Write-Host "🚀 Starting Docker Desktop..." -ForegroundColor Yellow
    
    # Try to start Docker Desktop
    try {
        Start-Process "C:\Program Files\Docker\Docker\Docker Desktop.exe" -WindowStyle Hidden
    }
    catch {
        Write-Host "❌ Could not start Docker Desktop automatically." -ForegroundColor Red
        Write-Host "Please start Docker Desktop manually and run this script again." -ForegroundColor Yellow
        Read-Host "Press Enter to exit"
        exit 1
    }
    
    if (-not (Wait-ForDocker)) {
        Write-Host "❌ Docker Desktop is not responding. Please:" -ForegroundColor Red
        Write-Host "   1. Open Docker Desktop manually" -ForegroundColor Yellow
        Write-Host "   2. Wait for it to fully start" -ForegroundColor Yellow
        Write-Host "   3. Run this script again" -ForegroundColor Yellow
        Read-Host "Press Enter to exit"
        exit 1
    }
} else {
    Write-Host "✅ Docker Desktop is already running" -ForegroundColor Green
}

# Navigate to project directory
$scriptPath = Split-Path -Parent $MyInvocation.MyCommand.Path
Set-Location $scriptPath

Write-Host ""
Write-Host "🏗️ Building and starting the healthcare analytics platform..." -ForegroundColor Green
Write-Host "   This will take 3-5 minutes for the first run..." -ForegroundColor Cyan
Write-Host ""

# Stop any existing containers
try {
    docker-compose down 2>$null
}
catch {
    # Ignore errors if no containers are running
}

# Build and start the application
try {
    docker-compose up --build -d
    
    Write-Host ""
    Write-Host "🎉 SUCCESS! The Kansas Healthcare Analytics Platform is starting..." -ForegroundColor Green
    Write-Host ""
    Write-Host "📱 Access the application at: http://localhost:4192" -ForegroundColor Cyan
    Write-Host "🔧 Backend API available at: http://localhost:3247" -ForegroundColor Cyan
    Write-Host ""
    Write-Host "⏳ Please wait 30-60 seconds for the application to fully load..." -ForegroundColor Yellow
    Write-Host ""
    
    # Wait a moment then try to open the browser
    Start-Sleep -Seconds 10
    
    Write-Host "🌐 Opening application in your default browser..." -ForegroundColor Green
    Start-Process "http://localhost:4192"
    
    Write-Host ""
    Write-Host "📋 Demo Instructions:" -ForegroundColor Cyan
    Write-Host "   • Use the dropdown filters to explore different provider networks" -ForegroundColor White
    Write-Host "   • Click on any county to see detailed analytics" -ForegroundColor White
    Write-Host "   • Export PDF reports using the Export button" -ForegroundColor White
    Write-Host "   • Toggle between light/dark themes" -ForegroundColor White
    Write-Host ""
    Write-Host "🛑 To stop the demo: Press Ctrl+C or close this window" -ForegroundColor Yellow
    Write-Host ""
    
    # Keep the script running and show logs
    Write-Host "📊 Application logs (press Ctrl+C to stop):" -ForegroundColor Green
    docker-compose logs -f
    
}
catch {
    Write-Host "❌ Failed to start the application: $($_.Exception.Message)" -ForegroundColor Red
    Write-Host ""
    Write-Host "🔧 Troubleshooting steps:" -ForegroundColor Yellow
    Write-Host "   1. Make sure Docker Desktop is running" -ForegroundColor White
    Write-Host "   2. Close any applications using ports 3247 or 4192" -ForegroundColor White
    Write-Host "   3. Run this script again" -ForegroundColor White
    Read-Host "Press Enter to exit"
    exit 1
}