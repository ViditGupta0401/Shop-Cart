# 🚀 Setup Instructions - Shopping Cart E-Commerce Platform

This guide provides detailed step-by-step instructions to set up and run the Shopping Cart project locally and deploy it to production.

## 🌐 **Live Application Links**

### **🎯 Primary Links (Click to Test)**
- **🚀 Live Frontend**: [https://shop-cart-pi-bice.vercel.app/](https://shop-cart-pi-bice.vercel.app/)
- **🔧 Live Backend API**: [https://shop-cart-production.up.railway.app](https://shop-cart-production.up.railway.app)
- **📹 Demo Video**: [https://drive.google.com/file/d/15LiBifVBaYWgMypvJNIfeRiXlOFeomFR/view?usp=sharing](https://drive.google.com/file/d/15LiBifVBaYWgMypvJNIfeRiXlOFeomFR/view?usp=sharing)

### **🔑 Demo Credentials**
- **Username**: `admin`
- **Password**: `artisan123`

## 📋 Prerequisites

Before starting, ensure you have the following installed:

### Required Software
- **Go 1.21+** - [Download here](https://golang.org/dl/)
- **Node.js 18+** - [Download here](https://nodejs.org/)
- **Git** - [Download here](https://git-scm.com/)

### Verify Installation
```bash
# Check Go version
go version

# Check Node.js version
node --version

# Check npm version
npm --version

# Check Git version
git --version
```

## 🛠️ Step-by-Step Local Setup

### Step 1: Clone the Repository

```bash
# Clone the repository
git clone <your-repository-url>

# Navigate to project directory
cd Shop-Cart

# Verify project structure
ls -la
```

**Expected output:**
```
Shop-Cart/
├── backend/
├── frontend/
├── postman/
├── README.md
└── SETUP_INSTRUCTIONS.md
```

### Step 2: Backend Setup

#### 2.1 Navigate to Backend Directory
```bash
cd backend
```

#### 2.2 Install Go Dependencies
```bash
# Download and install dependencies
go mod tidy

# Verify dependencies
go mod download
```

#### 2.3 Create Environment File
```bash
# Create .env file
touch .env
```

**Add the following to `.env`:**
```env
PORT=8080
JWT_SECRET=your-secret-key-change-in-production
CORS_ORIGIN=*
```

#### 2.4 Run Backend Server
```bash
# Start the server
go run main.go
```

**Expected output:**
```
ArtisanCraft API Server starting on port 8080...
```

**✅ Backend is now running on `http://localhost:8080`**

### Step 3: Frontend Setup

#### 3.1 Open New Terminal
Keep the backend running and open a new terminal window.

#### 3.2 Navigate to Frontend Directory
```bash
cd frontend
```

#### 3.3 Install Node.js Dependencies
```bash
# Install dependencies
npm install

# Verify installation
npm list --depth=0
```

#### 3.4 Start Frontend Development Server
```bash
# Start development server
npm run dev
```

**Expected output:**
```
  VITE v5.x.x  ready in xxx ms

  ➜  Local:   http://localhost:3000/
  ➜  Network: use --host to expose
```

**✅ Frontend is now running on `http://localhost:3000`**

### Step 4: Verify Setup

#### 4.1 Test Backend API
```bash
# Test server status
curl http://localhost:8080/

# Expected response:
# {
#   "message": "Shopping Cart Backend Server is running!",
#   "version": "1.0.0",
#   "status": "active"
# }
```

#### 4.2 Test Frontend
1. **Open browser** and go to `http://localhost:3000`
2. **You should see** the login page
3. **Login with demo credentials:**
   - Username: `admin`
   - Password: `artisan123`

#### 4.3 Test Complete Flow
1. **Login** with demo credentials
2. **Browse products** on the items page
3. **Add items** to cart by clicking "Add" buttons
4. **View cart** by clicking "Cart" button
5. **Checkout** by clicking "Checkout" button
6. **View order history** by clicking "Order History" button

## 🧪 Testing the Application

### Backend Testing
```bash
# Navigate to backend directory
cd backend

# Run all tests
go test ./tests/... -v

# Run specific test categories
go test -run TestUserSignup -v
go test -run TestUserLogin -v
go test -run TestCartAPI -v
go test -run TestCreateOrder -v
```

### Frontend Testing
```bash
# Navigate to frontend directory
cd frontend

# Run linting
npm run lint

# Build for production
npm run build
```

## 📚 API Testing with Postman

### Import Postman Collections
1. **Open Postman**
2. **Import collection** from `postman/Shopping_Cart_API.postman_collection.json`
3. **Set environment variables:**
   - `base_url`: `http://localhost:8080` (for local testing)
   - `base_url`: `https://shop-cart-production.up.railway.app` (for production testing)
   - `auth_token`: (will be auto-filled after login)

### Test API Endpoints
1. **Server Status**: `GET {{base_url}}/`
2. **User Signup**: `POST {{base_url}}/users`
3. **User Login**: `POST {{base_url}}/users/login`
4. **Get Items**: `GET {{base_url}}/items`
5. **Add to Cart**: `POST {{base_url}}/carts`
6. **Get Cart**: `GET {{base_url}}/carts`
7. **Create Order**: `POST {{base_url}}/orders`
8. **Get Orders**: `GET {{base_url}}/orders`

## 🚀 Production Deployment

### Backend Deployment (Railway)

#### Step 1: Create Railway Account
1. **Go to [Railway.app](https://railway.app)**
2. **Click "Start for Free"**
3. **Sign up with GitHub account**
4. **Complete the registration process**

#### Step 2: Deploy Backend
1. **Click "New Project"** on Railway dashboard
2. **Select "Deploy from GitHub repo"**
3. **Connect your GitHub account** (if not already connected)
4. **Select your repository** (Shop-Cart)
5. **Set root directory to `backend`**
6. **Add environment variables:**
   ```
   PORT=8080
   JWT_SECRET=your-super-secret-key-change-in-production
   CORS_ORIGIN=*
   ```
7. **Click "Deploy"**

#### Step 3: Configure Build Settings
1. **Go to "Settings" tab**
2. **Set "Build Command"** to: `go build -o main`
3. **Set "Start Command"** to: `./main`
4. **Click "Save"**

#### Step 4: Get Backend URL
1. **Wait for deployment** to complete (green checkmark)
2. **Copy your Railway URL** (e.g., `https://your-app.railway.app`)

### Frontend Deployment (Vercel)

#### Step 1: Create Vercel Account
1. **Go to [Vercel.com](https://vercel.com)**
2. **Click "Sign Up"**
3. **Sign up with GitHub account**
4. **Complete the registration process**

#### Step 2: Import Project
1. **Click "New Project"** on Vercel dashboard
2. **Click "Import Git Repository"**
3. **Select your GitHub repository** (Shop-Cart)
4. **Click "Import"**

#### Step 3: Configure Frontend Settings
1. **Set Project Settings:**
   - **Framework Preset**: `Vite`
   - **Root Directory**: `frontend`
   - **Build Command**: `npm run build`
   - **Output Directory**: `dist`
   - **Install Command**: `npm install`

2. **Add Environment Variables:**
   - Click "Environment Variables"
   - Add variable:
   ```
   VITE_API_URL=https://your-railway-url.railway.app
   ```

#### Step 4: Deploy Frontend
1. **Click "Deploy"** in Vercel
2. **Wait for deployment** to complete
3. **Copy your Vercel URL** (e.g., `https://your-app.vercel.app`)

### Step 5: Connect Frontend to Backend

#### Update CORS in Railway
1. **Go back to Railway dashboard**
2. **Click on your service**
3. **Click "Variables" tab**
4. **Update CORS_ORIGIN** to your Vercel URL:
   ```
   CORS_ORIGIN=https://your-vercel-app.vercel.app
   ```

#### Update Frontend API URL
1. **Go to Vercel dashboard**
2. **Update environment variable:**
   ```
   VITE_API_URL=https://your-railway-url.railway.app
   ```

## 🔧 Troubleshooting

### Common Issues and Solutions

#### Issue 1: Go Module Errors
```bash
# Error: go: module not found
# Solution:
go mod tidy
go mod download
```

#### Issue 2: Node.js Dependencies
```bash
# Error: Cannot find module
# Solution:
rm -rf node_modules package-lock.json
npm install
```

#### Issue 3: Port Already in Use
```bash
# Error: Address already in use
# Solution:
# For backend (port 8080)
lsof -ti:8080 | xargs kill -9

# For frontend (port 3000)
lsof -ti:3000 | xargs kill -9
```

#### Issue 4: Database Issues
```bash
# Error: database is locked
# Solution:
# Stop the server and restart
# The database will be recreated automatically
```

#### Issue 5: CORS Errors
```bash
# Error: CORS policy
# Solution:
# Check that backend is running on port 8080
# Check that frontend is running on port 3000
# Verify CORS_ORIGIN in .env file
```

### Debug Mode

#### Backend Debug
```bash
# Run with debug logging
GIN_MODE=debug go run main.go
```

#### Frontend Debug
```bash
# Run with verbose logging
npm run dev -- --debug
```

## 📋 Environment Variables

### Backend (.env)
```env
PORT=8080
JWT_SECRET=your-secret-key-change-in-production
CORS_ORIGIN=http://localhost:3000
DB_TYPE=sqlite
```

### Frontend (vite.config.ts)
```typescript
// Update API_BASE_URL in src/services/api.ts
const API_BASE_URL = 'http://localhost:8080'; // Development
const API_BASE_URL = 'https://your-backend-domain.railway.app'; // Production
```

## 🔍 Monitoring and Logs

### Backend Logs
```bash
# View real-time logs
tail -f backend/logs/app.log

# Check server status
curl http://localhost:8080/
```

### Frontend Logs
```bash
# View browser console logs
# Open Developer Tools (F12) in browser
# Check Console tab for errors
```

## 📞 Getting Help

### Check These First
1. **Verify all prerequisites** are installed
2. **Check both servers** are running
3. **Test API endpoints** with Postman
4. **Check browser console** for errors
5. **Review logs** for detailed error messages

### Common Commands
```bash
# Restart backend
cd backend && go run main.go

# Restart frontend
cd frontend && npm run dev

# Clear cache
npm cache clean --force

# Reset database
rm backend/artisancraft.db
```

### Support Resources
- **Backend Documentation**: Check `backend/README.md`
- **Frontend Documentation**: Check `frontend/README.md`
- **API Documentation**: Check `postman/` collections
- **Project README**: Check main `README.md`

---

## 🌟 **Live Application Links**

### **🎯 Primary Links (Click to Test)**
- **🚀 Live Frontend**: [https://shop-cart-pi-bice.vercel.app/](https://shop-cart-pi-bice.vercel.app/)
- **🔧 Live Backend API**: [https://shop-cart-production.up.railway.app](https://shop-cart-production.up.railway.app)
- **📹 Demo Video**: [https://drive.google.com/file/d/15LiBifVBaYWgMypvJNIfeRiXlOFeomFR/view?usp=sharing](https://drive.google.com/file/d/15LiBifVBaYWgMypvJNIfeRiXlOFeomFR/view?usp=sharing)

### **🔑 Demo Credentials**
- **Username**: `admin`
- **Password**: `artisan123`

**✅ Setup Complete! Your Shopping Cart application is now running locally and deployed to production.**

**Next Steps:**
1. Test the complete user flow
2. Explore the API with Postman
3. Run the test suite
4. Share your live URLs!

Thank You...