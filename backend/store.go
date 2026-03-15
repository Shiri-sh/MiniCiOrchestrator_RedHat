package main

import(
	"fmt"
)

func (app *App) AddBuild(b Build) Build {
	app.Mu.Lock()
	defer app.Mu.Unlock()
	b.ID = app.NextBuildID
	app.NextBuildID++
	app.Builds = append(app.Builds, b)
	return b
}

func (app *App) GetAllBuilds() []Build {
	app.Mu.Lock()
	defer app.Mu.Unlock()
	//returns a copy of the builds slice to prevent external modification
	copiedBuilds := make([]Build, len(app.Builds))
	copy(copiedBuilds, app.Builds)
	return copiedBuilds
}

func (app *App) UpdateBuildStatus(buildID int, status string) error {
	app.Mu.Lock()
	defer app.Mu.Unlock()
	for i, b := range app.Builds {
		if b.ID == buildID {
			app.Builds[i].Status = status
			return nil
		}
	}
	return fmt.Errorf("Build with ID %d not found", buildID)
}