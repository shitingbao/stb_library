package rebot

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func load() {
	// robotgo.MouseSleep = 100

	robotgo.ScrollMouse(10, "up") // 滚动条,向上滚动10行
	time.Sleep(time.Second)
	robotgo.ScrollMouse(10, "down") // 滚动条,向下滚动10行

	robotgo.Scroll(0, -10)
	robotgo.Scroll(100, 0)

	// robotgo.MilliSleep(100)
	robotgo.ScrollSmooth(-10, 6)
	// robotgo.ScrollRelative(10, -100)

	robotgo.Move(10, 20) // 坐标移动，绝对的定位

	robotgo.MoveRelative(0, -10) // 相对位移

	// 一起使用
	// 按住鼠标左键    robotgo.MouseToggle(`down`, `left`)
	// 解除按住鼠标左键    robotgo.MouseToggle(`up`, `left`)}

	// 将鼠标移动到屏幕 x:800 y:400 的位置（闪现到指定位置）
	robotgo.MoveMouse(800, 400)

	// 将鼠标移动到屏幕 x:800 y:400 的位置（模仿人类操作）
	robotgo.MoveMouseSmooth(800, 400)

	// 将鼠标移动到屏幕 x:800 y:400 的位置（模仿人类操作）
	// 第3个参数：纵坐标x 的延迟到达时间    // 第4个参数：横坐标y 的延迟到达时间
	// robotgo.MoveMouseSmooth(800, 400, 20.0, 200.0)}

	// 移动鼠标到 x:800 y:400 后，双击鼠标左键    robotgo.MoveClick(800, 400, `left`, true)}

	// 获取当前鼠标所在的位置    x, y := robotgo.GetMousePos()

	robotgo.DragSmooth(10, 10) // 点击选择，拖动

	robotgo.Click("wheelRight")
	robotgo.Click("left", true) // true 双击

	robotgo.MoveSmooth(100, 200, 1.0, 10.0)

	robotgo.Toggle("left")
	robotgo.Toggle("left", "up")
}

func keydown() {
	// 模拟按下1个键：打开开始菜单（win）
	robotgo.KeyTap(`command`)
	// 模拟按下2个键：打开资源管理器（win + e）
	robotgo.KeyTap(`e`, `command`)
	// 模拟按下3个键：打开任务管理器（Ctrl + Shift + ESC）
	robotgo.KeyTap(`esc`, `control`, `shift`)

	// 一直按住 A键不放
	robotgo.KeyToggle(`a`, `down`)
	// 解除按住 A键
	robotgo.KeyToggle(`a`, `up`)

	robotgo.KeyTap("enter")     // 回车
	robotgo.KeyTap("backspace") // 删除，注意不要用大写

}
