<script lang="ts">
    import { onMount } from "svelte";
    import type { PageData } from "./$types";
    import Anime from "./anime.svelte";
    import Button from "../components/button.svelte";
    import { fade } from "svelte/transition";
    export let data: PageData;
    let currentPage = 1;

    const itemsPerPage = 10;
    $: paginatedData = data.recommended.data;

    function paginateData() {
        const startIndex = (currentPage - 1) * itemsPerPage;
        const endIndex = startIndex + itemsPerPage;
        paginatedData = data.recommended.data.slice(startIndex, endIndex);
    }
    onMount(() => {
        paginateData();
    });

    function nextPage() {
        currentPage += 1;
        paginateData();
    }

    function previousPage() {
        currentPage -= 1;
        paginateData();
    }
</script>

{#each paginatedData as recommendation}
    <div
        class="list mb-10 grid grid-cols-2 rounded border border-gray-500 p-3 shadow-2xl"
    >
        <h1 class="col-span-2 p-2 text-center">
            {recommendation.content}
        </h1>
        {#each recommendation.entry as subRecommendation}
            <div class="mb-2 p-2">
                <Anime
                    title={subRecommendation.title}
                    mal_id={subRecommendation.mal_id}
                    image={subRecommendation.images.webp.image_url}
                />
            </div>
        {/each}
    </div>
{/each}
<div class="flex justify-center gap-4">
    <Button on:click={previousPage} disabled={currentPage === 1}>
        Previous
    </Button>
    <Button
        on:click={nextPage}
        disabled={currentPage + 1 ===
            Math.ceil(data.recommended.data.length / itemsPerPage)}
    >
        Next
    </Button>
</div>
