import { SERVER_URL } from "$env/static/private";
import api from "$lib/server/api";
import { error } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";
export type Favorite = {
    id: string;
    created: Date;
    updated: Date;
    title: string;
    image: string;
    mal_id: number;
};

export const load = (async () => {
    const go_server_favorites = await api<Favorite[]>(
        SERVER_URL + "/favorites",
    );

    if (!go_server_favorites.success) {
        throw error(500, "Failed to fetch favorites anime");
    } else {
        console.log("Loaded favorites from server", go_server_favorites.data);
    }

    return {
        favorites: go_server_favorites,
    };
}) satisfies LayoutServerLoad;
