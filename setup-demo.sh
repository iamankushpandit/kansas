#!/bin/bash

echo "Kansas Healthcare Analytics Platform - Demo Setup"
echo "================================================="
echo ""

# Check for test parameter
if [[ "$1" == "--test" || "$1" == "-t" ]]; then
    echo "[INFO] Running frontend tests with coverage..."
    cd kansas-healthcare-map
    
    if [ ! -d "node_modules" ]; then
        echo "[INFO] Installing test dependencies..."
        npm install
    fi
    
    npm run test:coverage
    echo "[SUCCESS] Tests completed! Coverage report in coverage/index.html"
    echo "[INFO] Cleaning up Docker containers..."
    docker-compose down 2>/dev/null
    read -p "Press Enter to exit"
    exit 0
fi

# Function to install Homebrew
install_homebrew() {
    echo "[INFO] Installing Homebrew..."
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    
    # Add Homebrew to PATH for current session
    if [[ $(uname -m) == "arm64" ]]; then
        eval "$(/opt/homebrew/bin/brew shellenv)"
    else
        eval "$(/usr/local/bin/brew shellenv)"
    fi
}

# Check and install Homebrew if needed
echo "Checking Homebrew..."
if ! command -v brew &> /dev/null; then
    echo "[INFO] Homebrew not found. Installing automatically..."
    install_homebrew
    echo "[SUCCESS] Homebrew installed successfully!"
else
    echo "[OK] Homebrew is available"
fi

# Check and install Docker Desktop if needed
echo "Checking Docker Desktop..."
if ! command -v docker &> /dev/null; then
    echo "[INFO] Docker Desktop not found. Installing automatically..."
    brew install --cask docker
    echo "[SUCCESS] Docker Desktop installed successfully!"
    echo "[INFO] Please start Docker Desktop from Applications and run this script again."
    read -p "Press Enter to exit"
    exit 0
else
    echo "[OK] Docker is available"
fi

# Check if Docker is running
echo "Checking if Docker is running..."
if ! docker ps &> /dev/null; then
    echo "[INFO] Docker Desktop is not running. Starting it..."
    
    # Try to start Docker Desktop
    open -a Docker
    echo "[INFO] Waiting for Docker Desktop to start (up to 60 seconds)..."
    
    # Wait up to 60 seconds for Docker to start
    timeout=60
    elapsed=0
    while [ $elapsed -lt $timeout ]; do
        sleep 5
        elapsed=$((elapsed + 5))
        if docker ps &> /dev/null; then
            echo "[OK] Docker is now running"
            break
        else
            echo "[INFO] Still waiting for Docker... ($elapsed/$timeout seconds)"
        fi
    done
    
    # Final check
    if ! docker ps &> /dev/null; then
        echo "[ERROR] Docker failed to start. Please start Docker Desktop manually."
        read -p "Press Enter to exit"
        exit 1
    fi
else
    echo "[OK] Docker is running"
fi

# Navigate to script directory
cd "$(dirname "$0")"

echo ""
echo "Running tests before build..."
echo ""

# Run frontend tests
echo "[TEST] Running frontend unit tests with coverage..."
cd kansas-healthcare-map
if [ ! -d "node_modules" ]; then
    npm install --silent
fi
npm run test:run -- --coverage --reporter=verbose
if [ $? -ne 0 ]; then
    echo "[ERROR] Frontend tests failed!"
    echo "[INFO] Cleaning up Docker containers..."
    docker-compose down 2>/dev/null
    read -p "Press Enter to exit"
    exit 1
fi
echo "[OK] Frontend tests passed"
cd ..

# Run backend tests
echo "[TEST] Running backend unit tests..."
cd kansas-healthcare-backend
go test ./... -cover -v
if [ $? -ne 0 ]; then
    echo "[ERROR] Backend tests failed!"
    echo "[INFO] Cleaning up Docker containers..."
    docker-compose down 2>/dev/null
    read -p "Press Enter to exit"
    exit 1
fi
echo "[OK] Backend tests passed"
cd ..

echo ""
echo "Test Summary:"
echo "[TESTS] Frontend: Passed"
echo "[TESTS] Backend: Passed"
echo ""
echo "Starting the healthcare platform..."
echo ""

# Stop existing containers
docker-compose down 2>/dev/null

# Check for rebuild parameter
if [[ "$1" == "--build" || "$1" == "-b" ]]; then
    echo "[INFO] Force rebuilding images"
    docker-compose up --build -d
else
    # Check if images exist
    backend_exists=$(docker images kansas-backend -q)
    frontend_exists=$(docker images kansas-frontend -q)
    
    if [[ -n "$backend_exists" && -n "$frontend_exists" ]]; then
        echo "[INFO] Using existing images (add --build to force rebuild)"
        docker-compose up --build -d
    else
        echo "[INFO] Building images (first run)"
        docker-compose up --build -d
    fi
fi

echo ""
echo "[SUCCESS] Application is starting..."
echo ""
echo "Frontend: http://localhost:4192"
echo "Backend: http://localhost:3247"
echo ""
echo "Preparing to launch browser..."
echo ""

# 10-second countdown with ASCII rocket
for i in {10..1}; do
    echo "    ^    "
    echo "   /|\   "
    echo "  / | \  "
    echo " |  |  | "
    echo " |  |  | "
    echo " |_____| Launch in $i seconds..."
    echo "  \___/  "
    echo ""
    sleep 1
    clear
done

echo "    ^    "
echo "   /|\   "
echo "  / | \  "
 echo " |  |  | "
echo " |  |  | "
echo " |_____| LIFTOFF!"
echo "  \___/  "
echo ""
open "http://localhost:4192"

echo ""
echo "[READY] Demo is ready! Browser should open automatically."
echo "[INFO] To stop: run 'docker-compose down'"
echo "[INFO] To run tests: ./setup-demo.sh --test"

read -p "Press Enter to exit"
echo "[INFO] Cleaning up Docker containers..."
docker-compose down 2>/dev/null