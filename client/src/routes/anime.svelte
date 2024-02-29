<script lang="ts">
    export let title: string;
    export let mal_id: number;
    export let image: string;
</script>

<div
    class="h-max[500px] flex max-w-sm flex-col gap-2 rounded-lg border border-gray-200 bg-white p-6 shadow dark:border-gray-700 dark:bg-gray-800"
>
    <a href="/{mal_id}">
        <h5
            class="mb-2 text-xl font-bold tracking-tight text-gray-900 dark:text-white"
        >
            {title}
        </h5>
    </a>

    <a href="/{mal_id}">
        <img
            class="hover:cursor-pointer hover:opacity-50"
            src={image}
            alt={title}
            height="200px"
            width="200px"
        />
    </a>
</div>
