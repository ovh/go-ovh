package govh

import (
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSig(t *testing.T) {
	// Sig should be: "$1$" + SHA1_HEX(AS+"+"+CK+"+"+METHOD+"+"+QUERY+"+"+BODY+"+"+TSTAMP)
	as := "EXEgWIz07P0HYwtQDs7cNIqCiQaWSuHF"
	ck := "MtSwSrPpNjqfVSmJhLbPyr2i45lSwPU1"
	method := "GET"
	query := "https://eu.api.ovh.com/1.0/domains/"
	var body string
	ts := int64(1366560945)

	// Get sig
	got := sig(as, ck, method, query, body, ts)

	// Expected
	expected := "$1$d3705e8afb27a0d2970a322b96550abfc67bb798"

	if got != expected {
		t.Fatalf("invalid sig: got %q, expected %q", got, expected)
	}
}

func TestTime(t *testing.T) {
	var apiCallCount int
	expectedDelay := 10 * time.Second
	expectedTime := time.Now().Add(expectedDelay)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiCallCount++
		fmt.Fprintf(w, "%d", expectedTime.Unix())
	}))
	defer ts.Close()

	// Create a fake client
	endpoint := Endpoint(ts.URL)
	client := NewClient(endpoint, "ak", "as", "ck")

	// Check that the time is correct
	gotTime, err := client.Time()
	if err != nil {
		t.Fatalf("expected not error, gotTime %q", err)
	}

	if expectedTime.Unix() != gotTime.Unix() {
		t.Errorf("expeted %q, gotTime %q", expectedTime, gotTime)
	}

	// Get the delay between the current time and the API
	gotDelay, err := client.delay()
	if err != nil {
		t.Fatalf("expected not error, got %q", err)
	}

	// The delay should be negative, we should use math.Floor to avoid dealing
	// with nanoseconds
	gotDelaySec := int(math.Floor(gotDelay.Seconds()))
	expectedDelaySec := int(math.Floor(-1 * expectedDelay.Seconds()))
	if gotDelaySec != expectedDelaySec {
		t.Errorf("expeted %d, gotTime %d", expectedDelaySec, gotDelaySec)
	}

	// The delay functionn is used inside a once methhod, let's check if the
	// call on the API to get the delay is really called once and only once.
	// Let's call delay one more time and check that the time API as only been
	// called twice (for the previous tests)
	_, err = client.delay()
	if err != nil {
		t.Fatalf("expected not error, got %q", err)
	}

	if apiCallCount != 2 {
		t.Errorf("expected 2 calls on time api, got %d", apiCallCount)
	}
}
