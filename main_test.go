package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// Test the get directory function

/*
Test
	func GetUserFiles(UserConfig *Config) map[string][]string
*/

func TestGetUserFiles(t *testing.T) {

	// Build the test structure
	/*
		- name
		- files []string
		- video
		- sub
		- want

	*/
	tests := []struct {
		name                 string
		files                []string
		VideoConfigExtension string
		SubConfigExtension   string
		want                 map[string][]string
	}{
		/*
			Matching both MKV and SRT
			Matching both MKV and ASS
			Matching both MP4 and SRT
			Matching both MP4 and ASS
		*/
		{
			name:                 "mkv and srt only",
			files:                []string{"a.mkv", "b.mkv", "c.mkv", "a.srt", "b.srt", "c.srt"},
			VideoConfigExtension: ".mkv",
			SubConfigExtension:   ".srt",
			want: map[string][]string{
				".mkv": {"a.mkv", "b.mkv", "c.mkv"},
				".srt": {"a.srt", "b.srt", "c.srt"},
			},
		},
		{
			name:                 "mkv and ass only",
			files:                []string{"a.mkv", "b.mkv", "c.mkv", "a.ass", "b.ass", "c.ass"},
			VideoConfigExtension: ".mkv",
			SubConfigExtension:   ".ass",
			want: map[string][]string{
				".mkv": {"a.mkv", "b.mkv", "c.mkv"},
				".ass": {"a.ass", "b.ass", "c.ass"},
			},
		},
		{
			name:                 "mp4 and srt only",
			files:                []string{"a.mp4", "b.mp4", "c.mp4", "a.srt", "b.srt", "c.srt"},
			VideoConfigExtension: ".mp4",
			SubConfigExtension:   ".srt",
			want: map[string][]string{
				".mp4": {"a.mp4", "b.mp4", "c.mp4"},
				".srt": {"a.srt", "b.srt", "c.srt"},
			},
		},
		{
			name:                 "mp4 and ass only",
			files:                []string{"a.mp4", "b.mp4", "c.mp4", "a.ass", "b.ass", "c.ass"},
			VideoConfigExtension: ".mp4",
			SubConfigExtension:   ".ass",
			want: map[string][]string{
				".mp4": {"a.mp4", "b.mp4", "c.mp4"},
				".ass": {"a.ass", "b.ass", "c.ass"},
			},
		},
		/*
			No Matching MKV and SRT
			No Matching MKV and ASS
			No Matching MP4 and SRT
			No Matching MP4 and ASS
		*/
		{
			name:                 "No matching MKV and SRT files",
			files:                []string{"a.mp4", "b.mp4", "c.mp4", "a.ass", "b.ass", "c.ass"},
			VideoConfigExtension: ".mkv",
			SubConfigExtension:   ".srt",
			want:                 map[string][]string{},
		},
		{
			name:                 "No matching MKV and ASS files",
			files:                []string{"a.mp4", "b.mp4", "c.mp4", "a.srt", "b.srt", "c.srt"},
			VideoConfigExtension: ".mkv",
			SubConfigExtension:   ".ass",
			want:                 map[string][]string{},
		},
		{
			name:                 "No matching MP4 and SRT files",
			files:                []string{"a.mkv", "b.mkv", "c.mkv", "a.ass", "b.ass", "c.ass"},
			VideoConfigExtension: ".mp4",
			SubConfigExtension:   ".srt",
			want:                 map[string][]string{},
		},
		{
			name:                 "No matching MP4 and ASS files",
			files:                []string{"a.mkv", "b.mkv", "c.mkv", "a.srt", "b.srt", "c.srt"},
			VideoConfigExtension: ".mp4",
			SubConfigExtension:   ".ass",
			want:                 map[string][]string{},
		},
		/*
			Match MKV, but no matching SRT
			Match MKV, but no matching ASS
			Match MP4, but no matching SRT
			Match MP4, but no matching ASS
		*/
		{
			name:                 "Matching MKV, but No matching SRT",
			files:                []string{"a.mkv", "b.mkv", "c.mkv", "a.ass", "b.ass", "c.ass"},
			VideoConfigExtension: ".mkv",
			SubConfigExtension:   ".srt",
			want: map[string][]string{
				".mkv": {"a.mkv", "b.mkv", "c.mkv"},
			},
		},
		{
			name:                 "Matching MKV, but No matching ASS",
			files:                []string{"a.mkv", "b.mkv", "c.mkv", "a.srt", "b.srt", "c.srt"},
			VideoConfigExtension: ".mkv",
			SubConfigExtension:   ".ass",
			want: map[string][]string{
				".mkv": {"a.mkv", "b.mkv", "c.mkv"},
			},
		},
		{
			name:                 "Matching MP4, but No matching SRT",
			files:                []string{"a.mp4", "b.mp4", "c.mp4", "a.ass", "b.ass", "c.ass"},
			VideoConfigExtension: ".mp4",
			SubConfigExtension:   ".srt",
			want: map[string][]string{
				".mp4": {"a.mp4", "b.mp4", "c.mp4"},
			},
		},
		{
			name:                 "Matching MP4, but No matching ASS",
			files:                []string{"a.mp4", "b.mp4", "c.mp4", "a.srt", "b.srt", "c.srt"},
			VideoConfigExtension: ".mp4",
			SubConfigExtension:   ".ass",
			want: map[string][]string{
				".mp4": {"a.mp4", "b.mp4", "c.mp4"},
			},
		},
		/*
			Matching MKV, empty SRT
			Matching MKV, empty ASS
			Matching MP4, empty SRT
			Matching MP4, empty ASS
		*/
		{
			name:                 "Matching MKV, empty SRT",
			files:                []string{"a.mkv", "b.mkv", "c.mkv"},
			VideoConfigExtension: ".mkv",
			SubConfigExtension:   ".srt",
			want: map[string][]string{
				".mkv": {"a.mkv", "b.mkv", "c.mkv"},
			},
		},
		{
			name:                 "Matching MKV, empty ASS",
			files:                []string{"a.mkv", "b.mkv", "c.mkv"},
			VideoConfigExtension: ".mkv",
			SubConfigExtension:   ".ass",
			want: map[string][]string{
				".mkv": {"a.mkv", "b.mkv", "c.mkv"},
			},
		},
		{
			name:                 "Matching MP4 empty SRT",
			files:                []string{"a.mp4", "b.mp4", "c.mp4"},
			VideoConfigExtension: ".mp4",
			SubConfigExtension:   ".srt",
			want: map[string][]string{
				".mp4": {"a.mp4", "b.mp4", "c.mp4"},
			},
		},
		{
			name:                 "Matching MP4 empty ASS",
			files:                []string{"a.mp4", "b.mp4", "c.mp4"},
			VideoConfigExtension: ".mp4",
			SubConfigExtension:   ".ass",
			want: map[string][]string{
				".mp4": {"a.mp4", "b.mp4", "c.mp4"},
			},
		},
		/*
			Matching SRT, but no matching MKV
			Matching SRT, but no matching MP4
			Matching ASS, but no matching MKV
			Matching ASS, but no matching MP4
		*/
		{
			name:                 "Matching SRT, but No match MKV",
			files:                []string{"a.mp4", "b.mp4", "c.mp4", "a.srt", "b.srt", "c.srt"},
			VideoConfigExtension: ".mkv",
			SubConfigExtension:   ".srt",
			want: map[string][]string{
				".srt": {"a.srt", "b.srt", "c.srt"},
			},
		},
		{
			name:                 "Matching SRT, but no matching MP4",
			files:                []string{"a.mkv", "b.mkv", "c.mkv", "a.srt", "b.srt", "c.srt"},
			VideoConfigExtension: ".mp4",
			SubConfigExtension:   ".srt",
			want: map[string][]string{
				".srt": {"a.srt", "b.srt", "c.srt"},
			},
		},
		{
			name:                 "Matching ASS, but no matching MKV",
			files:                []string{"a.mp4", "b.mp4", "c.mp4", "a.ass", "b.ass", "c.ass"},
			VideoConfigExtension: ".mkv",
			SubConfigExtension:   ".ass",
			want: map[string][]string{
				".ass": {"a.ass", "b.ass", "c.ass"},
			},
		},
		{
			name:                 "Matching ASS, but No matching MP4",
			files:                []string{"a.mkv", "b.mkv", "c.mkv", "a.ass", "b.ass", "c.ass"},
			VideoConfigExtension: ".mp4",
			SubConfigExtension:   ".ass",
			want: map[string][]string{
				".ass": {"a.ass", "b.ass", "c.ass"},
			},
		},
		/*
			Matching SRT, empty MKV
			Matching SRT, empty MP4
			Matching ASS, empty MKV
			Matching ASS, empty MP4
		*/
		{
			name:                 "Matching SRT, empty MKV",
			files:                []string{"a.srt", "b.srt", "c.srt"},
			VideoConfigExtension: ".mkv",
			SubConfigExtension:   ".srt",
			want: map[string][]string{
				".srt": {"a.srt", "b.srt", "c.srt"},
			},
		},
		{
			name:                 "Matching SRT, empty MP4",
			files:                []string{"a.srt", "b.srt", "c.srt"},
			VideoConfigExtension: ".mp4",
			SubConfigExtension:   ".srt",
			want: map[string][]string{
				".srt": {"a.srt", "b.srt", "c.srt"},
			},
		},
		{
			name:                 "Matching ASS, empty MKV",
			files:                []string{"a.ass", "b.ass", "c.ass"},
			VideoConfigExtension: ".mkv",
			SubConfigExtension:   ".ass",
			want: map[string][]string{
				".ass": {"a.ass", "b.ass", "c.ass"},
			},
		},
		{
			name:                 "Matching ASS, empty MP4",
			files:                []string{"a.ass", "b.ass", "c.ass"},
			VideoConfigExtension: ".mp4",
			SubConfigExtension:   ".ass",
			want: map[string][]string{
				".ass": {"a.ass", "b.ass", "c.ass"},
			},
		},
		{
			name:                 "none",
			files:                []string{},
			VideoConfigExtension: ".mkv",
			SubConfigExtension:   ".srt",
			want:                 map[string][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir() // Temp directory for test

			// Writing to the temp directory
			for _, f := range tt.files {
				// 0o644; 0o is the hex 0x and 644 sets the permissions for owner and the group can read and write to the file, while everybody else can only read
				if err := os.WriteFile(filepath.Join(dir, f), nil, 0o644); err != nil {
					t.Fatal(err)
				}
			}

			// What we need
			/*
				- UserConifg
					- Directory
					- VideoConfigName
					- SubtitleConfigName
			*/

			cfg := &Config{
				directory:              dir,
				VideoUserConfigName:    tt.VideoConfigExtension,
				SubtitleUserConfigName: tt.SubConfigExtension,
			}

			got := GetUserFiles(cfg)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

/*
Test

	func HanldeFileRenaming(UserConfig *Config, FileMap map[string][]string)
*/
func TestHandleFileRenaming(t *testing.T) {
	tests := []struct {
		name                      string
		OriginalFiles             []string
		VideoConfigExtension      string
		SubConfigExtension        string
		TestRenameFileName        string
		TestTotalNumberOfEpisodes int
		want                      map[string][]string
		wantErr                   bool
	}{
		{
			name:                      "Rename MKV and SRT Files",
			OriginalFiles:             []string{"a.mkv", "b.mkv", "c.mkv", "a.srt", "b.srt", "c.srt"},
			VideoConfigExtension:      ".mkv",
			SubConfigExtension:        ".srt",
			TestRenameFileName:        "Show",
			TestTotalNumberOfEpisodes: 3,
			want: map[string][]string{
				".mkv": {"Show - 01.mkv", "Show - 02.mkv", "Show - 03.mkv"},
				".srt": {"Show - 01.srt", "Show - 02.srt", "Show - 03.srt"},
			},
		},
		{
			name:                      "Rename MKV and ASS Files",
			OriginalFiles:             []string{"a.mkv", "b.mkv", "c.mkv", "a.ass", "b.ass", "c.ass"},
			VideoConfigExtension:      ".mkv",
			SubConfigExtension:        ".ass",
			TestRenameFileName:        "testing_the_file",
			TestTotalNumberOfEpisodes: 3,
			want: map[string][]string{
				".mkv": {"testing_the_file - 01.mkv", "testing_the_file - 02.mkv", "testing_the_file - 03.mkv"},
				".ass": {"testing_the_file - 01.ass", "testing_the_file - 02.ass", "testing_the_file - 03.ass"},
			},
		},
		{
			name:                      "Rename MP4 and ASS Files",
			OriginalFiles:             []string{"a.mp4", "b.mp4", "c.mp4", "a.ass", "b.ass", "c.ass"},
			VideoConfigExtension:      ".mp4",
			SubConfigExtension:        ".ass",
			TestRenameFileName:        "The Show",
			TestTotalNumberOfEpisodes: 3,
			want: map[string][]string{
				".mp4": {"The Show - 01.mp4", "The Show - 02.mp4", "The Show - 03.mp4"},
				".ass": {"The Show - 01.ass", "The Show - 02.ass", "The Show - 03.ass"},
			},
		},
		{
			name:                      "Rename MP4 and SRT Files",
			OriginalFiles:             []string{"a.mp4", "b.mp4", "c.mp4", "a.srt", "b.srt", "c.srt"},
			VideoConfigExtension:      ".mp4",
			SubConfigExtension:        ".srt",
			TestRenameFileName:        "Hibi whatever_2",
			TestTotalNumberOfEpisodes: 3,
			want: map[string][]string{
				".mp4": {"Hibi whatever_2 - 01.mp4", "Hibi whatever_2 - 02.mp4", "Hibi whatever_2 - 03.mp4"},
				".srt": {"Hibi whatever_2 - 01.srt", "Hibi whatever_2 - 02.srt", "Hibi whatever_2 - 03.srt"},
			},
		},
		{
			name:                      "Mismatch Error MP4 and SRT Files",
			OriginalFiles:             []string{"a.mp4", "c.mp4", "a.srt", "b.srt", "c.srt"},
			VideoConfigExtension:      ".mp4",
			SubConfigExtension:        ".srt",
			TestRenameFileName:        "Hibi whatever_2",
			TestTotalNumberOfEpisodes: 3,
			wantErr:                   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()

			for _, f := range tt.OriginalFiles {
				if err := os.WriteFile(filepath.Join(dir, f), nil, 0o644); err != nil {
					t.Fatal(err)
				}
			}

			cfg := &Config{
				directory:              dir,
				VideoUserConfigName:    tt.VideoConfigExtension,
				SubtitleUserConfigName: tt.SubConfigExtension,
				RenameFileName:         tt.TestRenameFileName,
				TotalNumberOfEpisodes:  tt.TestTotalNumberOfEpisodes,
			}

			_, err := HandleFileRenaming(cfg, GetUserFiles(cfg))

			if (err != nil) != tt.wantErr {
				t.Fatalf("error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			got := GetUserFiles(cfg)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
