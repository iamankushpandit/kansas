Write-Host "Kansas Healthcare Analytics Platform - Demo Setup" -ForegroundColor Green
Write-Host "=================================================" -ForegroundColor Green
Write-Host ""

# Check for test parameter
if ($args -contains "--test" -or $args -contains "-t") {
    Write-Host "[INFO] Running frontend tests with coverage..." -ForegroundColor Cyan
    Set-Location "kansas-healthcare-map"
    
    if (-not (Test-Path "node_modules")) {
        Write-Host "[INFO] Installing test dependencies..." -ForegroundColor Cyan
        npm install
    }
    
    npm run test:coverage
    Write-Host "[SUCCESS] Tests completed! Coverage report in coverage/index.html" -ForegroundColor Green
    Read-Host "Press Enter to exit"
    exit 0
}

# Check and install Docker Desktop if needed
Write-Host "Checking Docker Desktop..." -ForegroundColor Cyan
try {
    docker --version | Out-Null
    Write-Host "[OK] Docker is available" -ForegroundColor Green
} catch {
    Write-Host "[INFO] Docker Desktop not found. Installing automatically..." -ForegroundColor Yellow
    
    # Download Docker Desktop installer
    $dockerUrl = "https://desktop.docker.com/win/main/amd64/Docker%20Desktop%20Installer.exe"
    $dockerInstaller = "$env:TEMP\DockerDesktopInstaller.exe"
    
    Write-Host "[INFO] Downloading Docker Desktop (this may take a few minutes)..." -ForegroundColor Cyan
    try {
        Invoke-WebRequest -Uri $dockerUrl -OutFile $dockerInstaller -UseBasicParsing
        Write-Host "[INFO] Installing Docker Desktop silently..." -ForegroundColor Cyan
        Start-Process -FilePath $dockerInstaller -ArgumentList "install", "--quiet", "--accept-license" -Wait
        Remove-Item $dockerInstaller -Force
        
        Write-Host "[SUCCESS] Docker Desktop installed successfully!" -ForegroundColor Green
        Write-Host "[INFO] Please restart your computer and run this script again." -ForegroundColor Yellow
        Read-Host "Press Enter to exit"
        exit 0
    } catch {
        Write-Host "[ERROR] Failed to install Docker Desktop automatically." -ForegroundColor Red
        Write-Host "[INFO] Please download and install Docker Desktop manually from: https://www.docker.com/products/docker-desktop" -ForegroundColor Yellow
        Read-Host "Press Enter to exit"
        exit 1
    }
}

# Check if Docker is running
Write-Host "Checking if Docker is running..." -ForegroundColor Cyan
try {
    docker ps | Out-Null
    Write-Host "[OK] Docker is running" -ForegroundColor Green
} catch {
    Write-Host "[INFO] Docker Desktop is not running. Starting it..." -ForegroundColor Yellow
    
    # Try to start Docker Desktop
    $dockerPath = "$env:ProgramFiles\Docker\Docker\Docker Desktop.exe"
    if (Test-Path $dockerPath) {
        Start-Process -FilePath $dockerPath
        Write-Host "[INFO] Waiting for Docker Desktop to start (30 seconds)..." -ForegroundColor Cyan
        
        # Wait up to 60 seconds for Docker to start
        $timeout = 60
        $elapsed = 0
        while ($elapsed -lt $timeout) {
            Start-Sleep -Seconds 5
            $elapsed += 5
            try {
                docker ps | Out-Null
                Write-Host "[OK] Docker is now running" -ForegroundColor Green
                break
            } catch {
                Write-Host "[INFO] Still waiting for Docker... ($elapsed/$timeout seconds)" -ForegroundColor Cyan
            }
        }
        
        # Final check
        try {
            docker ps | Out-Null
        } catch {
            Write-Host "[ERROR] Docker failed to start. Please start Docker Desktop manually." -ForegroundColor Red
            Read-Host "Press Enter to exit"
            exit 1
        }
    } else {
        Write-Host "[ERROR] Docker Desktop not found. Please install it manually." -ForegroundColor Red
        Read-Host "Press Enter to exit"
        exit 1
    }
}

# Navigate to script directory
$scriptPath = Split-Path -Parent $MyInvocation.MyCommand.Path
Set-Location $scriptPath

Write-Host ""
Write-Host "Running tests before build..." -ForegroundColor Yellow
Write-Host ""

# Run frontend tests with coverage
Write-Host "[TEST] Running frontend unit tests with coverage..." -ForegroundColor Cyan
Set-Location "kansas-healthcare-map"
if (-not (Test-Path "node_modules")) {
    npm install --silent
}
$frontendOutput = npm run test:run -- --coverage --reporter=verbose 2>&1
Write-Host $frontendOutput
if ($LASTEXITCODE -ne 0) {
    Write-Host "[ERROR] Frontend tests failed!" -ForegroundColor Red
    Read-Host "Press Enter to exit"
    exit 1
}

Write-Host "[OK] Frontend tests passed" -ForegroundColor Green
Set-Location ".."

# Run backend tests
Write-Host "[TEST] Running backend unit tests..." -ForegroundColor Cyan
Set-Location "kansas-healthcare-backend"
$backendOutput = go test ./... -cover -v 2>&1
Write-Host $backendOutput
if ($LASTEXITCODE -ne 0) {
    Write-Host "[ERROR] Backend tests failed!" -ForegroundColor Red
    Read-Host "Press Enter to exit"
    exit 1
}

Write-Host "[OK] Backend tests passed" -ForegroundColor Green
Set-Location ".."

Write-Host ""
Write-Host "Test Summary:" -ForegroundColor Yellow
Write-Host "[TESTS] Frontend: Passed" -ForegroundColor Green
Write-Host "[TESTS] Backend: Passed" -ForegroundColor Green
Write-Host ""
Write-Host "Starting the healthcare platform..." -ForegroundColor Green
Write-Host ""

# Stop existing containers
docker-compose down 2>$null

# Check for rebuild parameter
if ($args -contains "--build" -or $args -contains "-b") {
    Write-Host "[INFO] Force rebuilding images" -ForegroundColor Cyan
    docker-compose up --build -d
} else {
    # Check if images exist
    $backendExists = docker images kansas-backend -q
    $frontendExists = docker images kansas-frontend -q
    
    if ($backendExists -and $frontendExists) {
        Write-Host "[INFO] Using existing images (add --build to force rebuild)" -ForegroundColor Cyan
        docker-compose up --build -d
    } else {
        Write-Host "[INFO] Building images (first run)" -ForegroundColor Cyan
        docker-compose up --build -d
    }
}

Write-Host ""
Write-Host "[SUCCESS] Application is starting..." -ForegroundColor Green
Write-Host ""
Write-Host "Frontend: http://localhost:4192" -ForegroundColor Cyan
Write-Host "Backend: http://localhost:3247" -ForegroundColor Cyan
Write-Host ""
Write-Host "Waiting 7 seconds then opening browser..." -ForegroundColor Yellow

Start-Sleep -Seconds 7
Start-Process "http://localhost:4192"

Write-Host ""
Write-Host "[READY] Demo is ready! Browser should open automatically." -ForegroundColor Green
Write-Host "[INFO] To stop: run 'docker-compose down'" -ForegroundColor Yellow
Write-Host "[INFO] To run tests: .\setup-demo.ps1 --test" -ForegroundColor Cyan

Read-Host "Press Enter to exit"