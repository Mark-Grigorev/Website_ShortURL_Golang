function GetInfoShrtURL() {
    const shortUrl = document.getElementById('shortUrl').value;
    const apiURL = "http://localhost:8080/v1/urlshort/statusurl"
        fetch(apiURL, {
        method: 'POST',
        body: shortUrl,
    })
    .then(response => response.json())
    .then(data => {
         // Извлекаем все значения  и объединяем их в одну строку с разделителем \n
           
        const osValues = data.map(item => item.OS).join('\n');
        const userAgentValues = data.map(item => item.UserAgent).join('\n')
        const infoDeviceValues = data.map(item => item.Device).join('\n')
        const ipValues=data.map(item => item.IP).join('\n')

        document.getElementById('infOs').value = osValues;
        document.getElementById('infoUserAgent').value = userAgentValues;
        document.getElementById('infoDevice').value = infoDeviceValues;
        document.getElementById('infoIp').value = ipValues;
    })
    .catch(error => {
        console.error('Error:', error);
    });
}
