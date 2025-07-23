document.getElementById("clickMe").addEventListener("click", () => {
  document.getElementById("output").innerText = "You clicked the button!";
});

async function fetchCoolTimes() {
    const location = document.getElementById('location').value.trim();
    document.getElementById('results').textContent = "Loading...";
    const resp = await fetch('/api/cooltimes', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ location })
    });
    if (resp.ok) {
        const data = await resp.json();
        // Format output for display
        let output = `Cool times (<=70°F) for ${data.location}:\n\n`;
        for (const day in data.cool_times) {
            output += `${day}: ${data.cool_times[day].join(", ")}\n`;
        }
        document.getElementById('results').textContent = output || "No cool times found.";
    } else {
        document.getElementById('results').textContent = "Error fetching cool times.";
    }
}