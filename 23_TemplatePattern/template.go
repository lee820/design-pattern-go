package temp

import "fmt"

//Game 模板基类
type Game struct {
	Initialize func()
	StartPlay  func()
	EndPlay    func()
}

//Play 模板基类的Play方法
func (g Game) Play() {
	g.Initialize()
	g.StartPlay()
	g.EndPlay()
}

//FootBall 子类，继承ame类
type FootBall struct {
	Game
}

//NewFootBall 实例化football子类
func NewFootBall() *FootBall {
	ft := new(FootBall)
	ft.Game.Initialize = ft.Initialize
	ft.Game.StartPlay = ft.StartPlay
	ft.Game.EndPlay = ft.EndPlay
	return ft
}

//Initialize 子类的Initialize方法
func (ft *FootBall) Initialize() {
	fmt.Println("Football game initialize")
}

//StartPlay 子类的StartPlay方法
func (ft *FootBall) StartPlay() {
	fmt.Println("Football game started.")
}

//EndPlay 子类的EndPlay方法
func (ft *FootBall) EndPlay() {
	fmt.Println("Football game Finished!")
}

//Basketball 篮球游戏子类
type Basketball struct {
	Game
}

//NewBasketball 实例化篮球游戏子类
func NewBasketball() *Basketball {
	basketball := new(Basketball)
	basketball.Game.Initialize = basketball.Initialize
	basketball.Game.StartPlay = basketball.StartPlay
	basketball.Game.EndPlay = basketball.EndPlay
	return basketball
}

//Initialize 篮球游戏子类的初始化方法
func (b *Basketball) Initialize() {
	fmt.Println("Basketball game initialize")
}

//StartPlay 篮球游戏子类的开始方法
func (b *Basketball) StartPlay() {
	fmt.Println("Basketball game started.")
}

//EndPlay 篮球游戏子类的结束方法
func (b *Basketball) EndPlay() {
	fmt.Println("Basketball game Finished!")
}
