package metrics

import (
	"testing"
)

func TestTokenLabels_Add(t *testing.T) {
	type args struct {
		label string
		value string
	}
	tests := []struct {
		name       string
		labels     TokenLabels
		args       args
		expect     TokenLabels
		expectSize int
	}{{
		name:   "Normal way",
		labels: TokenLabels{},
		args: args{
			label: "host",
			value: "A",
		},
		expect: TokenLabels{{
			Key:   "host",
			Value: "A",
		}},
		expectSize: 1,
	}, {
		name:   "Empty key does nothing",
		labels: TokenLabels{},
		args: args{
			label: "",
			value: "A",
		},
		expect:     TokenLabels{},
		expectSize: 0,
	}, {
		name:   "Empty value does nothing",
		labels: TokenLabels{},
		args: args{
			label: "host",
			value: "",
		},
		expect:     TokenLabels{},
		expectSize: 0,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			labels := tt.labels.Add(tt.args.label, tt.args.value)

			if len(labels) != tt.expectSize {
				t.Errorf("Bad size on TokenLabels{} expect %d, got %d", tt.expectSize, len(labels))
			}

			for i, tokenLabel := range labels {
				if tokenLabel.Key != tt.expect[i].Key {
					t.Errorf("Bad key on TokenLabel expect %s, got %s", tt.expect[i].Key, tokenLabel.Key)
				}
				if tokenLabel.Value != tt.expect[i].Value {
					t.Errorf("Bad value on TokenLabel expect %s, got %s", tt.expect[i].Value, tokenLabel.Value)
				}
			}
		})
	}
}
