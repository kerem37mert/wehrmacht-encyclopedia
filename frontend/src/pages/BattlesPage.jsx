import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

const API_URL = '/api';

const BattlesPage = () => {
    const [battles, setBattles] = useState([]);
    const [filteredBattles, setFilteredBattles] = useState([]);
    const [searchQuery, setSearchQuery] = useState('');
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetchBattles();
    }, []);

    useEffect(() => {
        filterBattles();
    }, [searchQuery, battles]);

    const fetchBattles = async () => {
        try {
            const response = await axios.get(`${API_URL}/battles`);
            setBattles(response.data || []);
            setFilteredBattles(response.data || []);
        } catch (error) {
            console.error('Error fetching battles:', error);
            setBattles([]);
            setFilteredBattles([]);
        } finally {
            setLoading(false);
        }
    };

    const filterBattles = () => {
        if (!searchQuery) {
            setFilteredBattles(battles);
            return;
        }

        const filtered = battles.filter(b =>
            b.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
            b.location.toLowerCase().includes(searchQuery.toLowerCase()) ||
            b.description.toLowerCase().includes(searchQuery.toLowerCase())
        );

        setFilteredBattles(filtered);
    };

    const getOutcomeBadgeClass = (outcome) => {
        if (outcome.toLowerCase().includes('victory')) {
            return 'badge-kriegsmarine';
        } else if (outcome.toLowerCase().includes('defeat')) {
            return 'badge-heer';
        }
        return 'badge-luftwaffe';
    };

    return (
        <div className="section">
            <div className="container">
                <h1 className="text-center mb-4">Battles & Operations</h1>
                <p className="text-center mb-4" style={{ color: 'var(--color-text-secondary)' }}>
                    Major military operations and battles of World War II
                </p>

                {/* Search Bar */}
                <div className="search-bar">
                    <input
                        type="text"
                        className="search-input"
                        placeholder="Search battles by name, location, or description..."
                        value={searchQuery}
                        onChange={(e) => setSearchQuery(e.target.value)}
                    />
                </div>

                {/* Battles Grid */}
                {loading ? (
                    <div className="loading">
                        <div className="spinner"></div>
                    </div>
                ) : filteredBattles.length > 0 ? (
                    <div className="grid grid-2">
                        {filteredBattles.map((battle) => (
                            <Link to={`/battles/${battle.id}`} key={battle.id} className="card">
                                <div style={{ marginBottom: 'var(--spacing-sm)' }}>
                                    <span className={`badge ${getOutcomeBadgeClass(battle.outcome)}`}>
                                        {battle.date}
                                    </span>
                                </div>
                                <h3 className="card-title">{battle.name}</h3>
                                <p className="card-subtitle">{battle.location}</p>
                                <p className="card-text">
                                    {battle.description.substring(0, 150)}...
                                </p>
                                <div style={{ marginTop: 'var(--spacing-md)' }}>
                                    <span className={`badge ${getOutcomeBadgeClass(battle.outcome)}`}>
                                        {battle.outcome}
                                    </span>
                                </div>
                            </Link>
                        ))}
                    </div>
                ) : (
                    <div className="text-center" style={{ padding: 'var(--spacing-2xl)' }}>
                        <p style={{ color: 'var(--color-text-secondary)' }}>
                            No battles found matching your criteria.
                        </p>
                    </div>
                )}
            </div>
        </div>
    );
};

export default BattlesPage;
