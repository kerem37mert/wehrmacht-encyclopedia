import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams, Link } from 'react-router-dom';

const API_URL = '/api';

const GeneralDetail = () => {
    const { id } = useParams();
    const [general, setGeneral] = useState(null);
    const [quotes, setQuotes] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetchGeneralDetails();
        fetchGeneralQuotes();
    }, [id]);

    const fetchGeneralDetails = async () => {
        try {
            const response = await axios.get(`${API_URL}/generals/${id}`);
            setGeneral(response.data);
        } catch (error) {
            console.error('Error fetching general details:', error);
        } finally {
            setLoading(false);
        }
    };

    const fetchGeneralQuotes = async () => {
        try {
            const response = await axios.get(`${API_URL}/quotes?general_id=${id}`);
            setQuotes(response.data || []);
        } catch (error) {
            console.error('Error fetching quotes:', error);
        }
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

    if (!general) {
        return (
            <div className="section">
                <div className="container text-center">
                    <h2>General not found</h2>
                    <Link to="/generals" className="btn mt-3">Back to Generals</Link>
                </div>
            </div>
        );
    }

    return (
        <div className="section">
            <div className="container">
                <Link to="/generals" className="btn mb-3">← Back to Generals</Link>

                <div className="card" style={{ marginBottom: 'var(--spacing-xl)' }}>
                    <div style={{ marginBottom: 'var(--spacing-md)' }}>
                        <span className={`badge ${getBranchBadgeClass(general.branch)}`}>
                            {general.branch}
                        </span>
                    </div>

                    <h1 style={{ marginBottom: 'var(--spacing-sm)' }}>{general.name}</h1>
                    <h3 style={{ color: 'var(--color-gold)', marginBottom: 'var(--spacing-lg)' }}>
                        {general.rank}
                    </h3>

                    {general.birth_date && general.death_date && (
                        <p style={{ color: 'var(--color-text-muted)', marginBottom: 'var(--spacing-lg)' }}>
                            <strong>Life:</strong> {general.birth_date} — {general.death_date}
                        </p>
                    )}

                    <h3 style={{ marginTop: 'var(--spacing-lg)', marginBottom: 'var(--spacing-md)' }}>
                        Biography
                    </h3>
                    <p style={{ lineHeight: '1.8', color: 'var(--color-text-secondary)' }}>
                        {general.biography}
                    </p>

                    {general.notable_battles && (
                        <>
                            <h3 style={{ marginTop: 'var(--spacing-lg)', marginBottom: 'var(--spacing-md)' }}>
                                Notable Battles & Operations
                            </h3>
                            <p style={{ color: 'var(--color-text-secondary)' }}>
                                {general.notable_battles}
                            </p>
                        </>
                    )}
                </div>

                {/* Quotes Section */}
                {quotes.length > 0 && (
                    <div>
                        <h2 style={{ marginBottom: 'var(--spacing-lg)' }}>Notable Quotes</h2>
                        {quotes.map((quote) => (
                            <div key={quote.id} className="quote-box" style={{ marginBottom: 'var(--spacing-lg)' }}>
                                <p className="quote-text">{quote.quote_text}</p>
                                <p className="quote-author">
                                    {quote.context && `${quote.context} • `}
                                    {quote.date && `${quote.date}`}
                                </p>
                            </div>
                        ))}
                    </div>
                )}
            </div>
        </div>
    );
};

export default GeneralDetail;
