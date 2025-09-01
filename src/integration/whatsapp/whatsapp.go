package whatsapp

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/DanielHGimenez/yell/src/config"
	"github.com/mdp/qrterminal/v3"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

const (
	databaseFileName       = "wppstore.db"
	databaseFilePathFormat = "file:%s?_foreign_keys=on"
	whatsappLogFile        = "whatsapp_db.log"
	clientLogFile          = "whatsapp_client.log"
)

func CreateClient() (*whatsmeow.Client, error) {
	dbLog, err := NewFileLogger(whatsappLogFile)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	userDatabaseFile, err := config.GetFilePathInExecutableFolder(databaseFileName)
	if err != nil {
		return nil, err
	}
	container, err := sqlstore.New(ctx, "sqlite3", fmt.Sprintf(databaseFilePathFormat, userDatabaseFile), dbLog)
	if err != nil {
		return nil, err
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		return nil, err
	}
	store.DeviceProps.Os = proto.String("Yell")
	clientLog, err := NewFileLogger(clientLogFile)
	if err != nil {
		return nil, err
	}
	client := whatsmeow.NewClient(deviceStore, clientLog)
	return client, nil
}

func IsLoggedIn(client *whatsmeow.Client) bool {
	return client.Store.ID != nil
}

func Login(client *whatsmeow.Client) error {
	qrChan, _ := client.GetQRChannel(context.Background())
	err := client.Connect()
	if err != nil {
		return err
	}
	for evt := range qrChan {
		if evt.Event == "code" {
			// Render the QR code here
			// e.g. qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			// or just manually `echo 2@... | qrencode -t ansiutf8` in a terminal
			qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
		} else {
			log.Println("Login event:", evt.Event)
		}
	}

	return nil
}

func GetGroupByName(client *whatsmeow.Client, groupName string) *types.JID {
	info, err := client.GetJoinedGroups()
	if err != nil {
		log.Fatal("could not check if the channels exists:", err)
	}
	for _, group := range info {
		if group.Name == groupName {
			return &group.JID
		}
	}
	return nil
}

func SendNotification(client *whatsmeow.Client, jid *types.JID, message string) error {
	_, err := client.SendMessage(context.Background(), *jid, &waE2E.Message{
		Conversation: proto.String(message),
	})
	if err != nil {
		return err
	}
	return nil
}
