# StoreMaker - Ecommerce Store Builder Platform

A comprehensive ecommerce platform similar to Shopify that allows users to create and customize their online stores with AI-powered features.

## 🚀 Quick Start

### Prerequisites

Before running the application, make sure you have:

- **Go 1.21+** installed
- **Node.js 18+** and npm installed
- **PostgreSQL 12+** installed and running
- **Git** installed

### 1. Clone & Setup

```bash
git clone <your-repo-url>
cd storemaker
```

### 2. Database Setup

1. **Install PostgreSQL** if not already installed
2. **Create a database**:
   ```sql
   CREATE DATABASE storemaker;
   CREATE USER postgres WITH PASSWORD 'password';
   GRANT ALL PRIVILEGES ON DATABASE storemaker TO postgres;
   ```

### 3. Backend Setup

```bash
# Navigate to backend directory
cd backend

# Install dependencies
go mod tidy

# Create environment file
copy config.env.example .env
# Or on Linux/Mac: cp config.env.example .env

# Edit .env file with your database credentials
# DATABASE_URL=postgres://postgres:password@localhost:5432/storemaker?sslmode=disable

# Run the backend server
go run cmd/main.go
```

The backend will start on `http://localhost:8080`

### 4. Frontend Setup

Open a new terminal window:

```bash
# Navigate to frontend directory
cd frontend

# Install dependencies
npm install

# Create environment file (optional)
# Create .env.local and add:
# NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1

# Run the frontend development server
npm run dev
```

The frontend will start on `http://localhost:3000`

## 🏗️ Architecture

### Backend (Go)
- **Framework**: Gin
- **Database**: PostgreSQL with GORM
- **Authentication**: JWT tokens
- **API**: RESTful endpoints

### Frontend (Next.js)
- **Framework**: Next.js 15 (App Router)
- **Styling**: Tailwind CSS
- **Language**: TypeScript
- **State Management**: React hooks

## 📱 Features

- **User Authentication** (Login/Register)
- **Role-based Access** (Admin, Merchant, Customer)
- **Store Creation & Management**
- **Template System**
- **Product Management**
- **Order Processing**
- **Dark Theme Support**
- **Responsive Design**

## 🔧 Development

### Backend Development
```bash
cd backend
go run cmd/main.go
```

### Frontend Development
```bash
cd frontend
npm run dev
```

### Database Migrations
The backend automatically runs migrations on startup.

## 🌐 API Endpoints

### Authentication
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/refresh` - Refresh token

### Stores
- `GET /api/v1/stores` - Get user stores
- `POST /api/v1/stores` - Create store
- `GET /api/v1/stores/:id` - Get store details
- `PUT /api/v1/stores/:id` - Update store

### Templates
- `GET /api/v1/templates` - Get public templates
- `GET /api/v1/templates/:id` - Get template details

### Products
- `GET /api/v1/stores/:id/products` - Get store products
- `POST /api/v1/stores/:id/products` - Create product

## 🎨 UI/UX

- **Modern Dark Theme** as primary
- **Minimal Gradients** for elegant design
- **Clean Typography** using Inter font
- **Responsive Layout** for all devices
- **Smooth Animations** and transitions

## 🔐 Environment Variables

### Backend (.env)
```env
DATABASE_URL=postgres://postgres:password@localhost:5432/storemaker?sslmode=disable
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
PORT=8080
ENVIRONMENT=development
```

### Frontend (.env.local) - Optional
```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
```

## 🚦 Testing

Access the application:
1. Open `http://localhost:3000` for the frontend
2. Backend API is available at `http://localhost:8080`
3. Health check: `http://localhost:8080/health`

## 📝 Next Steps

1. Set up PostgreSQL database
2. Configure environment variables
3. Run both backend and frontend
4. Test user registration/login
5. Explore the dashboard features

## 🐛 Troubleshooting

### Common Issues

1. **Database Connection Error**
   - Ensure PostgreSQL is running
   - Check database credentials in .env

2. **Frontend Build Errors**
   - Run `npm install` to ensure all dependencies are installed
   - Check Node.js version (requires 18+)

3. **Backend Compilation Errors**
   - Run `go mod tidy` to resolve dependencies
   - Ensure Go version 1.21+

4. **CORS Issues**
   - Backend is configured for `http://localhost:3000`
   - Update CORS settings if using different ports

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.