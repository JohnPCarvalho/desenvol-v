import React, { createContext, useState } from 'react';

import api from '../api/api';

const Context = createContext();

function AuthProvider({ children }) {
    const [authenticated, setAuthenticated] = useState(false);

    async function handleLogin() {
        const { data } = await api.post('/login');

        setAuthenticated(true);
    }

    return (
        <Context.Provider value={{ authenticated, handleLogin }} >
            {children}
        </Context.Provider>
    )
}

export { Context, AuthProvider };