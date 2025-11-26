import React, { useState, useEffect } from 'react';
import axios from 'axios';

const API_URL = 'http://localhost:8080/api';

const DictionaryPage = () => {
    const [terms, setTerms] = useState([]);
    const [filteredTerms, setFilteredTerms] = useState([]);
    const [searchQuery, setSearchQuery] = useState('');
    const [activeCategory, setActiveCategory] = useState('all');
    const [categories, setCategories] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetchTerms();
    }, []);

    useEffect(() => {
        filterTerms();
    }, [searchQuery, activeCategory, terms]);

    const fetchTerms = async () => {
        try {
            const response = await axios.get(`${API_URL}/terms`);
            const termsData = response.data || [];
            setTerms(termsData);
            setFilteredTerms(termsData);

            // Extract unique categories
            const uniqueCategories = [...new Set(termsData.map(t => t.category))];
            setCategories(uniqueCategories);
        } catch (error) {
            console.error('Error fetching terms:', error);
            setTerms([]);
            setFilteredTerms([]);
        } finally {
            setLoading(false);
        }
    };

    const filterTerms = () => {
        let filtered = terms;

        // Filter by category
        if (activeCategory !== 'all') {
            filtered = filtered.filter(t => t.category === activeCategory);
        }

        // Filter by search query
        if (searchQuery) {
            filtered = filtered.filter(t =>
                t.term.toLowerCase().includes(searchQuery.toLowerCase()) ||
                t.definition.toLowerCase().includes(searchQuery.toLowerCase())
            );
        }

        setFilteredTerms(filtered);
    };

    return (
        <div className="section">
            <div className="container">
                <h1 className="text-center mb-4">Military Dictionary</h1>
                <p className="text-center mb-4" style={{ color: 'var(--color-text-secondary)' }}>
                    Comprehensive glossary of Wehrmacht military terminology and definitions
                </p>

                {/* Search Bar */}
                <div className="search-bar">
                    <input
                        type="text"
                        className="search-input"
                        placeholder="Search terms or definitions..."
                        value={searchQuery}
                        onChange={(e) => setSearchQuery(e.target.value)}
                    />
                </div>

                {/* Category Filter */}
                <div className="tabs" style={{ flexWrap: 'wrap' }}>
                    <button
                        className={`tab ${activeCategory === 'all' ? 'active' : ''}`}
                        onClick={() => setActiveCategory('all')}
                    >
                        All Categories
                    </button>
                    {categories.map((category) => (
                        <button
                            key={category}
                            className={`tab ${activeCategory === category ? 'active' : ''}`}
                            onClick={() => setActiveCategory(category)}
                        >
                            {category}
                        </button>
                    ))}
                </div>

                {/* Terms List */}
                {loading ? (
                    <div className="loading">
                        <div className="spinner"></div>
                    </div>
                ) : filteredTerms.length > 0 ? (
                    <div className="grid grid-2">
                        {filteredTerms.map((term) => (
                            <div key={term.id} className="card">
                                <div style={{ marginBottom: 'var(--spacing-sm)' }}>
                                    <span className="badge badge-heer">{term.category}</span>
                                </div>
                                <h3 className="card-title">{term.term}</h3>
                                <p className="card-text">{term.definition}</p>
                            </div>
                        ))}
                    </div>
                ) : (
                    <div className="text-center" style={{ padding: 'var(--spacing-2xl)' }}>
                        <p style={{ color: 'var(--color-text-secondary)' }}>
                            No terms found matching your criteria.
                        </p>
                    </div>
                )}
            </div>
        </div>
    );
};

export default DictionaryPage;
