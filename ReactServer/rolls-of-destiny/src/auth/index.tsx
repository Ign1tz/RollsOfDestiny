import {createAuthProvider} from 'react-token-auth';

type Session = { access_token: string; refreshToken: string };

export const {useAuth, authFetch, login, logout} =
    createAuthProvider<Session>({
        storageKey: 'access_token',
        onUpdateToken: (token) => fetch('http://10.0.0.2:9090/refresh', {
            method: 'POST',
            body: token.access_token,
            headers: {
                        'Accept': 'application/json, text/plain',
                        'Content-Type': 'application/json;charset=UTF-8'
                    },
        })
        .then(r => r.json())
});