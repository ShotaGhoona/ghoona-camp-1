// Package value イベントドメインの値オブジェクトを定義する
package value

import (
	"errors"
	"time"
)

// TimeSlot イベントの開始時刻と終了時刻をペアで管理する値オブジェクト
// 時間の前後関係を保証し、不正な時間設定を防ぐ
// 朝活イベントの時間枠を適切に管理するためのドメインルールを内包
type TimeSlot struct {
	StartTime time.Time
	EndTime   time.Time
}

// NewTimeSlot 新しいタイムスロットを作成する
func NewTimeSlot(startTime, endTime time.Time) (*TimeSlot, error) {
	if endTime.Before(startTime) || endTime.Equal(startTime) {
		return nil, errors.New("end time must be after start time")
	}
	
	return &TimeSlot{
		StartTime: startTime,
		EndTime:   endTime,
	}, nil
}

// Duration タイムスロットの持続時間を計算する
func (ts TimeSlot) Duration() time.Duration {
	return ts.EndTime.Sub(ts.StartTime)
}

// IsValid タイムスロットが有効かどうかを判定する
func (ts TimeSlot) IsValid() bool {
	return ts.EndTime.After(ts.StartTime)
}