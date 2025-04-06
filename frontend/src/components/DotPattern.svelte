<script lang="ts">
    import { type ClassValue, clsx } from "clsx";
    import { twMerge } from "tailwind-merge";

    function cn(...inputs: ClassValue[]) {
        return twMerge(clsx(inputs));
    }

    export let width: string | number = 40;
    export let height: string | number = 40;
    export let x = 0;
    export let y = 0;
    export let cx = 1;
    export let cy = 1;
    export let cr = 1;
    export let fillColor = "rgb(163 163 163 / 0.8)";

    let className: any = "";
    export { className as class };
    let id = crypto.randomUUID().toString().slice(0, 10); // generating a unique ID for Components
</script>

<svg
    aria-hidden="true"
    class={cn("pointer-events-none absolute inset-0 w-full h-screen -z-10", className)}
    {...$$restProps}
    fill={fillColor}
>
    <defs>
        <pattern
            {id}
            {width}
            {height}
            patternUnits="userSpaceOnUse"
            patternContentUnits="userSpaceOnUse"
            {x}
            {y}
        >
            <circle id="pattern-circle" {cx} {cy} r={cr} />
        </pattern>
    </defs>
    <rect width="100%" height="100%" stroke-width={0} fill={`url(#${id})`} />
</svg>
