<script>
    export let disabled;
</script>

<button
    on:click
    on:mouseover
    on:mouseenter
    on:mouseleave
    {disabled}
    class="cursor-pointer rounded-lg border border-gray-300 bg-white px-4 py-2 font-medium text-gray-500 hover:bg-gray-100 hover:text-gray-700 disabled:cursor-default disabled:opacity-80"
>
    <slot />
</button>
