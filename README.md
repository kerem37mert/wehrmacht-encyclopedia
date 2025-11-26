# Wehrmacht Encyclopedia

A comprehensive web application featuring Wehrmacht history, generals, battles, and military terminology from World War II.

## Features

- **Generals Database**: Detailed biographies of Wehrmacht generals across all branches (Heer, Kriegsmarine, Luftwaffe)
- **Military Dictionary**: Comprehensive glossary of Wehrmacht terminology, ranks, and equipment
- **Battles Archive**: Information about major military operations and battles
- **Daily Quotes**: Rotating quotes from prominent military leaders
- **Search Functionality**: Search across all sections
- **Dark Military Theme**: Immersive dark theme reflecting the historical period

## Tech Stack

### Backend
- **Go** with Echo framework
- **SQLite** database
- RESTful API architecture

### Frontend
- **React** with Vite
- **React Router** for navigation
- **Axios** for API calls
- Custom CSS with dark military theme

## Getting Started

### Prerequisites
- Go 1.21 or higher
- Node.js 18 or higher
- npm

### Backend Setup

```bash
cd backend
go mod download
go run main.go
```

The backend server will start on `http://localhost:8080`

### Frontend Setup

```bash
cd frontend
npm install
npm run dev
```

The frontend will start on `http://localhost:5173` (or another port if 5173 is in use)

## API Endpoints

- `GET /api/generals` - List all generals (optional: `?branch=Heer|Kriegsmarine|Luftwaffe`)
- `GET /api/generals/:id` - Get general details
- `GET /api/terms` - List dictionary terms (optional: `?search=query`)
- `GET /api/battles` - List all battles
- `GET /api/battles/:id` - Get battle details
- `GET /api/quotes/daily` - Get daily rotating quote
- `GET /api/quotes` - List all quotes (optional: `?general_id=id`)
- `GET /api/search?q=query` - Global search across all sections

## Project Structure

```
wehrmacht/
├── backend/
│   ├── database/
│   │   ├── database.go
│   │   └── seed.go
│   ├── handlers/
│   │   └── handlers.go
│   ├── models/
│   │   └── models.go
│   ├── go.mod
│   └── main.go
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   │   └── Navigation.jsx
│   │   ├── pages/
│   │   │   ├── HomePage.jsx
│   │   │   ├── GeneralsPage.jsx
│   │   │   ├── GeneralDetail.jsx
│   │   │   ├── DictionaryPage.jsx
│   │   │   ├── BattlesPage.jsx
│   │   │   └── BattleDetail.jsx
│   │   ├── App.jsx
│   │   ├── main.jsx
│   │   └── index.css
│   ├── index.html
│   └── package.json
└── README.md
```

## Historical Note

This application is created for educational and historical documentation purposes. All information is presented in a factual, historical context.

## License

This project is for educational purposes.
