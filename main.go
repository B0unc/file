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
func main() {
	fmt.Println("Started")
	dir := GetDirectory()
	TakeUserInputForVideo()
	TakeUserInputForSubtitle()

	FileMap := listFiles(dir) // Print all the files in the current directory
	DebugFileMapOutput(FileMap)

	//Maybe TO-DO sort the file map values. But I think the OS does that already so its not a big issue

	// TO-DO: create a while loop that waits until the user enters a key so they can exit
	fmt.Println("\nPress Enter to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// Get the user specificed directory Current working directory of the program and the folder
func GetDirectory() string {
	UserSpecifiedDirectory, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	FullUserDirectory := filepath.Join(UserSpecifiedDirectory, "video_test")

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

func DebugFileMapOutput(FileMap map[string][]string) {

	println("Entered Debug Output for FileMap")
	keys := make([]string, 0, len(FileMap))

	for ext := range FileMap {
		keys = append(keys, ext)
	}

	for _, ext := range keys {
		fmt.Printf("%s:\n%s\n", ext, strings.Join(FileMap[ext], "\n"))
	}
}

func listFiles(dir string) map[string][]string {

	entries, err := os.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	var files []string
	FileMap := make(map[string][]string)
	for _, v := range entries {
		files = append(files, v.Name())
		extension := filepath.Ext(v.Name())
		FileMap[extension] = append(FileMap[extension], v.Name())
	}

	return FileMap
}
