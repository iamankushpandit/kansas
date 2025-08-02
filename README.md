# Kansas Healthcare Provider Network Analytics Platform

A comprehensive full-stack web application for analyzing healthcare provider networks across Kansas counties, providing data-driven insights for network optimization and strategic planning.

## 🚀 Features

### Interactive Data Visualization
- **Dynamic County Map**: Interactive Highcharts-powered map of Kansas with real-time data visualization
- **Multi-Metric Analysis**: Toggle between Provider Density, Claims Volume, and Network Coverage metrics
- **County-Specific Details**: Click-to-explore functionality with detailed county analytics
- **Responsive Design**: Fully responsive UI built with Vue.js and Vuetify

### Advanced Analytics
- **Specialty Density Analysis**: Automated calculation of provider specialty distribution by county
- **Network Termination Analytics**: Historical analysis of provider departures (2-5 year timeframe)
- **Claims-to-Provider Ratios**: Network coverage efficiency metrics
- **Priority-Based Recommendations**: AI-driven recommendations for network expansion

### Data Export & Reporting
- **PDF Export**: Professional PDF reports with county-specific recommendations
- **Multi-Page Support**: Automatic pagination for comprehensive reports
- **Custom Naming**: Dynamic filename generation with county and date information

### Network Management Tools
- **Provider Filtering**: Filter by specialty, network type (Commercial/Medicare/Tricare)
- **Real-Time Updates**: Dynamic data updates based on filter selections
- **Statewide Overview**: Comprehensive network stability metrics

## 🏗️ System Architecture

### Frontend (Vue.js 3 + Vuetify)
```
├── Vue.js 3 (Composition API)
├── Vuetify 3 (Material Design)
├── Highcharts (Interactive Maps)
├── Axios (HTTP Client)
├── jsPDF + html2canvas (PDF Export)
├── Vite (Build Tool)
└── Modular Component Architecture (11 Components)
```

### Component Architecture
```
App.vue (Main Container)
├── AppHeader.vue (Navigation & Actions)
├── KansasMap.vue (Interactive Map)
│   └── MapLegend.vue (Color Legend)
├── CountyDetails.vue (County Information)
│   ├── CountyMetrics.vue (Statistics)
│   ├── NetworkRecommendations.vue (Strategic Insights)
│   └── SpecialtyDensityAnalysis.vue (Provider Analysis)
└── ControlPanel.vue (Sidebar Controls)
    ├── FilterControls.vue (Data Filters)
    ├── AnalyticsCards.vue (Network Analytics)
    └── MissingDataAlert.vue (Data Warnings)
```

### Backend (Go + Gin Framework)
```
├── Go 1.21
├── Gin Web Framework
├── JSON-based Data Repository
├── RESTful API Design
├── CORS-enabled
└── Modular Architecture
```

### Data Layer
```
├── Provider Data (JSON)
├── Network Associations (JSON)
├── Claims Data (JSON)
├── Service Locations (JSON)
└── County Mappings (JSON)
```

### Deployment
```
├── Docker Containerization
├── Multi-stage Builds
├── Nginx Reverse Proxy
├── Docker Compose Orchestration
└── Production-ready Configuration
```

## 🛠️ Technical Implementation

### Component Design Principles
- **Single Responsibility**: Each component handles one specific feature
- **Props Down, Events Up**: Unidirectional data flow pattern
- **Composition over Inheritance**: Modular component composition
- **Accessibility First**: WCAG 2.1 AA compliance in every component
- **Reusability**: Components designed for reuse across features

### Backend API Endpoints
- `GET /api/v1/county-data` - Retrieve all county statistics
- `GET /api/v1/county-data/:county` - Get specific county data
- `POST /api/v1/filters` - Apply provider filters
- `GET /api/v1/recommendations/:county` - Get county recommendations
- `GET /api/v1/terminated-analysis` - Network termination analysis
- `GET /api/v1/specialty-density/:county` - Specialty density analysis

### Key Algorithms
1. **Density Calculation**: Provider-to-area ratio with dynamic unit selection
2. **Recommendation Engine**: Priority-based specialty gap analysis
3. **Network Stability**: Historical termination rate calculations
4. **Geographic Mapping**: County FIPS code to map coordinate translation

### Data Processing
- **Real-time Filtering**: Dynamic provider filtering by specialty and network
- **Statistical Analysis**: Claims-per-provider ratios and coverage metrics
- **Temporal Analysis**: 2-5 year historical termination tracking

## 📋 Requirements

### System Requirements
- **Docker**: Version 20.0 or higher
- **Docker Compose**: Version 2.0 or higher
- **Memory**: Minimum 2GB RAM
- **Storage**: 500MB available space

