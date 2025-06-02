package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: just drag and drop a video file onto the executable or call it with a video file as an argument")
		return
	}

	inputFile := os.Args[1]
	ext := filepath.Ext(inputFile)
	base := strings.TrimSuffix(inputFile, ext)
	outputFile := fmt.Sprintf("%s_sneed%s", base, ext)

	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i", inputFile,
		"-f", "lavfi",
		"-i", "color=c=0x00000000@0.0:s=qcif:r=10,format=rgba",
		"-filter_complex",
		"[1]drawtext=text='Nigger':fontsize=48:fontcolor=#00000077[totile];[totile]tile=25x25[tooverlay];[0:v][tooverlay]overlay=x=sin(n / 5)*5:y=sin(n / 3) * 2",
		"-shortest",
		"-threads", "4",
		outputFile,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Running the watermark thing on %s...\n", inputFile)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("FFmpreg is currently sucking and fucking, not sneeding and feeding: %v\n", err)
		os.Exit(1)
	}
}
