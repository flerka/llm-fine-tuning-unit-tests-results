package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListReleases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		githubReleaseURL string
		githubToken      string
		want      []string
		wantErr   bool
	}{
		{
			name: "success",
			githubReleaseURL: "https://api.github.com/repos/tenv/tenv/releases",
			want: []string{
				"v1.0.0",
				"v1.0.1",
				"v1.0.2",
				"v1.0.3",
				"v1.0.4",
				"v1.0.5",
				"v1.0.6",
				"v1.0.7",
				"v1.0.8",
				"v1.0.9",
				"v1.0.10",
				"v1.0.11",
				"v1.0.12",
				"v1.0.13",
				"v1.0.14",
				"v1.0.15",
				"v1.0.16",
				"v1.0.17",
				"v1.0.18",
				"v1.0.19",
				"v1.0.20",
				"v1.0.21",
				"v1.0.22",
				"v1.0.23",
				"v1.0.24",
				"v1.0.25",
				"v1.0.26",
				"v1.0.27",
				"v1.0.28",
				"v1.0.29",
				"v1.0.30",
				"v1.0.31",
				"v1.0.32",
				"v1.0.33",
				"v1.0.34",
				"v1.0.35",
				"v1.0.36",
				"v1.0.37",
				"v1.0.38",
				"v1.0.39",
				"v1.0.40",
				"v1.0.41",
				"v1.0.42",
				"v1.0.43",
				"v1.0.44",
				"v1.0.45",
				"v1.0.46",
				"v1.0.47",
				"v1.0.48",
				"v1.0.49",
				"v1.0.50",
				"v1.0.51",
				"v1.0.52",
				"v1.0.53",
				"v1.0.54",
				"v1.0.55",
				"v1.0.56",
				"v1.0.57",
				"v1.0.58",
				"v1.0.59",
				"v1.0.60",
				"v1.0.61",
				"v1.0.62",
				"v1.0.63",
				"v1.0.64",
				"v1.0.65",
				"v1.0.66",
				"v1.0.67",
				"v1.0.68",
				"v1.0.69",
				"v1.0.70",
				"v1.0.71",
				"v1.0.72",
				"v1.0.73",
				"v1.0.74",
				"v1.0.75",
				"v1.0.76",
				"v1.0.77",
				"v1.0.78",
				"v1.0.79",
				"v1.0.80",
				"v1.0.81",
				"v1.0.82",
				"v1.0.83",
				"v1.0.84",
				"v1.0.85",
				"v1.0.86",
				"v1.0.87",
				"v1.0.88",
				"v1.0.89",
				"v1.0.90",
				"v1.0.91",
				"v1.0.92",
				"v1.0.93",
				"v1.0.94",
				"v1.0.95",
				"v1.0.96",
				"v1.0.97",
				"v1.0.98",
				"v1.0.99",
				"v1.0.100",
				"v1.0.101",
				"v1.0.102",
				"v1.0.103",
				"v1.0.104",
				"v1.0.105",
				"v1.0.106",
				"v1.0.107",
				"v1.0.108",
				"v1.0.109",
				"v1.0.110",
				"v1.0.111",
				"v1.0.112",
				"v1.0.113",
				"v1.0.114",
				"v1.0.115",
				"v1.0.116",
				"v1.0.117",
				"v1.0.118",
				"v1.0.119",
				"v1.0.120",
				"v1.0.121",
				"v1.0.122",
				"v1.0.123",
				"v1.0.124",
				"v1.0.125",
				"v1.0.126",
				"v1.0.127",
				"v1.0.128",
				"v1.0.129",
				"v1.