package alert

import (
	"log"

	"github.com/DanielHGimenez/yell/src/integration/whatsapp"
)

func SendWhatsappNotification(message, groupName string) {
	client, err := whatsapp.CreateClient()
	if err != nil {
		log.Fatal("Error creating WhatsApp client:", err)
	}
	defer client.Disconnect()
	client.Connect()

	if !whatsapp.IsLoggedIn(client) {
		log.Fatal("WhatsApp is not logged in. Use 'config' command to set up.")
	}

	jid := whatsapp.GetGroupByName(client, groupName)
	if jid == nil {
		log.Fatal("Group not found:", groupName)
	}

	err = whatsapp.SendNotification(client, jid, message)
	if err != nil {
		log.Fatal("Error sending WhatsApp notification:", err)
	}
}
