<script lang="ts">
    import { Line, getDatasetAtEvent, getElementAtEvent } from "svelte-chartjs";
    import {
        Chart as ChartJS,
        Tooltip,
        LineElement,
        LinearScale,
        PointElement,
        CategoryScale,
    } from "chart.js";

    export let labels = [];
    export let counts = [];
    export let setVideoTime;

    let chartRef;

    let data = {
        labels,
        datasets: [
            {
                data: counts,
                borderColor: "rgb(255,90, 132)",
                backgroundColor: "rgba(255,99,132,0.5)",
            },
        ],
    };

    const options = {
        plugins: {
            legend: {
                display: false,
            },
        },
        responsive: true,
        maintainAspectRatio: false,
    };

    ChartJS.register(
        Tooltip,
        LinearScale,
        PointElement,
        CategoryScale,
        LineElement,
    );

    export function update() {
        chartRef.update();
    }

    function onClick(event) {
        if (!chartRef) {
            return;
        }
        const element = getElementAtEvent(chartRef, event);
        const index = element[0].index;
        const time = index * 30;
        setVideoTime(time);
    }
</script>

<Line {data} {options} bind:chart={chartRef} on:click={onClick} />
