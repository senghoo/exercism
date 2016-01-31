package hamming

import "testing"

const testVersion = 2

func TestHamming(t *testing.T) {
	if TestVersion != testVersion {
		t.Errorf("Found TestVersion = %v, want %v.", TestVersion, testVersion)
	}
	for _, tc := range testCases {
		switch got, err := Distance(tc.s1, tc.s2); {
		case err != nil:
			if tc.want >= 0 {
				t.Fatalf("Distance(%q, %q) returned error: %v",
					tc.s1, tc.s2, err)
			}
		case tc.want < 0:
			t.Fatalf("Distance(%q, %q) = %d.  Expected error.",
				tc.s1, tc.s2, got)
		case got != tc.want:
			t.Fatalf("Distance(%q, %q) = %d, want %d.",
				tc.s1, tc.s2, got, tc.want)
		}
	}
}

func BenchmarkHamming(b *testing.B) {
	// bench combined time to run through all test cases
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Distance(tc.s1, tc.s2)
		}
	}
}
