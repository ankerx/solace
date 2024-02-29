import api from "$lib/server/api";
import { error } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import { SERVER_URL } from "$env/static/private";
import * as z from "zod";
import { fail } from "@sveltejs/kit";
const schema = z.object({
    mal_id: z.number(),
    title: z.string(),
    image: z.string(),
});

export type Anime = {
    data: {
        mal_id: number;
        title: string;
        url: string;
        synopsis: string;
        images: {
            webp: {
                image_url: string;
                small_image_url: string;
                large_image_url: string;
            };
        };
    };
};

export const load = (async ({ params }) => {
    const id = params.anime_id;
    const anime = await api<Anime>(`https://api.jikan.moe/v4/anime/${id}`);
    if (!anime.success) {
        console.error("Failed to fetch anime", anime.error);
        throw error(500, "Failed to fetch anime");
    }

    return {
        anime: anime.data.data,
    };
}) satisfies PageServerLoad;

export const actions = {
    addToFavorites: async ({ request }) => {
        const form = await request.formData();

        const mal_id = Number(form.get("mal_id"));
        const title = form.get("title");
        const image = form.get("image");

        try {
            const data = schema.parse({
                mal_id,
                title,
                image,
            });
            const go_server_add_favorites = await api(
                SERVER_URL + "/favorites",
                {
                    method: "POST",
                    body: data,
                },
            );
            if (!go_server_add_favorites.success) {
                console.error(
                    "Failed to add to favorites",
                    go_server_add_favorites.error,
                );
            } else {
                console.log("Added to favorites", go_server_add_favorites.data);
            }
            return { success: true };
        } catch (error) {
            console.log(error);

            if (error instanceof z.ZodError) {
                return fail(400, {
                    message: error.issues,
                });
            }
            return fail(400, {
                message: "Failed to add to favorites!",
            });
        }
    },
} satisfies Actions;
