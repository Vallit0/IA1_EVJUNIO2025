package main

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/go-vgo/robotgo"
)

var frasesMotivacionales = []string{
	"¡Hoy es un buen día para escribir buen código! 💻",
	"Cada sprint es una oportunidad. ¡Dale con todo! 🚀",
	"Tu mente es tu IDE más poderoso. 🧠",
	"No es magia, es metodología ágil. ✨",
	"Stand-up mental completado. ¡A trabajar! ☕",
}

func reproducirSonido() {
	// Cambia la ruta si tienes un .wav personalizado
	exec.Command("aplay", "/usr/share/sounds/alsa/Front_Center.wav").Start()
}

func mensajeAleatorio() string {
	rand.Seed(time.Now().UnixNano())
	return frasesMotivacionales[rand.Intn(len(frasesMotivacionales))]
}

// hacer endpoints para testear automatizaciones
// cae una requ
func main() {
	// Esto se debe configurar desde el celular
	bot, err := tgbotapi.NewBotAPI("TU_TOKEN_DE_BOT")
	if err != nil {
		log.Panic(err)
	}
	// Configuramos el bot desde Telegram
	log.Printf("Bot autorizado: %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Cuando alguien escribe /inicio
		if update.Message.Text == "/inicio" {
			// go ritualAgil()
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Bienvenido al Curso de IA1 - USAC")
			bot.Send(msg)

			// Ejecutar lo que nosotros querramos
			fmt.Println("👋 Bienvenido al Curso de IA1 - USAC")
			reproducirSonido()
		}

		if update.Message.Text == "/historias" {
			// go ritualAgil()
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "⏳ Ejecutando ritual ágil en tu máquina Linux...")
			bot.Send(msg)
			// Leer el Mensaje que el usuario envio despues de /historias
			fmt.Println("📖 Abriendo Trello para añadir una tarjeta de inicio de jornada")

			exec.Command("xdg-open", "https://trello.com/b/zSKP0j5g/ia1").Start()
			// Simular clic y escritura
			fmt.Println("📝 Añadiendo tarjeta: Inicio de jornada")
			robotgo.MoveMouseSmooth(600, 400) // testearlo con
			robotgo.MouseClick()
			time.Sleep(500 * time.Millisecond)
			robotgo.TypeStr("🌄 Inicio de jornada ágil - " + time.Now().Format("15:04"))
			robotgo.KeyTap("enter")
			robotgo.Alert("Tarjeta añadida", "¡Tarjeta de inicio de jornada creada exitosamente en Trello!")
		}

		if update.Message.Text == "/iniciar-daily" {
			// go ritualAgil()
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "⏳ Ejecutando ritual ágil en tu máquina Linux...")
			bot.Send(msg)
			// Abrir VSCode
			fmt.Println("🔄 Iniciando Daily en VSCode")
			exec.Command("code", "--new-window").Start()
		}

		if update.Message.Text == "/mover" {
			// go ritualAgil()
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "⏳ Ejecutando ritual ágil en tu máquina Linux...")
			bot.Send(msg)
		}
	}
}
