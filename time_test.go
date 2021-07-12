package time_interval_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/zakon47/time_interval"
	"testing"
	"time"
)

func TestNewMarker1D(t *testing.T) {
	data := []struct {
		name   string
		metka  string
		result uint32
	}{
		{name: "test1_m", metka: " 1m", result: 60},
		{name: "test2_M", metka: "2M ", result: 5356800},
		{name: "test3_m", metka: " 10m", result: 600},
		{name: "test4_h", metka: "1h ", result: 3600},
		{name: "test5_d", metka: "1d", result: 86400},
		{name: "test6_w", metka: "1w", result: 604800},
		{name: "test7_mm", metka: "1M", result: 2678400},
		{name: "test8_!m", metka: "1!m", result: 0},
		{name: "test9_!n", metka: "77", result: 0},
		{name: "test10_!n", metka: "m77", result: 0},
		{name: "test11_!n", metka: "7m7", result: 0},
		{name: "test12_m7m", metka: "m7m", result: 0},
	}
	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.result, time_interval.NewMarker1D(test.metka).GetSeconds())
		})
	}
}
func TestMarkerName(t *testing.T) {
	data := []struct {
		name string
		swap bool //вернуть ответ в обратном порядке
		in   string
		out  string
	}{
		{name: "test1_m", swap: true, in: "1m", out: "m1"},
		{name: "test2_M", swap: false, in: "2M", out: "2M"},
		{name: "test3_m", swap: false, in: "10m", out: "10m"},
		{name: "test3_m", swap: true, in: "10m", out: "m10"},
		{name: "test4_h", swap: false, in: "1h", out: "1h"},
		{name: "test5_d", swap: false, in: "1d", out: "1d"},
		{name: "test5_d", swap: false, in: "d1d", out: "0d1d"},
		{name: "test5_d", swap: false, in: "1d1", out: "1d1"},
		{name: "test5_d", swap: true, in: "1d1", out: "d11"},
		{name: "test5_d", swap: false, in: "55", out: "55"},
		{name: "test5_d", swap: false, in: "dd", out: "0dd"},
	}
	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			if test.swap {
				assert.Equal(t, test.out, time_interval.NewMarker1D(test.in).GetNameD1())
			} else {
				assert.Equal(t, test.out, time_interval.NewMarker1D(test.in).GetName1D())
			}
		})
	}
}
func TestNewMarkerD1(t *testing.T) {
	data := []struct {
		name   string
		metka  string
		result uint32
	}{
		{name: "test1_m", metka: " m1", result: 60},
		{name: "test2_M", metka: "M2 ", result: 5356800},
		{name: "test3_m", metka: " m10", result: 600},
		{name: "test4_h", metka: "h1 ", result: 3600},
		{name: "test5_d", metka: "d1", result: 86400},
		{name: "test6_w", metka: "w1", result: 604800},
		{name: "test7_mm", metka: "M1", result: 2678400},
		{name: "test8_!m", metka: "!m1", result: 0},
		{name: "test9_!n", metka: "77", result: 0},
		{name: "test10_!n", metka: "77m", result: 0},
		{name: "test11_!n", metka: "7m7", result: 0},
		{name: "test12_m7m", metka: "m7m", result: 0},
	}
	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.result, time_interval.NewMarkerD1(test.metka).GetSeconds())
		})
	}
}
func TestMarkerNameLast(t *testing.T) {
	data := []struct {
		name string
		swap bool //вернуть ответ в обратном порядке
		in   string
		out  string
	}{
		{name: "test1_m", swap: true, in: "m1", out: "m1"},
		{name: "test2_M", swap: false, in: "M2", out: "2M"},
		{name: "test3_m", swap: false, in: "m10", out: "10m"},
		{name: "test3_m", swap: true, in: "m10", out: "m10"},
		{name: "test4_h", swap: false, in: "h1", out: "1h"},
		{name: "test5_d", swap: false, in: "d1", out: "1d"},
		{name: "test5_d", swap: false, in: "d1d", out: "0d1d"},
		{name: "test5_d", swap: false, in: "1d1", out: "11d"},
		{name: "test5_d", swap: false, in: "55", out: "55"},
		{name: "test5_d", swap: false, in: "dd", out: "0dd"},
	}
	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			if test.swap {
				assert.Equal(t, test.out, time_interval.NewMarkerD1(test.in).GetNameD1())
			} else {
				assert.Equal(t, test.out, time_interval.NewMarkerD1(test.in).GetName1D())
			}
		})
	}
}

