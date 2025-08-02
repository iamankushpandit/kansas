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

# Check Docker
Write-Host "Checking Docker..." -ForegroundColor Cyan
try {
    docker --version | Out-Null
    Write-Host "[OK] Docker is available" -ForegroundColor Green
} catch {
    Write-Host "[ERROR] Docker not found. Please install Docker Desktop first." -ForegroundColor Red
    Read-Host "Press Enter to exit"
    exit 1
}

# Check if Docker is running
Write-Host "Checking if Docker is running..." -ForegroundColor Cyan
try {
    docker ps | Out-Null
    Write-Host "[OK] Docker is running" -ForegroundColor Green
} catch {
    Write-Host "[ERROR] Docker is not running. Please start Docker Desktop." -ForegroundColor Red
    Read-Host "Press Enter to exit"
    exit 1
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

# Extract actual coverage percentage
$coverageMatch = $frontendOutput | Select-String "All files.*?(\d+\.\d+)"
if ($coverageMatch) {
    $actualCoverage = [double]$coverageMatch.Matches[0].Groups[1].Value
    Write-Host "[OK] Frontend tests passed - coverage: $actualCoverage%" -ForegroundColor Green
} else {
    Write-Host "[OK] Frontend tests passed" -ForegroundColor Green
    $actualCoverage = 0
}
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

# Extract backend coverage (simplified check)
$backendCoverageMatch = $backendOutput | Select-String "coverage: (\d+\.\d+)%"
if ($backendCoverageMatch) {
    $backendCoverage = [double]$backendCoverageMatch.Matches[0].Groups[1].Value
    Write-Host "[OK] Backend tests passed - coverage: $backendCoverage%" -ForegroundColor Green
} else {
    Write-Host "[OK] Backend tests passed" -ForegroundColor Green
    $backendCoverage = 0
}
Set-Location ".."

Write-Host ""
Write-Host "Test Coverage Summary:" -ForegroundColor Yellow
if ($actualCoverage -gt 0) {
    Write-Host "[COVERAGE] Frontend: $actualCoverage% code coverage" -ForegroundColor Green
} else {
    Write-Host "[COVERAGE] Frontend: Tests passed" -ForegroundColor Green
}
if ($backendCoverage -gt 0) {
    Write-Host "[COVERAGE] Backend: $backendCoverage% code coverage" -ForegroundColor Green
} else {
    Write-Host "[COVERAGE] Backend: Tests passed" -ForegroundColor Green
}
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