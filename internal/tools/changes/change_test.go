package changes

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

func TestParseChangeType(t *testing.T) {
	var testCases = map[string]struct {
		input    string
		wantType ChangeType
		wantErr  string
	}{
		"feature":       {"feature", FeatureChangeType, ""},
		"feature-case":  {"FEATURE", FeatureChangeType, ""},
		"bugfix":        {"bugfix", BugFixChangeType, ""},
		"bugfix-case":   {"BugFix", BugFixChangeType, ""},
		"invalid":       {"not-a-type", "", "unknown change type: not-a-type"},
		"invalid-empty": {"", "", "unknown change type:"},
	}

	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			c, err := ParseChangeType(tt.input)
			if tt.wantErr != "" {
				if err == nil {
					t.Fatal("expected non-nil err, got nil")
				}

				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Errorf("expected err to contain %s, got %s", tt.wantErr, err.Error())
				}
			} else {
				if c != tt.wantType {
					t.Errorf("expected type %s, got %s", tt.wantType, c)
				}
			}
		})
	}
}

func TestNewChanges(t *testing.T) {
	var changeTests = map[string]struct {
		modules     []string
		changeType  ChangeType
		description string
		wantErr     bool
	}{
		"valid feature 1 module":      {[]string{"a"}, FeatureChangeType, "this is a description", false},
		"valid feature 2 modules":     {[]string{"a", "b"}, FeatureChangeType, "this is a description", false},
		"valid bugfix 2 modules":      {[]string{"a", "b"}, BugFixChangeType, "this is a description", false},
		"invalid missing description": {[]string{"a", "b"}, BugFixChangeType, "", true},
		"invalid missing modules":     {[]string{}, FeatureChangeType, "this is a description", true},
	}

	for name, tt := range changeTests {
		t.Run(name, func(t *testing.T) {
			changes, err := NewChanges(tt.modules, tt.changeType, tt.description)
			if err != nil && !tt.wantErr {
				t.Errorf("expected nil err, got %v", err)
			} else if err == nil {
				if tt.wantErr {
					t.Errorf("expected non-nil err, got nil")
				}

				if len(changes) != len(tt.modules) {
					t.Errorf("expected %d changes, got %d", len(tt.modules), len(changes))
				}

				for _, c := range changes {
					want := Change{
						ID:            c.ID,
						SchemaVersion: SchemaVersion,
						Module:        c.Module,
						Type:          tt.changeType,
						Description:   tt.description,
					}

					if diff := cmp.Diff(want, c); diff != "" {
						t.Errorf("expect changes to match:\n%v", diff)
					}
				}
			}
		})
	}
}

func TestChangeToTemplate(t *testing.T) {
	const wantTemplate = `modules:
- test
type: feature
description: test description

# type may be one of "feature" or "bugfix".
# multiple modules may be listed. A change metadata file will be created for each module.`

	template, err := ChangeToTemplate(Change{
		ID:          "test-feature-1",
		Module:      "test",
		Type:        FeatureChangeType,
		Description: "test description",
	})
	if err != nil {
		t.Fatalf("expected nil err, got %v", err)
	}

	if bytes.Compare(template, []byte(wantTemplate)) != 0 {
		t.Errorf("expected template \"%s\", got \"%s\"", wantTemplate, string(template))
	}
}

func TestTemplateToChanges(t *testing.T) {
	const template = `modules:
- test
type: feature
description: test description

# type may be one of "feature" or "bugfix".
# multiple modules may be listed. A change metadata file will be created for each module.`

	changes, err := TemplateToChanges([]byte(template))
	if err != nil {
		t.Fatalf("expected nil err, got %v", err)
	}

	if len(changes) != 1 {
		t.Fatalf("expected 1 change, got %d", len(changes))
	}

	change := changes[0]

	want := Change{
		ID:            change.ID,
		SchemaVersion: SchemaVersion,
		Module:        "test",
		Type:          FeatureChangeType,
		Description:   "test description",
	}

	if diff := cmp.Diff(want, change); diff != "" {
		t.Errorf("expect changes to match:\n%v", diff)
	}
}

// assertHasChangeLike asserts that the given changes contains a change with the same type and description as want.
func assertHasChangeLike(t *testing.T, changes []Change, want Change) bool {
	want.SchemaVersion = SchemaVersion
	for _, c := range changes {
		want.ID = c.ID

		if diff := cmp.Diff(want, c); diff == "" {
			return true
		}
	}

	want.ID = ""
	t.Errorf("expected changes to contain %v", want)
	return true
}