var tt0 = time.Date(2020, time.May, 7, 22, 30, 0, 0, time.UTC)   //07-05-2020 22:30:00
var tt1 = time.Date(2020, time.May, 7, 22, 31, 0, 0, time.UTC)   //07-05-2020 22:31:00
var tt2 = time.Date(2020, time.May, 7, 22, 31, 14, 0, time.UTC)  //07-05-2020 22:31:14
var tt3 = time.Date(2020, time.May, 7, 22, 31, 29, 0, time.UTC)  //07-05-2020 22:31:29
var tt4 = time.Date(2020, time.May, 7, 22, 31, 30, 0, time.UTC)  //07-05-2020 22:31:30
var tt5 = time.Date(2020, time.May, 7, 22, 31, 31, 0, time.UTC)  //07-05-2020 22:31:31
var tt6 = time.Date(2020, time.May, 7, 22, 31, 59, 0, time.UTC)  //07-05-2020 22:31:59
var tt7 = time.Date(2020, time.May, 7, 22, 32, 00, 0, time.UTC)  //07-05-2020 22:32:00
var tt8 = time.Date(2020, time.May, 7, 22, 35, 33, 0, time.UTC)  //07-05-2020 22:35:33
var tt9 = time.Date(2020, time.May, 7, 22, 33, 59, 0, time.UTC)  //07-05-2020 22:33:59
var tt10 = time.Date(2020, time.May, 7, 22, 31, 30, 0, time.UTC) //07-05-2020 22:31:30
var tt11 = time.Date(2020, time.May, 7, 22, 35, 00, 0, time.UTC) //07-05-2020 22:35:00
var tt12 = time.Date(2020, time.May, 7, 22, 33, 00, 0, time.UTC) //07-05-2020 22:33:00
var tt13 = time.Date(2020, time.May, 7, 22, 40, 00, 0, time.UTC) //07-05-2020 22:40:00
var tt14 = time.Date(2020, time.May, 7, 22, 36, 00, 0, time.UTC) //07-05-2020 22:36:00
var tt15 = time.Date(2020, time.May, 7, 22, 34, 00, 0, time.UTC) //07-05-2020 22:34:00
var tt16 = time.Date(2020, time.May, 7, 22, 40, 00, 0, time.UTC) //07-05-2020 22:40:00