### Development Requirements
- **Node.js**: Version 18 or higher
- **Go**: Version 1.21 or higher
- **Git**: For version control

## 🚀 Quick Start

### Prerequisites (Fresh Machine Setup)

#### Option 1: Docker Setup (Recommended)
**Windows:**
1. Install [Docker Desktop for Windows](https://docs.docker.com/desktop/install/windows-install/)
2. Install [Git for Windows](https://git-scm.com/download/win)
3. Restart your computer after installation

**macOS:**
1. Install [Docker Desktop for Mac](https://docs.docker.com/desktop/install/mac-install/)
2. Install Git: `xcode-select --install` (or install Xcode from App Store)

**Linux:**
```bash
# Ubuntu/Debian
sudo apt update
sudo apt install docker.io docker-compose git
sudo systemctl start docker
sudo usermod -aG docker $USER
# Log out and back in
```

#### Option 2: Manual Development Setup
**Windows:**
1. Install [Node.js 18+](https://nodejs.org/en/download/) (includes npm)
2. Install [Go 1.21+](https://golang.org/dl/)
3. Install [Git for Windows](https://git-scm.com/download/win)
4. Install [Visual Studio Code](https://code.visualstudio.com/) (optional)

**macOS:**
1. Install [Homebrew](https://brew.sh/): `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`
2. Install dependencies:
   ```bash
   brew install node@18 go git
   ```
3. Install [Visual Studio Code](https://code.visualstudio.com/) (optional)

**Linux:**
```bash
# Ubuntu/Debian
sudo apt update
sudo apt install nodejs npm golang-go git

# Verify versions
node --version  # Should be 18+
go version     # Should be 1.21+
```

### Running the Application

#### Using Docker (Recommended)
```bash
# Clone the repository
git clone <repository-url>
cd kansas-healthcare-network

# Start the application
docker-compose up --build

# Access the application
# Windows/Linux: http://localhost
# macOS: http://localhost
```

#### Manual Development Setup
```bash
# Clone the repository
git clone <repository-url>
cd kansas-healthcare-network

# Backend setup (Terminal 1)
cd kansas-healthcare-backend
go mod download
go run main.go
# Backend will run on http://localhost:8080

# Frontend setup (Terminal 2 - new terminal window)
cd kansas-healthcare-map
npm install
npm run dev
# Frontend will run on http://localhost:5173
```

### Verification Steps
1. **Docker**: Navigate to `http://localhost` - you should see the Kansas Healthcare map
2. **Manual**: Navigate to `http://localhost:5173` - you should see the application
3. **API Test**: Visit `http://localhost:8080/health` - should return `{"status":"healthy"}`

### Troubleshooting

**Docker Issues:**
- Windows: Ensure Docker Desktop is running and WSL2 is enabled
- macOS: Ensure Docker Desktop is running
- Linux: Check `sudo systemctl status docker`

**Port Conflicts:**
- If port 80 is busy: `docker-compose up --build -p 8080:80`
- If port 5173 is busy: Change port in `vite.config.js`

**Permission Issues (Linux):**
```bash
sudo chown -R $USER:$USER .
sudo chmod -R 755 .
```

## 🧪 Testing

### Health Check Testing
```bash
# Test backend health
curl http://localhost:8080/health

# Test frontend health (via proxy)
curl http://localhost/health
```

### API Testing
```bash
# Test backend endpoints
cd kansas-healthcare-backend
curl http://localhost:8080/api/v1/county-data
curl http://localhost:8080/api/v1/county-data/Sedgwick
```

### Frontend Testing
```bash
# Run development server
cd kansas-healthcare-map
npm run dev

# Build for production
npm run build
```

### Integration Testing
1. Start both services
2. Verify health endpoints respond
3. Navigate to county selection
4. Verify data loading and recommendations
5. Test PDF export functionality
6. Validate filter operations
7. Test graceful shutdown (Ctrl+C)

## 📊 Data Model

### Provider Entity
```json
{
  "provider_id": "string",
  "npi": "string",
  "provider_type": "string",
  "status": "Active|Terminated",
  "county": "string"
}
```

### County Statistics
```json
{
  "county": "string",
  "provider_count": "number",
  "claims_count": "number",
  "avg_claim_amount": "number",
  "density": "string",
  "density_miles": "string"
}
```

## 🔧 Configuration

### Environment Variables
```bash
# Backend
PORT=8080
DATA_SOURCE=json

# Frontend
VITE_API_BASE_URL=http://localhost:8080/api/v1
```

### Docker Configuration
- **Backend**: Runs on port 8080
- **Frontend**: Runs on port 80 with Nginx
- **Network**: Internal Docker bridge network

## 📈 Performance Optimizations

### Frontend
- **Component Splitting**: 11 focused components for better tree-shaking
- **Lazy Loading**: Dynamic imports for PDF libraries
- **Efficient Rendering**: Vuetify's optimized components
- **Modular Architecture**: Smaller bundle sizes and faster loading
- **Caching**: Axios response caching for repeated requests

### Backend
- **JSON Repository**: Fast in-memory data access
- **Efficient Filtering**: Optimized provider filtering algorithms
- **CORS Optimization**: Minimal CORS overhead

### Infrastructure
- **Multi-stage Builds**: Reduced Docker image sizes
- **Nginx Caching**: Static asset caching
- **Gzip Compression**: Reduced payload sizes

## 🔒 Security Features

- **CORS Configuration**: Controlled cross-origin access
- **Input Validation**: Server-side request validation
- **Error Handling**: Comprehensive error management
- **Container Security**: Non-root container execution

## 📋 12-Factor App Compliance

### ✅ **Fully Compliant Factors**
1. **Codebase** - Single repository tracked in Git with multiple deployments
2. **Dependencies** - Explicit dependency declaration via Go modules and npm
3. **Config** - Configuration stored in environment variables (PORT, DATA_SOURCE, API_BASE_URL)
4. **Backing Services** - JSON data treated as attached resources
5. **Build, Release, Run** - Strict separation via Docker multi-stage builds
6. **Processes** - Stateless backend processes with no sticky sessions
7. **Port Binding** - Services export HTTP via port binding (8080, 80)
8. **Concurrency** - Horizontally scalable via Docker container replication
9. **Disposability** - Fast startup/shutdown with graceful termination handling
10. **Dev/Prod Parity** - Docker ensures identical environments across stages
11. **Logs** - Structured logging to stdout as event streams
12. **Admin Processes** - Health check endpoints for administrative tasks

### 🔄 **Demo Limitations**
- **Backing Services**: Uses JSON files instead of external database (acceptable for demo)
- **Admin Processes**: Limited to health checks (sufficient for demo scope)
- **Concurrency**: Single process model (adequate for demo scale)

*Note: Non-compliant factors are intentionally simplified for demo purposes. In production, these would use external databases, comprehensive admin tooling, and advanced concurrency patterns.*

## ♿ Accessibility & Usability Compliance

### ✅ **WCAG 2.1 AA Compliance**
- **Semantic HTML**: Proper heading hierarchy, landmarks, and structure
- **Keyboard Navigation**: Full keyboard accessibility with skip links
- **Screen Reader Support**: ARIA labels, live regions, and descriptive text
- **Color Contrast**: Enhanced contrast ratios meeting AA standards
- **Focus Management**: Visible focus indicators and logical tab order
- **Alternative Text**: Comprehensive alt text and ARIA descriptions
- **Error Handling**: Clear error messages with ARIA live regions

### ✅ **Nielsen's 10 Usability Heuristics**
1. **Visibility of System Status**: Loading states and progress indicators
2. **Match System & Real World**: Intuitive medical terminology and icons
3. **User Control**: Theme toggle, filter controls, and error recovery
4. **Consistency**: Uniform design patterns and interactions
5. **Error Prevention**: Input validation and confirmation dialogs
6. **Recognition vs Recall**: Clear labels and contextual help
7. **Flexibility**: Multiple interaction methods and shortcuts
8. **Aesthetic Design**: Clean, professional medical interface
9. **Error Recovery**: Graceful error handling with recovery options
10. **Help Documentation**: Tooltips and contextual guidance

### ✅ **Inclusive Design Principles**
- **Motor Disabilities**: Large touch targets (44px minimum)
- **Visual Impairments**: High contrast mode and scalable text
- **Cognitive Accessibility**: Clear information hierarchy and progressive disclosure
- **Reduced Motion**: Respects prefers-reduced-motion settings
- **Multiple Input Methods**: Mouse, keyboard, and touch support
- **Screen Reader Optimization**: Structured content with proper semantics

## 📱 Browser Compatibility

- **Chrome**: 90+ (Full accessibility support)
- **Firefox**: 88+ (Full accessibility support)
- **Safari**: 14+ (Full accessibility support)
- **Edge**: 90+ (Full accessibility support)
- **Screen Readers**: NVDA, JAWS, VoiceOver compatible

## 🤝 Contributing

### Development Workflow
1. Fork the repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open Pull Request

### Code Standards
- **Go**: Follow Go formatting standards (`gofmt`)
- **JavaScript**: ESLint configuration included
- **Vue**: Vue 3 Composition API patterns