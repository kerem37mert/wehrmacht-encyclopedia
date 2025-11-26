import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Navigation from './components/Navigation';
import HomePage from './pages/HomePage';
import GeneralsPage from './pages/GeneralsPage';
import GeneralDetail from './pages/GeneralDetail';
import DictionaryPage from './pages/DictionaryPage';
import BattlesPage from './pages/BattlesPage';
import BattleDetail from './pages/BattleDetail';
import './index.css';

function App() {
  return (
    <Router>
      <div className="app">
        <Navigation />
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/generals" element={<GeneralsPage />} />
          <Route path="/generals/:id" element={<GeneralDetail />} />
          <Route path="/dictionary" element={<DictionaryPage />} />
          <Route path="/battles" element={<BattlesPage />} />
          <Route path="/battles/:id" element={<BattleDetail />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
