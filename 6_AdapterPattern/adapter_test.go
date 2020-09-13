package adapter

import "testing"

func Test(t *testing.T) {
	t.Run("audio Test:", AudioPlayerTest)
}

func AudioPlayerTest(t *testing.T) {
	audio := NewAudioPlayer()
	audio.Play("mp4", "yello.mp4")
	audio.Play("rmvb", "夜的第七章.rmvb")
	audio.Play("mp3", "以父之名.mp3")
	audio.Play("avi", "viva la Vida.avi")
}
