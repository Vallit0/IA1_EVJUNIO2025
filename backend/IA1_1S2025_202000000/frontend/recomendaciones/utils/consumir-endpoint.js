async function consumirEndpoint(url, data) {
    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });
        if (!response.ok) {
            throw new Error(`Error: ${response.status}`);
        }
        return await response.json();
    } catch (error) {
        console.error('Error al consumir el endpoint:', error);
        throw error;
    }
}

// Ejemplo de uso:
// consumirEndpoint('https://api.ejemplo.com/endpoint', { clave: 'valor' })
//   .then(respuesta => console.log(respuesta))
//   .catch(error => console.error(error));

export default consumirEndpoint;