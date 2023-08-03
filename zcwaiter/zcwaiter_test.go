package zcwaiter_test

import (
	"gitee.com/zhaochuninhefei/zcgolog/zclog"
	"gitee.com/zhaochuninhefei/zcutils-go/zcpath"
	"gitee.com/zhaochuninhefei/zcutils-go/zcwaiter"
	"testing"
	"time"
)

func TestWaiter_WaitUntil(t *testing.T) {
	zclog.Level = zclog.LOG_LEVEL_DEBUG
	type fields struct {
		waitMaxTimes  int
		waitMSPerTime time.Duration
	}
	type args struct {
		supplier func() bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestWaiter_WaitUntil",
			fields: fields{
				waitMaxTimes:  300,
				waitMSPerTime: 1000,
			},
			args: args{
				supplier: func() bool {
					exist, _ := zcpath.FileExists("testdata/test.txt")
					return exist
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 5秒后创建文件 testdata/test.txt
			go func() {
				time.Sleep(5 * time.Second)
				if err := createTestFile(); err != nil {
					t.Error(err)
					return
				}
			}()
			_ = zcwaiter.NewWaiter()
			w, err := zcwaiter.NewCustomWaiter(tt.fields.waitMaxTimes, tt.fields.waitMSPerTime)
			if err != nil {
				t.Fatal(err)
			}
			if err = w.WaitUntil(tt.args.supplier); err != nil {
				t.Fatal(err)
			}
			// 删除 testdata/test.txt
			if err = zcpath.RemoveFile("testdata/test.txt"); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func createTestFile() error {
	// 创建 testdata 目录
	ok, err := zcpath.CreateDir("testdata")
	if !ok {
		return err
	}
	err = zcpath.CreateFile("testdata/test.txt")
	if err != nil {
		return err
	}
	return nil
}