func TestPointer_Ceil(t *testing.T) {
	data := []struct {
		name   string
		time   time.Time
		diff   uint32
		result time.Time
	}{
		{name: "test0", time: tt6, diff: 120, result: tt0},
		{name: "test1", time: tt1, diff: 60, result: tt1},
		{name: "test2", time: tt1, diff: 30, result: tt1},
		{name: "test3", time: tt2, diff: 30, result: tt1},
		{name: "test4", time: tt3, diff: 30, result: tt1},
		{name: "test5", time: tt4, diff: 30, result: tt4},
		{name: "test6", time: tt5, diff: 60, result: tt1},
		{name: "test7", time: tt5, diff: 30, result: tt4},
		{name: "test8", time: tt6, diff: 30, result: tt4},
		{name: "test9", time: tt6, diff: 60, result: tt1},
		{name: "test10", time: tt7, diff: 120, result: tt7},

		{name: "test11", time: tt8, diff: time_interval.NewMarker1D("5m").GetSeconds(), result: tt11},
		{name: "test12", time: tt8, diff: time_interval.NewMarker1D("10m").GetSeconds(), result: tt0},
		{name: "test13", time: tt9, diff: time_interval.NewMarker1D("5m").GetSeconds(), result: tt0},
		{name: "test14", time: tt10, diff: time_interval.NewMarker1D("3m").GetSeconds(), result: tt0},
		{name: "test15", time: tt11, diff: time_interval.NewMarker1D("5m").GetSeconds(), result: tt11},
		{name: "test16", time: tt11, diff: time_interval.NewMarker1D("3m").GetSeconds(), result: tt12},
		{name: "test22", time: tt7, diff: 0, result: tt7},
		{name: "test23", time: tt8, diff: 0, result: tt8},
	}
	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.result.Unix(), time_interval.Pointer(test.time.Unix()).Ceil(test.diff).Time().Unix())
		})
	}
}
func TestPointer_Floor(t *testing.T) {
	data := []struct {
		name   string
		time   time.Time
		diff   uint32
		result time.Time
	}{
		{name: "test0", time: tt6, diff: 120, result: tt7},
		{name: "test1", time: tt1, diff: 60, result: tt7},
		{name: "test2", time: tt1, diff: 30, result: tt4},
		{name: "test3", time: tt2, diff: 30, result: tt4},
		{name: "test4", time: tt3, diff: 30, result: tt4},
		{name: "test5", time: tt4, diff: 30, result: tt7},
		{name: "test6", time: tt5, diff: 60, result: tt7},
		{name: "test7", time: tt5, diff: 30, result: tt7},
		{name: "test8", time: tt6, diff: 30, result: tt7},
		{name: "test9", time: tt6, diff: 60, result: tt7},
		{name: "test10", time: tt7, diff: 120, result: tt15},

		{name: "test11", time: tt8, diff: time_interval.NewMarker1D("5m").GetSeconds(), result: tt13},
		{name: "test12", time: tt8, diff: time_interval.NewMarker1D("10m").GetSeconds(), result: tt13},
		{name: "test13", time: tt9, diff: time_interval.NewMarker1D("5m").GetSeconds(), result: tt11},
		{name: "test14", time: tt10, diff: time_interval.NewMarker1D("3m").GetSeconds(), result: tt12},
		{name: "test15", time: tt11, diff: time_interval.NewMarker1D("5m").GetSeconds(), result: tt16},
		{name: "test16", time: tt11, diff: time_interval.NewMarker1D("3m").GetSeconds(), result: tt14},
		{name: "test22", time: tt7, diff: 0, result: tt7},
		{name: "test23", time: tt8, diff: 0, result: tt8},
	}
	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.result.Unix(), time_interval.Pointer(test.time.Unix()).Floor(test.diff).Time().Unix())
		})
	}
}

func TestInterval_Linear(t *testing.T) {
	data := []struct {
		name   string
		from   time.Time
		to     time.Time
		step   uint32
		result []int64
	}{
		{name: "test1", from: tt6, to: tt6.Add(5 * time.Minute), step: uint32(60), result: []int64{1588890660, 1588890720, 1588890780, 1588890840, 1588890900, 1588890960, 1588891020}},
		{name: "test1", from: tt6, to: tt6.Add(-5 * time.Minute), step: uint32(60), result: []int64{1588890720, 1588890660, 1588890600, 1588890540, 1588890480, 1588890420, 1588890360}},
		{name: "test1", from: tt6, to: tt6.Add(5 * time.Minute), step: uint32(0), result: []int64{}},
		{name: "test1", from: tt6, to: tt6.Add(-5 * time.Minute), step: uint32(0), result: []int64{}},
		{name: "test3", from: tt6, to: tt6.Add(-1*time.Minute + 10*time.Second), step: uint32(18), result: []int64{1588890726, 1588890708, 1588890690, 1588890672, 1588890654}},
		{name: "test4", from: tt6, to: tt6.Add(1*time.Minute + 10*time.Second), step: uint32(18), result: []int64{1588890708, 1588890726, 1588890744, 1588890762, 1588890780, 1588890798}},
		{name: "test5", from: tt6, to: tt6, step: uint32(18), result: []int64{1588890726, 1588890708}},
	}
	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			interv := time_interval.NewInterval(test.from, test.to)
			assert.Equal(t, test.result, interv.Linear(test.step))
		})
	}
}
