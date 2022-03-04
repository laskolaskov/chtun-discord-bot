package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var red = color.New(color.FgRed).SprintfFunc()
var whisper = color.New(color.FgMagenta).SprintfFunc()

var whispers = []string{"Death is close...", "You are already dead.", "Your courage will fail.", "Your friends will abandon you.", "You will betray your friends.", "You are weak.", "You will die.", "Your heart will explode."}

func main() {
	godotenv.Load()
	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	defer discord.Close()

	discord.AddHandler(listen)

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(red("\nC'Thun is here."), whisper("There is no escape...\n\n"))
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}

func listen(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	fmt.Println(whisper("%v disturbed C'Thun!", red(m.Author.Username)))

	ch, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		fmt.Println(red("Error"), "while creating DM channel:", err)
	}

	//somethimes C'Thun will not whisper
	if random(0, 10) < 3 {
		fmt.Println(whisper("C'Thun is still sleeping. %v got lucky this time...", red(m.Author.Username)))
		return
	}

	d := time.Duration(random(5, 25)) * time.Second
	fmt.Println(whisper("C'Thun will whisper to %v in %v ...", red(m.Author.Username), whisper(fmt.Sprint(d))))

	time.Sleep(d)
	fmt.Println(whisper("Whispering to %v...", red(m.Author.Username)))
	msg := fmt.Sprintf("```yaml\n%v\n```", whispers[random(0, len(whispers))])
	s.ChannelMessageSend(ch.ID, msg)
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
