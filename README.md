## USD to NPR Currency Converter
A simple and efficient web application built with Go (Golang) that converts US Dollars (USD) to Nepalese Rupees (NPR). This application provides real-time currency conversion with a clean, user-friendly interface.

<img width="1202" alt="image" src="https://github.com/user-attachments/assets/6fcbce75-6dca-4b18-8747-4d3ea6396236">

After using the Ingress:
<img width="1395" alt="image" src="https://github.com/user-attachments/assets/a903c4ad-c16c-4c7a-92c7-73fd3209614c">


## Features
```
🚀 Real-time currency conversion
💱 Current exchange rate display
📱 Responsive design for all devices
⚡ Fast and lightweight
🛡️ Input validation and error handling
🎨 Clean and intuitive user interface
```
## Technologies Used
```
Backend: Go (Golang)
Frontend: HTML, CSS, JavaScript
Server: Native Go HTTP server
Architecture: RESTful API
```
## Prerequisites
Before running this application, make sure you have:
```
Go (version 1.16 or higher) installed
Git installed (for cloning the repository)
A web browser
```
## Installation
Clone the repository:
```
git clone https://github.com/Bhaktabahadurthapa/Currency-Converter-USD-to-NPR.git
```
Navigate to the project directory:
```
cd currency-converter
```
Run the application:
```
run main.go
```
Open your web browser and visit:

```
http://localhost:8080
```
## Project Structure
```
currency-converter/
├── main.go              # Main application file
├── templates/
│   └── index.html      # HTML template
├── static/
│   ├── styles.css      # CSS styles
│   └── script.js       # JavaScript functionality
└── README.md           # Project documentation
```

## API Endpoints
Convert Currency
```
URL: /convert
Method: POST
Request Body:

jsonCopy{
    "amount": 1.00
}
```
## Response:
```
jsonCopy{
    "usd_amount": 1.00,
    "npr_amount": 132.50,
    "rate": 132.50
}
```

## Configuration
The current exchange rate is set in main.go. To modify the exchange rate, update the EXCHANGE_RATE constant:
```
goCopyconst EXCHANGE_RATE = 132.50 // 1 USD = 132.50 NPR
```
## License
This project is licensed under the MIT License - see the LICENSE file for details.

















