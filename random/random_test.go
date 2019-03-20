package random

import (
	"math"
	"testing"

	"github.com/namnhce/weighted-random-messages/random/util"
)

const (
	AcceptedThreshold = 0.69
)

func Test_F(t *testing.T) {
	type args struct {
		in    []int64
		times int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test with minimum quantity of version list",
			args: args{
				in:    []int64{50, 30, 60},
				times: 1000,
			},

			want: AcceptedThreshold,
		},
		{
			name: "Test with maximum quantity of version list",
			args: args{
				in: []int64{
					0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
					132413220, 121411241, 124115122, 235124343, 53432514, 153262365, 163343436, 446313743, 414822459, 351615449,
					232413220, 221411241, 124115122, 235124343, 53432524, 253262365, 263343436, 446323743, 424822459, 351625449,
					332413220, 321411241, 124115122, 235124343, 53432534, 353262365, 363343436, 446333743, 434822459, 351635449,
					432413220, 421411241, 124115122, 235124343, 53432544, 453262365, 463343436, 446343743, 444822459, 351645449,
					532413220, 521411241, 124115122, 235124343, 53432554, 553262365, 563343436, 446353743, 454822459, 351655449,
					632413220, 621411241, 124115122, 235124343, 53432564, 653262365, 663343436, 446363743, 464822459, 351665449,
					732413220, 721411241, 124115122, 235124343, 53432574, 753262365, 763343436, 446373743, 474822459, 351675449,
					832413220, 821411241, 124115122, 235124343, 53432584, 853262365, 863343436, 446383743, 484822459, 351685449,
					932413220, 921411241, 124115122, 235124343, 53432594, 953262365, 963343436, 446393743, 494822459, 351695449,
				},
				times: 1000,
			},

			want: AcceptedThreshold,
		},
		{
			name: "Test with big weight of version",
			args: args{
				in: []int64{
					0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
					10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
					20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
					30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
					40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
					50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
					60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
					70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
					80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
					90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
				},
				times: 1000,
			},

			want: AcceptedThreshold,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := []int64{}

			for i := 0; i < tt.args.times; i++ {
				x := F(tt.args.in)
				got = append(got, x)
			}

			sorted := util.SortInt64s(got)
			rs := make(map[int64]int64)

			for _, g := range sorted {
				_, ok := rs[g]
				if ok {
					rs[g]++
				} else {
					rs[g] = 1
				}
			}

			inResult := []int64{}
			for _, v := range rs {
				inResult = append(inResult, v)
			}

			actualResult := util.CalFrequency(inResult)
			expectedResult := util.CalFrequency(tt.args.in)

			testResult := true
			failedThresshold := 0.0

			for i := 0; i < len(actualResult); i++ {
				x := math.Abs(actualResult[i] - expectedResult[i])
				if x > tt.want {
					failedThresshold = x
					testResult = false
					break
				}
			}

			// Compare the rate
			if !testResult {
				t.Errorf("Message Version Random is not in accepted threshold() error = %v, wantErr %v", failedThresshold, tt.want)
			}
		})
	}
}

func Test_validate(t *testing.T) {
	type args struct {
		in []int64
	}

	// empty
	emptyArr := []int64{}

	// 100 elements
	validVersions := []int64{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
		60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
		70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
		80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
		90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
	}

	// 101 elements
	invalidVersions := []int64{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
		60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
		70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
		80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
		90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
		100,
	}

	validWeight := []int64{0, 1, 2, 3, 4}
	invalidWeight := []int64{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		132413220, 121411241, 124115122, 235124343, 53432514, 153262365, 163343436, 446313743, 414822459, 3351615449,
		232413220, 221411241, 124115122, 235124343, 53432524, 253262365, 263343436, 446323743, 424822459, 3351625449,
		332413220, 321411241, 124115122, 235124343, 53432534, 353262365, 363343436, 446333743, 434822459, 3351635449,
		432413220, 421411241, 124115122, 235124343, 53432544, 453262365, 463343436, 446343743, 444822459, 3351645449,
		532413220, 521411241, 124115122, 235124343, 53432554, 553262365, 563343436, 446353743, 454822459, 3351655449,
		632413220, 621411241, 124115122, 235124343, 53432564, 653262365, 663343436, 446363743, 464822459, 3351665449,
		732413220, 721411241, 124115122, 235124343, 53432574, 753262365, 763343436, 446373743, 474822459, 3351675449,
		832413220, 821411241, 124115122, 235124343, 53432584, 853262365, 863343436, 446383743, 484822459, 3351685449,
		932413220, 921411241, 124115122, 235124343, 53432594, 953262365, 963343436, 446393743, 494822459, 3351695449,
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "[Test_validate] - valid version length",
			args: args{
				in: validVersions,
			},
			wantErr: nil,
		},
		{
			name: "[Test_validate] - invalid version length",
			args: args{
				in: invalidVersions,
			},
			wantErr: errorVersionLengthInvalid,
		},
		{
			name: "[Test_validate] - invalid with empty version list",
			args: args{
				in: emptyArr,
			},
			wantErr: errorVersionLengthInvalid,
		},
		{
			name: "[Test_validate] - valid weight",
			args: args{
				in: validWeight,
			},
			wantErr: nil,
		},
		{
			name: "[Test_validate] - invalid weight",
			args: args{
				in: invalidWeight,
			},
			wantErr: errorWeightInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate(tt.args.in); err != nil && err != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}
