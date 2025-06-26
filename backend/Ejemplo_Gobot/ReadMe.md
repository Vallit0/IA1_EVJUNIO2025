## Pasos para Hacer un Bot en Telegram y Rotbo 

1. Descargar Telegram en Celular 
2. Crear un bot con @BotFather 
3. Escribirle /newbot y copiar el Token API del BotFather 
4. Agregar el token a mi programa de Go 
```Go 
bot, err := tgbotapi.NewBotAPI("TU_TOKEN_DE_BOT")
	if err != nil {
		log.Panic(err)
	}
```
