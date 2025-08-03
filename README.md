# 🛒 Shopping Cart E-Commerce Platform

A complete e-commerce web service built with **Go (Gin)** backend and **React** frontend, implementing the core shopping cart functionality with user authentication, product management, and order processing.

## 🌐 **Live Application Links**

### **🎯 Primary Links (Click to Test)**
- **🚀 Live Frontend**: [https://shop-cart-pi-bice.vercel.app/](https://shop-cart-pi-bice.vercel.app/)
- **🔧 Live Backend API**: [https://shop-cart-production.up.railway.app](https://shop-cart-production.up.railway.app)
- **📹 Demo Video**: [https://drive.google.com/file/d/1zb7pLGY5gmfa2OzqoSPOTC6W15do4bih/view?usp=sharing](https://drive.google.com/file/d/1zb7pLGY5gmfa2OzqoSPOTC6W15do4bih/view?usp=sharing)

### **🔑 Demo Credentials**
- **Username**: `admin`
- **Password**: `artisan123`

## 📋 **Project Overview**

This project implements a basic e-commerce store with the following user flow:
```
User → Add Items to Cart → Order with Items in Cart → Done!
```

### 🎯 **Core Requirements Implemented**

- ✅ **User Authentication**: Signup and login with JWT tokens
- ✅ **Single Session Policy**: One device login per user
- ✅ **Shopping Cart**: Add/remove items with quantity management
- ✅ **Order Processing**: Convert cart to order seamlessly
- ✅ **Product Management**: Browse and manage items
- ✅ **Real-time Updates**: Live cart count and status tracking

## 🛠️ **Technology Stack**

### Backend
- **Go 1.21+** - High-performance programming language
- **Gin** - Lightning-fast HTTP web framework
- **GORM** - Powerful ORM for database operations
- **SQLite** - Reliable database for development and production
- **JWT** - Secure JSON Web Tokens for authentication
- **bcrypt** - Military-grade password hashing

### Frontend
- **React 18** - Modern UI framework with hooks
- **TypeScript** - Type-safe development
- **Vite** - Ultra-fast build tool and dev server
- **Tailwind CSS** - Utility-first CSS framework
- **shadcn/ui** - Beautiful, accessible UI components

### Testing
- **Go Testing** - Comprehensive testing framework
- **testify/assert** - Powerful assertion library
- **SQLite in-memory** - Isolated test database

## 🚀 **Quick Start Guide**

### Prerequisites
- Go 1.21+ installed
- Node.js 18+ installed
- Git

### Step 1: Clone the Repository
```bash
git clone <your-repository-url>
cd Shop-Cart
```

### Step 2: Backend Setup

1. **Navigate to backend directory:**
   ```bash
   cd backend
   ```

2. **Install Go dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the server:**
   ```bash
   go run main.go
   ```
   
   The backend will start on `http://localhost:8080`

### Step 3: Frontend Setup

1. **Navigate to frontend directory:**
   ```bash
   cd frontend
   ```

2. **Install Node.js dependencies:**
   ```bash
   npm install
   ```

3. **Start the development server:**
   ```bash
   npm run dev
   ```
   
   The frontend will start on `http://localhost:3000`

### Step 4: Access the Application

1. **Open your browser** and go to `http://localhost:3000`
2. **Login with demo credentials:**
   - Username: `admin`
   - Password: `artisan123`
3. **Start shopping!** Browse products, add to cart, and place orders

## 🔐 **Authentication System**

### JWT Token Flow

1. **User Signup** (`POST /users`):
   ```json
   {
     "username": "newuser",
     "password": "securepassword"
   }
   ```

2. **User Login** (`POST /users/login`):
   ```json
   {
     "username": "newuser",
     "password": "securepassword"
   }
   ```
   Returns JWT token for authentication.

3. **Use Token**: Include in Authorization header:
   ```
   Authorization: Bearer <your_jwt_token>
   ```

### Single Session Policy
- Each user can only be logged in from **one device at a time**
- New login automatically invalidates previous tokens
- Tokens are stored in the database for validation

## 📊 **Database Schema**

The application uses the following database structure:

### Users Table
- `id` (Primary Key)
- `username` (Unique)
- `password` (Hashed)
- `token` (JWT token)
- `cart_id` (Foreign Key to Carts)
- `created_at`

### Carts Table
- `id` (Primary Key)
- `user_id` (Foreign Key to Users)
- `name`
- `status`
- `created_at`

### Items Table
- `id` (Primary Key)
- `name`
- `description`
- `price`
- `category`
- `rating`
- `reviews`
- `image`
- `in_stock`
- `status`
- `created_at`

### Orders Table
- `id` (Primary Key)
- `cart_id` (Foreign Key to Carts)
- `user_id` (Foreign Key to Users)
- `total`
- `status`
- `created_at`

### Cart_Items Table (Junction)
- `cart_id` (Foreign Key to Carts)
- `item_id` (Foreign Key to Items)
- `quantity`
- `price`

## 🔌 **API Endpoints**

### Base URL: `https://shop-cart-production.up.railway.app`

| Method | Endpoint | Description | Authentication |
|--------|----------|-------------|----------------|
| `POST` | `/users` | Create a new user | ❌ |
| `GET` | `/users` | List all users | ✅ |
| `POST` | `/users/login` | Login user | ❌ |
| `POST` | `/items` | Create an item | ✅ |
| `GET` | `/items` | List all items | ❌ |
| `POST` | `/carts` | Add items to cart | ✅ |
| `GET` | `/carts` | List all carts | ✅ |
| `POST` | `/orders` | Convert cart to order | ✅ |
| `GET` | `/orders` | List all orders | ✅ |

### Authentication Required Endpoints
**Note**: The following endpoints require a valid JWT token in the Authorization header:
- `POST /carts` - User's token identifies which user the cart belongs to
- `POST /orders` - User's token identifies which user the order belongs to

## 🎨 **Frontend Features**

### User Interface Flow

1. **Login Screen**
   - Username and password input
   - Error handling with `window.alert()` for invalid credentials
   - Automatic redirect to products page on success

2. **List Items Screen**
   - Display all available products
   - Click to add items to cart
   - Real-time cart count updates
   - Search and filter functionality

3. **Required UI Elements**
   - **Checkout Button**: Converts cart to order
   - **Cart Button**: Shows cart items in toast notification
   - **Order History Button**: Shows order IDs in toast notification

### User Experience Features
- **Responsive Design**: Works on desktop and mobile
- **Real-time Updates**: Live cart count and status
- **Toast Notifications**: Success/error feedback
- **Loading States**: Smooth user experience
- **Error Handling**: Graceful error management

## 🧪 **Testing**

### Backend Testing
```bash
cd backend
go test ./tests/... -v
```

**Test Coverage:**
- ✅ **User API**: Signup, login, list users (6 tests)
- ✅ **Cart API**: Add/remove items, get cart (6 tests)
- ✅ **Order API**: Create orders, list orders (7 tests)

### Frontend Testing
- TypeScript for type safety
- React Query for data management
- Error boundaries for graceful error handling

## 📁 **Project Structure**

```
Shop-Cart/
├── backend/
│   ├── controllers/     # HTTP request handlers
│   │   ├── user.go     # User authentication
│   │   ├── item.go     # Product management
│   │   ├── cart.go     # Cart operations
│   │   └── order.go    # Order processing
│   ├── models/         # Database models
│   │   ├── user.go     # User model
│   │   ├── item.go     # Item model
│   │   ├── cart.go     # Cart and CartItem models
│   │   └── order.go    # Order model
│   ├── routes/         # API route definitions
│   ├── middlewares/    # Authentication middleware
│   ├── utils/          # Database and JWT utilities
│   ├── tests/          # Test files
│   └── main.go         # Application entry point
├── frontend/
│   ├── src/
│   │   ├── components/ # React components
│   │   ├── services/   # API service functions
│   │   └── App.tsx     # Main application
│   └── package.json    # Frontend dependencies
├── postman/            # API testing collections
└── README.md           # This file
```

## 🚀 **Deployment**

### Backend Deployment (Railway)
- **Platform**: Railway
- **URL**: https://shop-cart-production.up.railway.app
- **Database**: SQLite
- **Environment**: Production

### Frontend Deployment (Vercel)
- **Platform**: Vercel
- **URL**: https://shop-cart-pi-bice.vercel.app/
- **Framework**: Vite + React
- **Environment**: Production

### Environment Variables
```env
PORT=8080
JWT_SECRET=your-secret-key-change-in-production
CORS_ORIGIN=https://shop-cart-pi-bice.vercel.app
```

## 📚 **API Documentation**

### Authentication Endpoints

#### POST /users
**Create a new user account**
```bash
curl -X POST https://shop-cart-production.up.railway.app/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "newuser",
    "password": "securepassword"
  }'
```

#### POST /users/login
**Login and get JWT token**
```bash
curl -X POST https://shop-cart-production.up.railway.app/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "newuser",
    "password": "securepassword"
  }'
```

### Product Endpoints

#### GET /items
**List all available products**
```bash
curl -X GET https://shop-cart-production.up.railway.app/items
```

#### POST /items
**Create a new product (requires authentication)**
```bash
curl -X POST https://shop-cart-production.up.railway.app/items \
  -H "Authorization: Bearer <your_jwt_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "New Product",
    "description": "Product description",
    "price": 29.99,
    "category": "Electronics"
  }'
```

### Cart Endpoints

#### POST /carts
**Add item to cart (requires authentication)**
```bash
curl -X POST https://shop-cart-production.up.railway.app/carts \
  -H "Authorization: Bearer <your_jwt_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "item_id": 1,
    "quantity": 2
  }'
```

#### GET /carts
**Get user's cart (requires authentication)**
```bash
curl -X GET https://shop-cart-production.up.railway.app/carts \
  -H "Authorization: Bearer <your_jwt_token>"
```

### Order Endpoints

#### POST /orders
**Create order from cart (requires authentication)**
```bash
curl -X POST https://shop-cart-production.up.railway.app/orders \
  -H "Authorization: Bearer <your_jwt_token>"
```

#### GET /orders
**List user's orders (requires authentication)**
```bash
curl -X GET https://shop-cart-production.up.railway.app/orders \
  -H "Authorization: Bearer <your_jwt_token>"
```

## 🎯 **Key Features Implemented**

### ✅ Core Requirements
- **User Authentication**: Complete signup/login system
- **Single Session Policy**: One device login per user
- **Shopping Cart**: Add/remove items with quantities
- **Order Processing**: Convert cart to order
- **Product Management**: Browse and manage items
- **Real-time Updates**: Live cart count and status

### ✅ Advanced Features
- **JWT Token Security**: Secure authentication
- **Database Relationships**: Proper foreign key constraints
- **Error Handling**: Comprehensive error management
- **Responsive Design**: Mobile-friendly interface
- **Toast Notifications**: User feedback system
- **Loading States**: Smooth user experience

### ✅ Testing & Documentation
- **Comprehensive Tests**: Backend API testing
- **Postman Collections**: Ready-to-use API tests
- **Clear Documentation**: Step-by-step setup guide
- **Deployment Ready**: Production-ready configuration

## 🔧 **Troubleshooting**

### Common Issues

1. **CORS Errors**:
   - Update CORS origin in backend to match frontend domain
   - Check environment variables

2. **Database Connection**:
   - Ensure SQLite file is writable
   - Check database permissions

3. **Build Failures**:
   - Verify all dependencies are installed
   - Check Node.js and Go versions

4. **Authentication Issues**:
   - Verify JWT token is valid
   - Check token expiration

### Getting Help

1. **Check logs** in your terminal
2. **Verify environment variables** are set correctly
3. **Test API endpoints** using Postman collections
4. **Check browser console** for frontend errors

## 📞 **Support**

For issues or questions:
1. Check the troubleshooting section above
2. Review the API documentation
3. Test with the provided Postman collections
4. Verify all setup steps are completed

---

## 🌟 **Live Application Links**

### **🎯 Primary Links (Click to Test)**
- **🚀 Live Frontend**: [https://shop-cart-pi-bice.vercel.app/](https://shop-cart-pi-bice.vercel.app/)
- **🔧 Live Backend API**: [https://shop-cart-production.up.railway.app](https://shop-cart-production.up.railway.app)
- **📹 Demo Video**: [https://drive.google.com/file/d/1zb7pLGY5gmfa2OzqoSPOTC6W15do4bih/view?usp=sharing](https://drive.google.com/file/d/1zb7pLGY5gmfa2OzqoSPOTC6W15do4bih/view?usp=sharing)

### **🔑 Demo Credentials**
- **Username**: `admin`
- **Password**: `artisan123`

**Built by Vidit Gupta ❤️ using Go, Gin, React, and TypeScript**

*This project implements a complete e-commerce shopping cart system with user authentication, product management, and order processing capabilities.*
