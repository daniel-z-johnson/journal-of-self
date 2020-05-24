package config

import "testing"

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("exampleConfig.json")
	if err != nil {
		t.Fatalf("There was an error: '%s'", err)
	}

	t.Logf("Config: '%+v'", config)
	if config.Database == nil {
		t.Fatal("database part of config should not be nil")
	}

	if config.Database.Database != "postgres" {
		t.Fatalf("Expected postgres but got '%s' instead",
			config.Database.Database)
	}

	if config.Database.Host != "127.0.0.1" {
		t.Fatalf("Expected host to be '127.0.0.1' but got '%s' instead",
			config.Database.Host)
	}
}

// really didn't need this test but felt like increasing test converage
func TestBadFileName(t *testing.T) {
	_, err := LoadConfig("noFile.txt")
	if err == nil {
		t.Fatal("Expected an error but didn't get one")
	}
}
