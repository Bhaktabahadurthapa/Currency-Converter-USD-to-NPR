async function convertCurrency() {
    const usdAmount = document.getElementById('usd-amount').value;
    const resultDiv = document.getElementById('result');

    if (!usdAmount || usdAmount < 0) {
        alert('Please enter a valid amount');
        return;
    }

    try {
        const response = await fetch('/convert', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ amount: parseFloat(usdAmount) })
        });

        const data = await response.json();
        
        resultDiv.style.display = 'block';
        resultDiv.innerHTML = `
            <p><strong>${data.usd_amount.toFixed(2)} USD</strong> = 
            <strong>${data.npr_amount.toFixed(2)} NPR</strong></p>
            <p>Exchange Rate: 1 USD = ${data.rate} NPR</p>
        `;
    } catch (error) {
        resultDiv.style.display = 'block';
        resultDiv.innerHTML = 'Error occurred during conversion';
        console.error('Error:', error);
    }
}
