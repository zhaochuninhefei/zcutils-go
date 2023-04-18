package zcsync

// Caller zcsync调用者
type Caller struct {
	// 定义一个chan
	lock chan struct{}
}

// NewCaller 初始化一个zcsync调用者
func NewCaller() *Caller {
	caller := &Caller{
		lock: make(chan struct{}, 1),
	}
	// 向chan中放入一个空结构体，以便首次调用CallSync时不会阻塞
	caller.lock <- struct{}{}
	return caller
}

// CallSync 利用chan实现同步调用, 保证并发时只有一个goroutine能够执行目标函数, 其他goroutine被阻塞, 直到上一个goroutine执行完毕, 才能继续执行。
func (z *Caller) CallSync(f func() error) error {
	// 从chan获取一个空结构体, chan为空时会阻塞
	<-z.lock
	// 执行传入的函数
	err := f()
	// 向chan中放入一个空结构体, 以便下一个goroutine能够继续执行
	z.lock <- struct{}{}
	// 返回传入的函数的执行结果
	return err
}
