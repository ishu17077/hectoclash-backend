import { SERVER, user, type User } from '$lib';
import { get } from 'svelte/store';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ cookies, fetch, url }) => {
    const token = cookies.get('token') || '';

    let currentUser: User | null = null;

    try {
        const response = await fetch(`${SERVER}/user`, {
            headers: {
                token: token
            }
        });

        if (response.ok) {
            const body = await response.json();
            currentUser = {
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

    let body = null;

    try {
        console.log(JSON.stringify({
            problems: 1,
            match_type: "SOLO",
            created_by_player_id: get(user)?.user_id
        }))
        const response = await fetch(`${SERVER}/match`, {
            method: 'POST',
            headers: {
                token: token
            },
            body: JSON.stringify({
                problems: 5,
                match_type: "SOLO",
                created_by_player_id: get(user)?.user_id
            })
        });

        if (response.ok) {
            body = await response.json();
            console.log(body);
        } else {
            console.error(await response.text());
        }
    } catch (error) {
        console.error(error);
    }

    return {
        // id: url.searchParams.get("id"),
        // type: url.searchParams.get("type"),
        // sequence: ['1', '2', '3', '4', '5', '6'],
        // operators: ['+', '×', '(', ')'],
        // no: 2,
        // total: 9,
        problems: body,
        currentUser: currentUser || null
    };
};