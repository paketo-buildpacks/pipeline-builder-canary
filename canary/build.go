package canary

import (
	"fmt"
	"github.com/buildpacks/libcnb"
	"github.com/paketo-buildpacks/libpak"
	"github.com/paketo-buildpacks/libpak/bard"
)

type Build struct {
	Logger bard.Logger
}

func (b Build) Build(context libcnb.BuildContext) (libcnb.BuildResult, error) {
	b.Logger.Title(context.Buildpack)
	result := libcnb.NewBuildResult()

	r := libpak.PlanEntryResolver{Plan: context.Plan}
	_, ok, err := r.Resolve("canary")
	if err != nil {
		return libcnb.BuildResult{}, fmt.Errorf("unable to resolve buildpack plan entry canary\n%w", err)
	} else if !ok {
		return libcnb.BuildResult{}, nil
	}

	return result, nil
}
