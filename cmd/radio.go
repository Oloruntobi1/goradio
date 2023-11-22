package cmd

import (
	"bufio"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/manifoldco/promptui"
	"github.com/neptunsk1y/goradio/version"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/neptunsk1y/goradio/radio"
	"github.com/spf13/cobra"
)

var clear map[string]func()

var spinnerMusic *spinner.Spinner
var selectStyle = &promptui.SelectTemplates{
	Active:   `> {{ . | blue | bold }}`,
	Inactive: `   {{ . }}`,
	Selected: `{{ "√" | green }} {{ "Playing station" | bold }} {{ . | blue }}`,
	Label:    `{{ . | bold }}`,
}

func init() {
	rootCmd.AddCommand(radioCmd)

	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println("error clear")
		}
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println("error clear")
		}
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println("error clear")
		}
	}
}

var radioCmd = &cobra.Command{
	Use:   "radio",
	Short: "launch radio",
	Run: func(cmd *cobra.Command, args []string) {
		radio.CheckMPV()
		_, err := version.Latest()
		if err != nil {
			fmt.Println("Error version check")
		}
		spinnerMusic = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		spinnerMusic.Suffix = color.GreenString(" Getting stations...")
		errSpinner := spinnerMusic.Color("blue")
		if errSpinner != nil {
			fmt.Println("error setup spinner color")
		}
		ClearScreen()
		RadioAPI := radio.RecordAPI()
		if err := RadioAPI.GetJSON(); err != nil {
			fmt.Println("Error getting json")
			os.Exit(1)
		}
		stationsTitles := []string{}
		stationsLinks := []string{}
		for {
			for _, i := range RadioAPI.RequestResult.Result.Stations {
				stationsTitles = append(stationsTitles, i.Title)
				stationsLinks = append(stationsLinks, i.Stream320)
			}
			prompt := promptui.Select{
				Label:     color.BlueString("?") + " Select station",
				Items:     stationsTitles,
				Templates: selectStyle,
				Size:      10,
				Searcher: func(input string, index int) bool {
					pepper := stationsTitles[index]
					name := strings.Replace(strings.ToLower(pepper), " ", "", -1)
					input = strings.Replace(strings.ToLower(input), " ", "", -1)

					return strings.Contains(name, input)
				},
			}
			fmt.Println("goradio v" + version.Version)
			index, _, err := prompt.Run()
			if err != nil {
				if err.Error() == "^C" {
					ClearScreen()
					return
				}
				log.Fatalf("Error: %v", err.Error())
				return
			}
			stationLink := stationsLinks[index]
			PlayStation(stationLink)
			ClearScreen()
		}
	},
}

func ClearScreen() {
	clearScreen, value := clear[runtime.GOOS]
	if value {
		clearScreen()
	} else {
		fmt.Println("Your terminal not supported clear command")
		os.Exit(2)
	}
}

func changeTitle(title string) {
	spinnerMusic.Suffix = " Playing:                                                                 "
	spinnerMusic.Suffix = " Playing: " + color.BlueString(title)
}

func PlayStation(url string) {
	cmd := exec.Command("mpv", "--no-video", "--no-cache-pause", url)
	signalCmd := make(chan os.Signal, 1)
	signal.Notify(signalCmd, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalCmd
		if spinnerMusic.Active() {
			spinnerMusic.Stop()
		}
		err := cmd.Process.Kill()
		if err != nil {
			fmt.Println("error close mpv")
			os.Exit(1)
		}
	}()
	stdout, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		fmt.Println("error mpv start")
		os.Exit(1)
	}
	stopAnimation := make(chan bool)
	animationSearch := func() {
		for {
			select {
			case <-stopAnimation:
				return
			default:
				changeTitle("_-_-_-_-_-_")
				time.Sleep(time.Millisecond * 250)
				changeTitle("-_-_-_-_-_-")
				time.Sleep(time.Millisecond * 250)
			}
		}
	}
	go animationSearch()
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	spinnerMusic.Start()
	isStopped := false
	for scanner.Scan() {
		m := scanner.Text()
		r := regexp.MustCompile(`icy-title: (.*)`)
		if r.MatchString(m) {
			if !isStopped {
				stopAnimation <- true
				isStopped = true
			}

			title := r.FindStringSubmatch(m)[1]
			if title == " - " {
				go animationSearch()
				time.Sleep(time.Second * 2)
				stopAnimation <- true
			} else {
				changeTitle(title)
			}

		}
	}
	errWait := cmd.Wait()
	if errWait != nil {
		fmt.Println("mpv wait error")
		os.Exit(1)
	}
	signal.Stop(signalCmd)
}
