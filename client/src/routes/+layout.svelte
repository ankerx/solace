<script lang="ts">
    import "../app.css";
    import type { LayoutData } from "./$types";
    import Anime from "./anime.svelte";

    export let data: LayoutData;
</script>

<div class="grid grid-cols-2 p-4">
    <div>
        <slot />
    </div>
    <div class="ml-10">
        <h2 class="my-5 text-xl">My favorites:</h2>
        {#each data.favorites.data as value}
            <Anime
                title={value.title}
                mal_id={value.mal_id}
                image={value.image}
            />
        {/each}
    </div>
</div>
