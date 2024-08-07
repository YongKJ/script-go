package PromptUtil

func PackageGoScript(crossBuildPath string) (string, []string) {
	return "busybox", []string{
		"bash",
		crossBuildPath,
	}
}

func CaptureVideoScreenshot(video string, picture string) (string, []string) {
	return "ffmpeg", []string{
		"-i",
		video,
		"-ss",
		"0:0:00",
		"-t",
		"0:0:01",
		"-r",
		"1",
		"-f",
		"image2",
		picture,
	}
}
