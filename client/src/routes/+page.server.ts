import api from "$lib/server/api";
import { error } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

type RecommendedAnime = {
    data: {
        content: string;
        entry: {
            mal_id: number;
            title: string;
            url: string;
            images: {
                webp: {
                    image_url: string;
                    small_image_url: string;
                    large_image_url: string;
                };
            };
        }[];
    }[];
};

export const load = (async () => {
    const recommended = await api<RecommendedAnime>(
        "https://api.jikan.moe/v4/recommendations/anime",
    );
    if (!recommended.success) {
        throw error(500, "Failed to fetch recommended anime");
    }

    return {
        recommended: recommended.data,
    };
}) satisfies PageServerLoad;
