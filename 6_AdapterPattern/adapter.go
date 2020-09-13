package adapter

import "fmt"

//MediaPlayer 多媒体播放器接口
type MediaPlayer interface {
	Play(audioType, fileName string)
}

//AdvanceMediaPlayer 新多媒体播放器接口
type AdvanceMediaPlayer interface {
	PlayRmvb(fileName string)
	PlayMP4(fileName string)
}

//RmvbPlayer rmvb播放器类
type RmvbPlayer struct{}

//MP4Player mp4播放器类
type MP4Player struct{}

//NewRmvbPlayer 实例化rmvb播放器类
func NewRmvbPlayer() *RmvbPlayer {
	return &RmvbPlayer{}
}

//PlayRmvb 播放rmvb格式视频方法
func (rmvb *RmvbPlayer) PlayRmvb(fileName string) {
	fmt.Println("Playing rmvb file. Name: ", fileName)
}

//PlayMP4 为了实现接口所写，不使用这个方法
func (rmvb *RmvbPlayer) PlayMP4(fileName string) {
	//do nothing
}

//NewMp4Player 实例化mp4播放器类
func NewMp4Player() *MP4Player {
	return &MP4Player{}
}

//PlayRmvb 为了实现接口，不使用这个方法
func (mp4 *MP4Player) PlayRmvb(fileName string) {
	//do nothing
}

//PlayMP4 播放MP4文件方法
func (mp4 *MP4Player) PlayMP4(fileName string) {
	fmt.Println("Playing mp4 file. Name: ", fileName)
}

//MediaAdapter 多媒体播放器适配器
type MediaAdapter struct {
	MusicPlayer AdvanceMediaPlayer
}

//NewMediaAdapter 初始化适配器
func NewMediaAdapter(audioType string) *MediaAdapter {
	switch audioType {
	case "rmvb":
		return &MediaAdapter{MusicPlayer: NewRmvbPlayer()}
	case "mp4":
		return &MediaAdapter{MusicPlayer: NewMp4Player()}
	}
	return nil
}

//Play 根据文件类型，调用相应的适配器进行播放
func (mAdapter *MediaAdapter) Play(audioType, fileName string) {
	switch audioType {
	case "rmvb":
		mAdapter.MusicPlayer.PlayRmvb(fileName)
		break
	case "mp4":
		mAdapter.MusicPlayer.PlayMP4(fileName)
		break
	}
}

//AudioPlayer 音频播放器类
type AudioPlayer struct {
	mediaAdapter *MediaAdapter
}

//NewAudioPlayer 音频播放器实例化
func NewAudioPlayer() *AudioPlayer {
	return &AudioPlayer{}
}

//Play 播放音频
func (auPlayer *AudioPlayer) Play(audioType, fileName string) {
	if audioType == "mp3" {
		fmt.Println("Playing mp3 file. Name: ", fileName)
	} else if audioType == "rmvb" || audioType == "mp4" {
		auPlayer.mediaAdapter = NewMediaAdapter(audioType)
		if auPlayer.mediaAdapter == nil {
			fmt.Println("mediaAdapter create fail.")
		}
		auPlayer.mediaAdapter.Play(audioType, fileName)
	} else {
		fmt.Println("Invalid media type. ", audioType, "format not supported")
	}
}
