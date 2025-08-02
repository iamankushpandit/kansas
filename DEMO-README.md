# 🏥 Kansas Healthcare Analytics Platform - Hackathon Demo

## 🚀 One-Click Demo Setup (For Non-Technical Users)

### Step 1: Download the Demo
1. Download this entire folder to your Windows laptop
2. Extract all files to a folder (e.g., `C:\kansas-demo\`)

### Step 2: Run the Automated Setup
1. **Right-click** on `setup-demo.ps1`
2. Select **"Run with PowerShell"**
3. If prompted, click **"Yes"** to allow the script to run
4. **Wait 5-10 minutes** for automatic installation and setup

### Step 3: Access the Application
- The script will automatically open your browser to: **http://localhost:4192**
- If it doesn't open automatically, manually navigate to that URL

## 🎯 Demo Features to Showcase

### Interactive Map
- **Click counties** to see detailed healthcare provider analytics
- **Change filters** using the dropdown menus on the right
- **Toggle themes** between light and dark mode

### Key Analytics
- **Provider Density** analysis across Kansas counties
- **Network Termination** analysis (2-5 year historical data)
- **Specialty Distribution** recommendations
- **Claims Volume** and coverage metrics

### Export Capabilities
- **PDF Export** of county-specific recommendations
- **Professional reports** for healthcare network planning

## 🛠️ What the Script Does Automatically

1. **Checks for Docker Desktop** - installs if missing
2. **Starts Docker Desktop** - waits for it to be ready
3. **Builds the application** - creates both frontend and backend
4. **Launches the platform** - opens in your browser
5. **Shows live logs** - for debugging if needed

## 🔧 Manual Setup (If Automated Script Fails)

### Prerequisites
1. Install [Docker Desktop](https://www.docker.com/products/docker-desktop)
2. Start Docker Desktop and wait for it to be ready

### Run Commands
```bash
# Open PowerShell in the project folder
docker-compose up --build

# Access at: http://localhost:4192
```

## 📱 Demo URLs
- **Main Application**: http://localhost:4192
- **Backend API**: http://localhost:3247/health
- **API Documentation**: http://localhost:3247/api/v1/county-data

## 🛑 Stopping the Demo
- Press **Ctrl+C** in the PowerShell window
- Or run: `docker-compose down`

## 🎪 Hackathon Presentation Tips

### Key Talking Points
1. **Healthcare Compliance**: WCAG 2.1 AA accessibility, HIPAA-ready architecture
2. **Performance**: Go backend provides sub-millisecond response times
3. **Scalability**: Docker containerization for easy deployment
4. **Real-world Impact**: Helps healthcare networks optimize provider coverage

### Demo Flow
1. **Start with overview** - Show the Kansas map with provider density
2. **Filter demonstration** - Change network types and specialties
3. **County deep-dive** - Click on a county to show detailed analytics
4. **Export feature** - Generate a PDF report
5. **Technical architecture** - Mention Go + Vue.js + Docker stack

## 🚨 Troubleshooting

### Common Issues
- **Port conflicts**: Close applications using ports 3247 or 4192
- **Docker not starting**: Restart Docker Desktop manually
- **Slow loading**: Wait 2-3 minutes for first-time container builds

### Support
- Check Docker Desktop is running (whale icon in system tray)
- Ensure Windows has administrator privileges for Docker installation
- Verify internet connection for downloading Docker and dependencies

## 🏆 Technical Highlights for Judges

- **Go Backend**: High-performance healthcare data processing
- **Vue.js 3 Frontend**: Modern, accessible healthcare UI
- **Docker Deployment**: Production-ready containerization
- **Healthcare Standards**: FIPS codes, CMS compliance, accessibility
- **Real Analytics**: Terminated provider analysis, specialty density calculations