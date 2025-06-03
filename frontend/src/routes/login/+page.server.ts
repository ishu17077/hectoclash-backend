import { goto } from '$app/navigation';
import { SERVER, user, type User } from '$lib';
import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions = {
    login: async ({ cookies, request, url }) => {
        const data = await request.formData();
        const email = data.get('email');
        const password = data.get('password');

        try {
            const response = await fetch(`${SERVER}/login`, {
                method: "POST",
                body: JSON.stringify({
                    email: email,
                    password: password
                })
            })

            if (response.ok) {
                const body = await response.json();
                console.log(body);

                cookies.set('token', body.token, { path: "/" })
                console.log(user);

                throw redirect(307, "/game")
            } else {
                console.error(response.text());
            }
        } catch (error) {
            console.error(error);
        }
    },
    register: async ({ request, cookies }) => {
        const data = await request.formData();
        const username = data.get('username');
        const email = data.get('email');
        const password = data.get('password');

        try {
            const response = await fetch(`${SERVER}/signin`, {
                method: "POST",
                body: JSON.stringify({
                    username: username,
                    email: email,
                    password: password
                })
            })

            if (response.ok) {
                const body = await response.json();
                cookies.set('token', body.token, { path: "/" })
                console.log(body);
            } else {
                console.error(response.text());
            }
        } catch (error) {
            console.error(error);
        }
    }
} satisfies Actions;