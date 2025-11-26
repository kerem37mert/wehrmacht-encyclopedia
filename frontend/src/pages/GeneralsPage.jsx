import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

const API_URL = 'http://localhost:8080/api';

const GeneralsPage = () => {
    const [generals, setGenerals] = useState([]);
    const [filteredGenerals, setFilteredGenerals] = useState([]);
    const [activeBranch, setActiveBranch] = useState('all');
    const [searchQuery, setSearchQuery] = useState('');
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetchGenerals();
    }, []);

    useEffect(() => {
        filterGenerals();
    }, [activeBranch, searchQuery, generals]);

    const fetchGenerals = async () => {
        try {
            const response = await axios.get(`${API_URL}/generals`);
            setGenerals(response.data || []);
            setFilteredGenerals(response.data || []);
        } catch (error) {
            console.error('Error fetching generals:', error);
            setGenerals([]);
            setFilteredGenerals([]);
        } finally {
            setLoading(false);
        }
    };

    const filterGenerals = () => {
        let filtered = generals;

        // Filter by branch
        if (activeBranch !== 'all') {
            filtered = filtered.filter(g => g.branch === activeBranch);
        }

        // Filter by search query
        if (searchQuery) {
            filtered = filtered.filter(g =>
                g.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
                g.rank.toLowerCase().includes(searchQuery.toLowerCase()) ||
                g.biography.toLowerCase().includes(searchQuery.toLowerCase())
            );
        }

        setFilteredGenerals(filtered);
    };

    const getBranchBadgeClass = (branch) => {
        switch (branch) {
            case 'Heer':
                return 'badge-heer';
            case 'Kriegsmarine':
                return 'badge-kriegsmarine';
            case 'Luftwaffe':
                return 'badge-luftwaffe';
            default:
                return '';
        }
    };

    return (
        <div className="section">
            <div className="container">
                <h1 className="text-center mb-4">Wehrmacht Generals</h1>
                <p className="text-center mb-4" style={{ color: 'var(--color-text-secondary)' }}>
                    Explore the military leaders of the Wehrmacht across all branches
                </p>

                {/* Search Bar */}
                <div className="search-bar">
                    <input
                        type="text"
                        className="search-input"
                        placeholder="Search generals by name, rank, or biography..."
                        value={searchQuery}
                        onChange={(e) => setSearchQuery(e.target.value)}
                    />
                </div>

                {/* Branch Filter Tabs */}
                <div className="tabs">
                    <button
                        className={`tab ${activeBranch === 'all' ? 'active' : ''}`}
                        onClick={() => setActiveBranch('all')}
                    >
                        All Branches
                    </button>
                    <button
                        className={`tab ${activeBranch === 'Heer' ? 'active' : ''}`}
                        onClick={() => setActiveBranch('Heer')}
                    >
                        Heer (Army)
                    </button>
                    <button
                        className={`tab ${activeBranch === 'Kriegsmarine' ? 'active' : ''}`}
                        onClick={() => setActiveBranch('Kriegsmarine')}
                    >
                        Kriegsmarine (Navy)
                    </button>
                    <button
                        className={`tab ${activeBranch === 'Luftwaffe' ? 'active' : ''}`}
                        onClick={() => setActiveBranch('Luftwaffe')}
                    >
                        Luftwaffe (Air Force)
                    </button>
                </div>

                {/* Generals Grid */}
                {loading ? (
                    <div className="loading">
                        <div className="spinner"></div>
                    </div>
                ) : filteredGenerals.length > 0 ? (
                    <div className="grid grid-3">
                        {filteredGenerals.map((general) => (
                            <Link to={`/generals/${general.id}`} key={general.id} className="card">
                                <div style={{ marginBottom: 'var(--spacing-sm)' }}>
                                    <span className={`badge ${getBranchBadgeClass(general.branch)}`}>
                                        {general.branch}
                                    </span>
                                </div>
                                <h3 className="card-title">{general.name}</h3>
                                <p className="card-subtitle">{general.rank}</p>
                                <p className="card-text">
                                    {general.birth_date && general.death_date && (
                                        <span style={{ display: 'block', marginBottom: 'var(--spacing-sm)', color: 'var(--color-text-muted)' }}>
                                            {general.birth_date} — {general.death_date}
                                        </span>
                                    )}
                                    {general.biography.substring(0, 150)}...
                                </p>
                            </Link>
                        ))}
                    </div>
                ) : (
                    <div className="text-center" style={{ padding: 'var(--spacing-2xl)' }}>
                        <p style={{ color: 'var(--color-text-secondary)' }}>
                            No generals found matching your criteria.
                        </p>
                    </div>
                )}
            </div>
        </div>
    );
};

export default GeneralsPage;
