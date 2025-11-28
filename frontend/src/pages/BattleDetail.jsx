import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams, Link } from 'react-router-dom';

const API_URL = '/api';

const BattleDetail = () => {
    const { id } = useParams();
    const [battle, setBattle] = useState(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetchBattleDetails();
    }, [id]);

    const fetchBattleDetails = async () => {
        try {
            const response = await axios.get(`${API_URL}/battles/${id}`);
            setBattle(response.data);
        } catch (error) {
            console.error('Error fetching battle details:', error);
        } finally {
            setLoading(false);
        }
    };

    const getOutcomeBadgeClass = (outcome) => {
        if (outcome.toLowerCase().includes('victory')) {
            return 'badge-kriegsmarine';
        } else if (outcome.toLowerCase().includes('defeat')) {
            return 'badge-heer';
        }
        return 'badge-luftwaffe';
    };

    if (loading) {
        return (
            <div className="section">
                <div className="container">
                    <div className="loading">
                        <div className="spinner"></div>
                    </div>
                </div>
            </div>
        );
    }

    if (!battle) {
        return (
            <div className="section">
                <div className="container text-center">
                    <h2>Battle not found</h2>
                    <Link to="/battles" className="btn mt-3">Back to Battles</Link>
                </div>
            </div>
        );
    }

    return (
        <div className="section">
            <div className="container">
                <Link to="/battles" className="btn mb-3">← Back to Battles</Link>

                <div className="card">
                    <div style={{ marginBottom: 'var(--spacing-md)' }}>
                        <span className={`badge ${getOutcomeBadgeClass(battle.outcome)}`}>
                            {battle.date}
                        </span>
                    </div>

                    <h1 style={{ marginBottom: 'var(--spacing-md)' }}>{battle.name}</h1>

                    <div style={{ marginBottom: 'var(--spacing-lg)' }}>
                        <p style={{ color: 'var(--color-text-muted)' }}>
                            <strong>Location:</strong> {battle.location}
                        </p>
                    </div>

                    <h3 style={{ marginTop: 'var(--spacing-lg)', marginBottom: 'var(--spacing-md)' }}>
                        Description
                    </h3>
                    <p style={{ lineHeight: '1.8', color: 'var(--color-text-secondary)', marginBottom: 'var(--spacing-lg)' }}>
                        {battle.description}
                    </p>

                    <h3 style={{ marginTop: 'var(--spacing-lg)', marginBottom: 'var(--spacing-md)' }}>
                        Participants
                    </h3>
                    <p style={{ color: 'var(--color-text-secondary)', marginBottom: 'var(--spacing-lg)' }}>
                        {battle.participants}
                    </p>

                    <h3 style={{ marginTop: 'var(--spacing-lg)', marginBottom: 'var(--spacing-md)' }}>
                        Outcome
                    </h3>
                    <div>
                        <span className={`badge ${getOutcomeBadgeClass(battle.outcome)}`}>
                            {battle.outcome}
                        </span>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default BattleDetail;
