async function fetchData() {
    try {
        const response = await fetch('/data');
        const data = await response.json();
        // TODO: Use D3.js or Chart.js to visualize the data
    } catch (error) {
        console.error('Error fetching data:', error);
    }
}

fetchData();