/*
@Time : 2021/4/23 21:09
@Author : LiuKun
@File : logExt
@Software: GoLand
@Description:
*/

package clog

type LogExt struct {
	title string
}

func NewLogExt(title string) *LogExt {
	t := new(LogExt)
	t.title = title
	return t
}

func (l *LogExt) Start() {
	Start("开始%s...\n", l.title)
}

func (l *LogExt) Completion(err error) {
	if err != nil {
		Fail("%s失败：%s\n", l.title, err.Error())
	} else {
		Success("%s成功\n", l.title)
	}
}

func (l *LogExt) Fail(err error) {
	if err != nil {
		Fail("%s失败：%s\n", l.title, err.Error())
	}
}

func (l *LogExt) Final() {
	Complete("%s完成\n", l.title)
}
