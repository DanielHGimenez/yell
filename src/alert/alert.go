package alert

import (
	"fmt"
	"log"

	"github.com/DanielHGimenez/yell/src/config"
	"github.com/spf13/pflag"
)

const (
	BEEP              string = "beep"
	SOUND             string = "sound"
	WPP               string = "wpp"
	GROUP             string = "group"
	MESSAGE           string = "message"
	MESSAGE_SHORTHAND string = "m"
)

func ConfigureFlags(flags *pflag.FlagSet) {
	flags.Bool(BEEP, false, "(Default) Set to the notification be a terminal beep.")
	flags.Bool(SOUND, false, "Set to the notification be a sound.")
	flags.Bool(WPP, false, "Set to the notification be sent using Whatsapp.")
	flags.String(GROUP, "", "The name of the Whatsapp group to send the notification to. Required if --wpp is set.")
	flags.StringP(MESSAGE, MESSAGE_SHORTHAND, "", "The message to send for whatsapp. Default is 'Task completed!'")
}

func Execute(flags *pflag.FlagSet) {
	if flags.Changed(BEEP) || (!flags.Changed(SOUND) && !flags.Changed(WPP)) {
		fmt.Print("\a") // Terminal beep
	}
	if flags.Changed(SOUND) {
		configuration, err := config.LoadConfig()
		var soundFile string
		if err == nil && configuration.SoundPath != "" {
			soundFile = configuration.SoundPath
		} else {
			defaultSoundFilePath, err := config.GetDefaultNotificationSoundFilePath()
			if err != nil {
				log.Fatal("could not get default notification sound file path: ", err)
			}
			soundFile = defaultSoundFilePath
		}
		PlaySound(soundFile)
	}
	if flags.Changed(WPP) {
		if !flags.Changed(GROUP) {
			log.Fatal("The 'group' flag is required when using 'wpp' flag.")
		}
		var message string
		if flags.Changed(MESSAGE) {
			msg, err := flags.GetString(MESSAGE)
			if err != nil {
				log.Fatal("failed to get 'message' flag: ", err)
			}
			message = msg
		} else {
			message = "Task completed!"
		}
		group, err := flags.GetString(GROUP)
		if err != nil {
			log.Fatal("failed to get 'group' flag: ", err)
		}
		SendWhatsappNotification(message, group)
	}
}
