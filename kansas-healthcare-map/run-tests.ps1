Write-Host "Kansas Healthcare Analytics - Frontend Test Suite" -ForegroundColor Green
Write-Host "=================================================" -ForegroundColor Green
Write-Host ""

# Check if we're in the right directory
if (-not (Test-Path "package.json")) {
    Write-Host "[ERROR] package.json not found. Please run from the frontend directory." -ForegroundColor Red
    Read-Host "Press Enter to exit"
    exit 1
}

# Check if node_modules exists
if (-not (Test-Path "node_modules")) {
    Write-Host "[INFO] Installing dependencies..." -ForegroundColor Cyan
    npm install
}

Write-Host "[INFO] Running unit tests with coverage..." -ForegroundColor Cyan
Write-Host ""

# Run tests with coverage
npm run test:coverage

Write-Host ""
Write-Host "[SUCCESS] Test execution completed!" -ForegroundColor Green
Write-Host "[INFO] Coverage report generated in coverage/ directory" -ForegroundColor Yellow
Write-Host "[INFO] Open coverage/index.html to view detailed coverage report" -ForegroundColor Yellow

Read-Host "Press Enter to exit"