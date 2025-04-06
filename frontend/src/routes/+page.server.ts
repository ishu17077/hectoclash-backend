import { SERVER, type User } from '$lib';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ cookies, fetch }) => {
    const token = cookies.get('token') || '';

    let user: User | null = null;

    try {
        const response = await fetch(`${SERVER}/user`, {
            headers: {
                token: token
            }
        });

        if (response.ok) {
            const body = await response.json();
            user = {
                email: body.email,
                username: body.username,
                token: token,
                avatar: body.avatar,
                user_id: body.user_id
            }

        } else {
            console.error(await response.text());
        }
    } catch (error) {
        console.error(error);
    }

    return {
        currentUser: user || null
    };
};