package main

import "sync"

// In-memory store for builds
var (
	builds []Build
	nextBuildID = 1
	mu sync.Mutex
)

func AddBuild(b Build) Build {
	mu.Lock()
	defer mu.Unlock()
	b.ID = nextBuildID
	nextBuildID++
	builds = append(builds, b)
	return b
}

func GetAllBuilds() []Build {
	mu.Lock()
	defer mu.Unlock()
	//returns a copy of the builds slice to prevent external modification
	copiedBuilds := make([]Build, len(builds))
	copy(copiedBuilds, builds)
	return copiedBuilds
}

func UpdateBuildStatus(buildID int, status string) {
	mu.Lock()
	defer mu.Unlock()
	for i, b := range builds {
		if b.ID == buildID {
			builds[i].Status = status
			break
		}
	}
}