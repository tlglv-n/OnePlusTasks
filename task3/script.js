document.addEventListener("DOMContentLoaded", function() {
    fetch("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1")
        .then(response => response.json())
        .then(data => {
            displayCryptoCurrencies(data);
        })
        .catch(error => {
            console.error("Error fetching data:", error);
        });
});

function displayCryptoCurrencies(currencies) {
    const tableDiv = document.getElementById("cryptoTable");
    const table = document.createElement("table");
    const headerRow = table.insertRow();
    const headers = ["ID", "Symbol", "Name"];
    headers.forEach(headerText => {
        const header = document.createElement("th");
        header.textContent = headerText;
        headerRow.appendChild(header);
    });

    currencies.forEach((currency, index) => {
        const row = table.insertRow();
        const isUSDT = currency.symbol.toLowerCase() === "usdt";
        const bgColor = isUSDT ? "green" : (index < 5 ? "blue" : "white");

        row.style.backgroundColor = bgColor;

        const idCell = row.insertCell();
        idCell.textContent = currency.id;

        const symbolCell = row.insertCell();
        symbolCell.textContent = currency.symbol;

        const nameCell = row.insertCell();
        nameCell.textContent = currency.name;
        if (isUSDT) {
            idCell.style.backgroundColor = "blue";
            nameCell.style.backgroundColor = "blue";
            row.style.backgroundColor = "green";
        }
    });

    tableDiv.appendChild(table);
}
