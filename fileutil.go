package quickfix

import "strings"

func sessionIDFilenamePrefix(s SessionID) string {
	sender := []string{s.SenderCompID}
	if s.SenderSubID != "" {
		sender = append(sender, s.SenderSubID)
	}
	if s.SenderLocationID != "" {
		sender = append(sender, s.SenderLocationID)
	}

	target := []string{s.TargetCompID}
	if s.TargetSubID != "" {
		target = append(target, s.TargetSubID)
	}
	if s.TargetLocationID != "" {
		target = append(target, s.TargetLocationID)
	}

	fname := []string{s.BeginString, strings.Join(sender, "_"), strings.Join(target, "_")}
	if s.Qualifier != "" {
		fname = append(fname, s.Qualifier)
	}
	return strings.Join(fname, "-")
}
