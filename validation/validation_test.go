// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package validation

import "testing"

func TestIsValidIP(t *testing.T) {
	goodValues := []string{
		"::1",
		"2a00:79e0:2:0:f1c3:e797:93c1:df80",
		"::",
		"2001:4860:4860::8888",
		"::fff:1.1.1.1",
		"1.1.1.1",
		"1.1.1.01",
		"255.0.0.1",
		"1.0.0.0",
		"0.0.0.0",
	}
	for _, val := range goodValues {
		if msgs := IsValidIP(val); len(msgs) != 0 {
			t.Errorf("expected true for %q: %v", val, msgs)
		}
	}

	badValues := []string{
		"[2001:db8:0:1]:80",
		"myhost.mydomain",
		"-1.0.0.0",
		"[2001:db8:0:1]",
		"a",
	}
	for _, val := range badValues {
		if msgs := IsValidIP(val); len(msgs) == 0 {
			t.Errorf("expected false for %q", val)
		}
	}
}

func TestIsValidWsName(t *testing.T) {
	goodName := []string{
		"a",
		"a1",
		"a-b",
		"a1-b-3",
		"a2222",
	}
	for _, val := range goodName {
		if msgs := IsValidWsName(val); msgs != nil {
			t.Errorf("expected true for %q: %v", val, msgs)
		}
	}
	badName := []string{
		"-",
		"-a",
		"x=",
		"a-",
	}
	for _, val := range badName {
		if msgs := IsValidWsName(val); msgs == nil {
			t.Errorf("expected true for %q: %v", val, msgs)
		}
	}
}

func TestVerifyEmailFormat(t *testing.T) {
	mails := []string{
		"a@a.com", "a.com", "a@", "@b", "a@a-com.-",
	}
	for _, mail := range mails {
		t.Logf("mail %v, %v", mail, VerifyEmailFormat(mail))
	}
}
