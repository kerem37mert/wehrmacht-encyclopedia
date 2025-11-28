import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

const API_URL = '/api';

const HomePage = () => {
    const [dailyQuote, setDailyQuote] = useState(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetchDailyQuote();
    }, []);

    const fetchDailyQuote = async () => {
        try {
            const response = await axios.get(`${API_URL}/quotes/daily`);
            setDailyQuote(response.data);
        } catch (error) {
            console.error('Error fetching daily quote:', error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div>
            {/* Hero Section */}
            <section className="hero">
                <div className="container hero-content">
                    <h1 className="hero-title">Wehrmacht Encyclopedia</h1>
                    <p className="hero-subtitle">
                        A comprehensive archive of Wehrmacht history, generals, battles, and military terminology
                    </p>
                </div>
            </section>

            {/* Daily Quote Section */}
            <section className="section">
                <div className="container">
                    <h2 className="text-center mb-4">Quote of the Day</h2>
                    {loading ? (
                        <div className="loading">
                            <div className="spinner"></div>
                        </div>
                    ) : dailyQuote ? (
                        <div className="quote-box">
                            <p className="quote-text">{dailyQuote.quote_text}</p>
                            <p className="quote-author">
                                — {dailyQuote.rank} {dailyQuote.general_name}
                                {dailyQuote.context && ` • ${dailyQuote.context}`}
                            </p>
                        </div>
                    ) : (
                        <p className="text-center">No quote available</p>
                    )}
                </div>
            </section>

            {/* Quick Navigation */}
            <section className="section" style={{ background: 'var(--color-bg-secondary)' }}>
                <div className="container">
                    <h2 className="text-center mb-4">Explore the Archive</h2>
                    <div className="grid grid-3">
                        <Link to="/generals" className="card">
                            <h3 className="card-title">⚔ Generals</h3>
                            <p className="card-text">
                                Explore the biographies, ranks, and achievements of Wehrmacht generals across all branches:
                                Heer (Army), Kriegsmarine (Navy), and Luftwaffe (Air Force).
                            </p>
                        </Link>

                        <Link to="/dictionary" className="card">
                            <h3 className="card-title">📖 Military Dictionary</h3>
                            <p className="card-text">
                                Comprehensive glossary of Wehrmacht military terminology, ranks, equipment, and organizational
                                structures from the Second World War era.
                            </p>
                        </Link>

                        <Link to="/battles" className="card">
                            <h3 className="card-title">⚡ Battles & Operations</h3>
                            <p className="card-text">
                                Detailed accounts of major military operations, battles, and campaigns involving Wehrmacht
                                forces throughout World War II.
                            </p>
                        </Link>
                    </div>
                </div>
            </section>

            {/* Features Section */}
            <section className="section">
                <div className="container">
                    <h2 className="text-center mb-4">Features</h2>
                    <div className="grid grid-4">
                        <div className="card">
                            <h4 className="card-title">Historical Accuracy</h4>
                            <p className="card-text">
                                All information is carefully researched and documented from historical sources.
                            </p>
                        </div>

                        <div className="card">
                            <h4 className="card-title">Comprehensive Database</h4>
                            <p className="card-text">
                                Extensive collection of generals, battles, and military terminology.
                            </p>
                        </div>

                        <div className="card">
                            <h4 className="card-title">Branch Organization</h4>
                            <p className="card-text">
                                Information organized by military branch: Army, Navy, and Air Force.
                            </p>
                        </div>

                        <div className="card">
                            <h4 className="card-title">Search Functionality</h4>
                            <p className="card-text">
                                Powerful search across all sections to find exactly what you need.
                            </p>
                        </div>
                    </div>
                </div>
            </section>
        </div>
    );
};

export default HomePage;
