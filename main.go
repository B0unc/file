/*
WHY:

	I need a program to mass rename each subtitle and video file
	Soon I want to remove subtitle and audio tracks from each file
	Hopefully be able to retime each subtitle track

Phases:

	1: Mass Rename of each file to match the subtitle and video tracks
	2: Remove subtitle and audio tracks from the video file
	3: automatically retime the target language subtitle to the audio
		3a: copy the method from the retiming python script
		3b: manually do it
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// TO-DO Create a Enum struct for Video and Subtitle Seleciton

// TO-DO create a argument list that is like this
// [0] - program path it self nothing special
// [1]
/*
	if there is something here we can assume that it is the folder we have to work through
	if there's nothing here we can assume that we are working in the current directory
	-- help list commands
*/

type Config struct {
	VideoUserConfig        int
	VideoUserConfigName    string
	SubtitleUserConfig     int
	SubtitleUserConfigName string
	directory              string
	RenameFileName         string
	TotalNumberOfEpisodes  int
}

// Map for the Extensions
var VideoExtensions = map[int]string{
	1: ".mkv",
	2: ".mp4",
}

var SubtitleExtensions = map[int]string{
	1: ".srt",
	2: ".ass",
}

/*
	Arguement len:
		if there is nothing added the len is 1
		if there is somethere the len is > 1
	Console arguments needed: expect a len 2
		help
			-help
		Directory: should expect a len 3
			-f current
			-f "Some folder to go to"
	We need to [0:2]
*/

func main() {
	UserConfig := Config{}
	HandleArguments(&UserConfig)
	fmt.Println("Started")

	UserConfig.VideoUserConfig = TakeUserInputForVideo()
	UserConfig.VideoUserConfigName = VideoExtensions[UserConfig.VideoUserConfig]
	UserConfig.SubtitleUserConfig = TakeUserInputForSubtitle()
	UserConfig.SubtitleUserConfigName = SubtitleExtensions[UserConfig.SubtitleUserConfig]
	UserConfig.RenameFileName = GetUserFileRename()
	UserConfig.TotalNumberOfEpisodes = GetUserNumberOfEpisodes()

	FileMap := GetUserFiles(&UserConfig) // Print all the files in the current directory
	DebugFileMapOutput(FileMap, &UserConfig)

	//Maybe TO-DO sort the file map values. But I think the OS does that already so its not a big issue

	// TO-DO: create a while loop that waits until the user enters a key so they can exit
	HandleUserExitPhase(0)
}

func HandleArguments(UserConfig *Config) {
	if len(os.Args) < 2 {
		println("No Arguments entered. Try -help to see commands")
		HandleUserExitPhase(-1)
	}

	// Switch case for args 1
	switch os.Args[1] {
	case "-help":
		HandleArgumentHelp()
		HandleUserExitPhase(-1)
	case "-f":
		HandleArgumentsFolderCommand(UserConfig)
	default:
		fmt.Println("Argument not found. Try the -help command for more info")
		HandleUserExitPhase(-1)
	}
}

func HandleArgumentsFolderCommand(UserConfig *Config) {
	if len(os.Args) > 3 {
		fmt.Println("Try putting the folder name in '' if the folder has any spaces.")
		HandleUserExitPhase(-1)
	} else if len(os.Args) == 2 {
		fmt.Println("No folder name was enter. Trying the current directory. If you dont want this try -f 'folder name here'.")
		UserConfig.directory = GetDirectory("")
	} else if len(os.Args) == 3 {
		fmt.Printf("Trying Folder '%s'\n", os.Args[2])
		UserConfig.directory = GetDirectory(os.Args[2])
	} else {
		fmt.Println("Something went wrong check the function HandleArgumentsFolderCommand")
		HandleUserExitPhase(-1)
	}

}

func HandleArgumentHelp() {
	fmt.Println("-help Shows commands to use for the program\n-f use for the destination of the fold\n-f 'folder name' enter the folder name to grab the files from that folder.")
}

// Get the user specificed directory Current working directory of the program and the folder
func GetDirectory(ArgumentFolderCommandName string) string {
	UserSpecifiedDirectory, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	FullUserDirectory := filepath.Join(UserSpecifiedDirectory, ArgumentFolderCommandName)
	info, err := os.Stat(FullUserDirectory)

	if err != nil {
		log.Fatal("Folder ", FullUserDirectory, " cannot be found or is not a directory")
	}

	if !info.IsDir() {
		log.Fatal(FullUserDirectory, " is not a directory")
	}

	// add a check to see if the folder exist

	return FullUserDirectory
}

// TODO Make sure that when the user press the enter key it defaults to the 1 option
func TakeUserInputForVideo() int {
	VideoUserInput := 0
	fmt.Print("(1) MKV\n(2) MP4\nEnter a value for which video file you want to rename: ")
	fmt.Scanln(&VideoUserInput)

	return VideoUserInput
}

// TODO Make sure that when the user press the enter key it defaults to the 1 option
func TakeUserInputForSubtitle() int {
	SubtitleUserInput := 0
	fmt.Print("(1) SRT\n(2) ASS\nEnter a value for which subtitle file you want to rename: ")
	fmt.Scanln(&SubtitleUserInput)

	return SubtitleUserInput
}

func DebugFileMapOutput(FileMap map[string][]string, UserConfig *Config) {

	fmt.Printf("User Config:\n%d user video input '%s' extension name\n%d user subtitle input '%s' extension name\n%s User directory\n'%s' show's name\n%d number of episodes\n",
		UserConfig.VideoUserConfig, UserConfig.VideoUserConfigName,
		UserConfig.SubtitleUserConfig, UserConfig.SubtitleUserConfigName,
		UserConfig.directory,
		UserConfig.RenameFileName,
		UserConfig.TotalNumberOfEpisodes)

	println("Entered Debug Output for FileMap")
	keys := make([]string, 0, len(FileMap))

	for ext := range FileMap {
		keys = append(keys, ext)
	}

	for _, ext := range keys {
		fmt.Printf("%s:\n%s\n", ext, strings.Join(FileMap[ext], "\n"))
	}
}

func GetUserFiles(UserConfig *Config) map[string][]string {

	entries, err := os.ReadDir(UserConfig.directory)

	if err != nil {
		log.Fatal(err, " something went wrong in the getuserfiles function")
	}

	FileMap := make(map[string][]string)
	for _, v := range entries {
		extension := filepath.Ext(v.Name())
		if extension == UserConfig.VideoUserConfigName || extension == UserConfig.SubtitleUserConfigName {
			FileMap[extension] = append(FileMap[extension], v.Name())
		}
	}

	return FileMap
}

func GetUserFileRename() string {
	UserFileRenameInput := ""

	fmt.Println("\nEnter a the name of the show: ")
	fmt.Scanln(&UserFileRenameInput)

	return UserFileRenameInput
}

func GetUserNumberOfEpisodes() int {
	NumberofEpisodesInput := 0
	fmt.Println("\nEnter the number of episodes: ")
	fmt.Scanln(&NumberofEpisodesInput)

	return NumberofEpisodesInput
}
func HanldeFileRenaming(UserConfig *Config) {
	// TODO handle the ranaming
	return
}

func HandleUserExitPhase(OSExitCode int) {
	// TO-DO: create a while loop that waits until the user enters a key so they can exit
	fmt.Println("\nPress Enter to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	os.Exit(OSExitCode)
}
