package cmd

import (
	"fmt"
	"github.com/Ressetkk/windblume-lyre-player/player"
	"github.com/Ressetkk/windblume-lyre-player/score"
	"github.com/spf13/cobra"
	"log"
)

type playOptions struct {
	score string
	tempo string
}

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play the score.",
	Long: `Play the given score in base64 encoded string.
The score must be encoded base64 JSON string.
After running the program you'll have 10 seconds before it starts generating key events.`,
	Run: func(cmd *cobra.Command, args []string) {
		encodedScore := args[0]
		dec, err := score.Decode(encodedScore)
		if err != nil {
			log.Fatal(err)
		}
		initMsg := "Playing song"
		if dec.Name != "" {
			initMsg += fmt.Sprintf("\"%s\"", dec.Name)
		}
		if dec.Author != "" {
			fmt.Sprintf(", by %s", dec.Author)
		}
		initMsg += "..."
		fmt.Println(initMsg)
		if p, err := player.New(debug); err != nil {
			log.Fatal(err)
		} else {
			if err := p.Play(dec); err != nil {
				log.Fatal(err)
			}
		}

	},
}
