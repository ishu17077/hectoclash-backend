import { writable } from "svelte/store";

export const SERVER = "http://localhost:8080"

export interface User {
    user_id: string,
    username: string,
    email: string,
    avatar: string,
    token: string
}

export let user = writable<User | null>(null);

export let score = writable<number>(-1);
export let tscore = writable<number>(1);